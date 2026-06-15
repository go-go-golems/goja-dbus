package dbuscore

import (
	"fmt"
	"strings"
)

// Policy is the Go-side authority for D-Bus access. JavaScript may request a
// narrower policy, but later host integration should intersect that with a
// host-provided maximum policy before connecting.
type Policy struct {
	AllowSessionBus bool
	AllowSystemBus  bool
	AllowCall       []string
}

func DefaultPolicy() Policy {
	return Policy{AllowSessionBus: true, AllowCall: []string{"*"}}
}

func (p Policy) IsZero() bool {
	return !p.AllowSessionBus && !p.AllowSystemBus && len(p.AllowCall) == 0
}

func (p Policy) CheckConnect(kind BusKind) error {
	switch kind {
	case BusSession:
		if !p.AllowSessionBus {
			return fmt.Errorf("dbus: session bus is not allowed by policy")
		}
	case BusSystem:
		if !p.AllowSystemBus {
			return fmt.Errorf("dbus: system bus is not allowed by policy")
		}
	case BusAddress:
		if !p.AllowSessionBus && !p.AllowSystemBus {
			return fmt.Errorf("dbus: address bus is not allowed by policy")
		}
	default:
		return fmt.Errorf("dbus: unsupported bus kind %q", kind)
	}
	return nil
}

func (p Policy) CheckCall(req MethodCallRequest) error {
	if len(p.AllowCall) == 0 {
		return nil
	}
	candidates := []string{
		req.Destination,
		req.Interface,
		req.Interface + "." + req.Member,
		req.Destination + ":" + req.Interface,
		req.Destination + ":" + req.Interface + "." + req.Member,
	}
	for _, pattern := range p.AllowCall {
		for _, candidate := range candidates {
			if matchPattern(pattern, candidate) {
				return nil
			}
		}
	}
	return fmt.Errorf("dbus: call to %s %s.%s is not allowed by policy", req.Destination, req.Interface, req.Member)
}

func matchPattern(pattern, value string) bool {
	pattern = strings.TrimSpace(pattern)
	if pattern == "" {
		return false
	}
	if pattern == "*" || pattern == value {
		return true
	}
	if strings.HasSuffix(pattern, "*") {
		return strings.HasPrefix(value, strings.TrimSuffix(pattern, "*"))
	}
	return false
}
