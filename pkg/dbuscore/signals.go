package dbuscore

import (
	"context"
	"fmt"
	"strings"
	"sync"

	godbus "github.com/godbus/dbus/v5"
)

type SignalMatchRequest struct {
	Sender    string
	Path      godbus.ObjectPath
	Interface string
	Member    string
	Buffer    int
}

type SignalPayload struct {
	Sender string
	Path   string
	Name   string
	Body   []any
}

type SignalSink func(context.Context, SignalPayload) error

type Subscription struct {
	conn    *godbus.Conn
	ch      chan *godbus.Signal
	cancel  context.CancelFunc
	options []godbus.MatchOption
	onClose func(*Subscription)

	closeOnce sync.Once
	closeErr  error
}

func (b *Bus) Listen(ctx context.Context, req SignalMatchRequest, sink SignalSink) (*Subscription, error) {
	if sink == nil {
		return nil, fmt.Errorf("dbus: signal sink is required")
	}
	conn, err := b.connection()
	if err != nil {
		return nil, err
	}
	options, err := req.MatchOptions()
	if err != nil {
		return nil, err
	}
	if err := conn.AddMatchSignalContext(ctx, options...); err != nil {
		return nil, err
	}
	buffer := req.Buffer
	if buffer <= 0 {
		buffer = 16
	}
	ch := make(chan *godbus.Signal, buffer)
	conn.Signal(ch)
	subCtx, cancel := context.WithCancel(ctx)
	sub := &Subscription{conn: conn, ch: ch, cancel: cancel, options: options}
	b.registerSubscription(sub)
	sub.onClose = b.unregisterSubscription
	go sub.run(subCtx, sink)
	return sub, nil
}

func (r SignalMatchRequest) MatchOptions() ([]godbus.MatchOption, error) {
	options := []godbus.MatchOption{}
	if strings.TrimSpace(r.Sender) != "" {
		options = append(options, godbus.WithMatchSender(r.Sender))
	}
	if r.Path != "" {
		if !r.Path.IsValid() {
			return nil, fmt.Errorf("dbus: invalid signal object path %q", r.Path)
		}
		options = append(options, godbus.WithMatchObjectPath(r.Path))
	}
	if strings.TrimSpace(r.Interface) != "" {
		options = append(options, godbus.WithMatchInterface(r.Interface))
	}
	if strings.TrimSpace(r.Member) != "" {
		options = append(options, godbus.WithMatchMember(r.Member))
	}
	return options, nil
}

func (s *Subscription) Close(ctx context.Context) error {
	if s == nil {
		return nil
	}
	s.closeOnce.Do(func() {
		if s.cancel != nil {
			s.cancel()
		}
		if s.conn != nil && s.ch != nil {
			s.conn.RemoveSignal(s.ch)
		}
		if s.conn != nil && len(s.options) > 0 {
			s.closeErr = s.conn.RemoveMatchSignalContext(ctx, s.options...)
		}
		if s.onClose != nil {
			s.onClose(s)
		}
	})
	return s.closeErr
}

func (s *Subscription) run(ctx context.Context, sink SignalSink) {
	for {
		select {
		case <-ctx.Done():
			return
		case sig := <-s.ch:
			if sig == nil {
				continue
			}
			_ = sink(ctx, SignalPayload{
				Sender: sig.Sender,
				Path:   string(sig.Path),
				Name:   sig.Name,
				Body:   normalizeSignalBody(sig.Body),
			})
		}
	}
}

func normalizeSignalBody(body []any) []any {
	out := make([]any, 0, len(body))
	for _, value := range body {
		out = append(out, normalizeDBusValue(value))
	}
	return out
}

func (b *Bus) registerSubscription(sub *Subscription) {
	if b == nil || sub == nil {
		return
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return
	}
	if b.subscriptions == nil {
		b.subscriptions = map[*Subscription]struct{}{}
	}
	b.subscriptions[sub] = struct{}{}
}

func (b *Bus) unregisterSubscription(sub *Subscription) {
	if b == nil || sub == nil {
		return
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.subscriptions != nil {
		delete(b.subscriptions, sub)
	}
}
