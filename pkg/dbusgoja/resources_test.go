package dbusgoja

import (
	"context"
	"testing"
	"time"

	"github.com/go-go-golems/goja-dbus/pkg/dbuscore"
)

func TestResourceRegistryClosesBusesOnLifetimeCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	registry := newResourceRegistry(ctx)
	bus := &dbuscore.Bus{}
	registry.addBus(bus)
	cancel()

	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		if bus.Closed() {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatalf("bus was not closed after lifetime cancellation")
}

func TestResourceRegistryRemoveBusPreventsLifetimeClose(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	registry := newResourceRegistry(ctx)
	bus := &dbuscore.Bus{}
	registry.addBus(bus)
	registry.removeBus(bus)
	cancel()
	time.Sleep(20 * time.Millisecond)
	if bus.Closed() {
		t.Fatalf("removed bus was closed by registry")
	}
}
