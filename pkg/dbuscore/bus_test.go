package dbuscore

import (
	"context"
	"testing"
)

func TestBusCloseClosesTrackedSubscriptions(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	bus := &Bus{subscriptions: map[*Subscription]struct{}{}}
	sub := &Subscription{cancel: cancel}
	bus.registerSubscription(sub)
	if len(bus.subscriptions) != 1 {
		t.Fatalf("subscription was not registered")
	}
	if err := bus.Close(context.Background()); err != nil {
		t.Fatalf("close bus: %v", err)
	}
	select {
	case <-ctx.Done():
	default:
		t.Fatalf("subscription cancel was not called")
	}
	if len(bus.subscriptions) != 0 {
		t.Fatalf("subscriptions were not cleared")
	}
	if err := bus.Close(context.Background()); err != nil {
		t.Fatalf("second close: %v", err)
	}
}
