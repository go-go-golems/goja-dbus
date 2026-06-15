package dbusgoja

import (
	"context"
	"errors"
	"sync"

	"github.com/go-go-golems/goja-dbus/pkg/dbuscore"
)

type resourceRegistry struct {
	mu     sync.Mutex
	buses  map[*dbuscore.Bus]struct{}
	closed bool
}

func newResourceRegistry(lifetime context.Context) *resourceRegistry {
	if lifetime == nil {
		lifetime = context.Background()
	}
	registry := &resourceRegistry{buses: map[*dbuscore.Bus]struct{}{}}
	go func() {
		<-lifetime.Done()
		_ = registry.closeAll(context.Background())
	}()
	return registry
}

func (r *resourceRegistry) addBus(bus *dbuscore.Bus) {
	if r == nil || bus == nil {
		return
	}
	r.mu.Lock()
	if r.closed {
		r.mu.Unlock()
		_ = bus.Close(context.Background())
		return
	}
	r.buses[bus] = struct{}{}
	r.mu.Unlock()
}

func (r *resourceRegistry) removeBus(bus *dbuscore.Bus) {
	if r == nil || bus == nil {
		return
	}
	r.mu.Lock()
	delete(r.buses, bus)
	r.mu.Unlock()
}

func (r *resourceRegistry) closeAll(ctx context.Context) error {
	if r == nil {
		return nil
	}
	r.mu.Lock()
	if r.closed {
		r.mu.Unlock()
		return nil
	}
	r.closed = true
	buses := make([]*dbuscore.Bus, 0, len(r.buses))
	for bus := range r.buses {
		buses = append(buses, bus)
	}
	r.buses = nil
	r.mu.Unlock()

	var retErr error
	for _, bus := range buses {
		retErr = errors.Join(retErr, bus.Close(ctx))
	}
	return retErr
}
