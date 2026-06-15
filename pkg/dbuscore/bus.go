package dbuscore

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	godbus "github.com/godbus/dbus/v5"
)

type BusKind string

const (
	BusSession BusKind = "session"
	BusSystem  BusKind = "system"
	BusAddress BusKind = "address"
)

type ConnectOptions struct {
	Kind    BusKind
	Address string
	Timeout time.Duration
	Policy  Policy
}

type Bus struct {
	conn    *godbus.Conn
	kind    BusKind
	policy  Policy
	timeout time.Duration

	mu     sync.Mutex
	closed bool
}

type Arg struct {
	Signature string
	Value     any
}

type MethodCallRequest struct {
	Destination     string
	Path            godbus.ObjectPath
	Interface       string
	Member          string
	Inputs          []Arg
	OutputSignature string
	Timeout         time.Duration
}

func Connect(ctx context.Context, opts ConnectOptions) (*Bus, error) {
	policy := opts.Policy
	if policy.IsZero() {
		policy = DefaultPolicy()
	}
	if err := policy.CheckConnect(opts.Kind); err != nil {
		return nil, err
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var (
		conn *godbus.Conn
		err  error
	)
	switch opts.Kind {
	case BusSession:
		conn, err = godbus.ConnectSessionBus()
	case BusSystem:
		conn, err = godbus.ConnectSystemBus()
	case BusAddress:
		if strings.TrimSpace(opts.Address) == "" {
			return nil, fmt.Errorf("dbus: address is required")
		}
		conn, err = godbus.Connect(opts.Address)
	default:
		err = fmt.Errorf("dbus: unsupported bus kind %q", opts.Kind)
	}
	if err != nil {
		return nil, err
	}
	return &Bus{conn: conn, kind: opts.Kind, policy: policy, timeout: opts.Timeout}, nil
}

func (b *Bus) Close(context.Context) error {
	if b == nil {
		return nil
	}
	b.mu.Lock()
	if b.closed {
		b.mu.Unlock()
		return nil
	}
	b.closed = true
	conn := b.conn
	b.conn = nil
	b.mu.Unlock()
	if conn == nil {
		return nil
	}
	return conn.Close()
}

func (b *Bus) Call(ctx context.Context, req MethodCallRequest) (any, error) {
	conn, err := b.connection()
	if err != nil {
		return nil, err
	}
	if err := b.policy.CheckCall(req); err != nil {
		return nil, err
	}
	if strings.TrimSpace(req.Destination) == "" {
		return nil, fmt.Errorf("dbus: destination is required")
	}
	if !req.Path.IsValid() {
		return nil, fmt.Errorf("dbus: invalid object path %q", req.Path)
	}
	if strings.TrimSpace(req.Interface) == "" || strings.TrimSpace(req.Member) == "" {
		return nil, fmt.Errorf("dbus: interface and member are required")
	}

	args := make([]any, 0, len(req.Inputs))
	for _, input := range req.Inputs {
		value, err := Marshal(input.Signature, input.Value)
		if err != nil {
			return nil, err
		}
		args = append(args, value)
	}

	timeout := req.Timeout
	if timeout == 0 {
		timeout = b.timeout
	}
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	obj := conn.Object(req.Destination, req.Path)
	call := obj.CallWithContext(ctx, req.Interface+"."+req.Member, 0, args...)
	if call.Err != nil {
		return nil, call.Err
	}
	return Unmarshal(req.OutputSignature, call.Body)
}

func (b *Bus) connection() (*godbus.Conn, error) {
	if b == nil {
		return nil, fmt.Errorf("dbus: nil bus")
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed || b.conn == nil {
		return nil, fmt.Errorf("dbus: bus is closed")
	}
	return b.conn, nil
}
