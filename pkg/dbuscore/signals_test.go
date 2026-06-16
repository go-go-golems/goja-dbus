package dbuscore

import (
	"testing"

	godbus "github.com/godbus/dbus/v5"
)

func TestSignalMatchOptionsValidatePath(t *testing.T) {
	_, err := SignalMatchRequest{Path: godbus.ObjectPath("not/a/path")}.MatchOptions()
	if err == nil {
		t.Fatalf("expected invalid path error")
	}
}

func TestSignalMatchOptionsAllowEmptyMatch(t *testing.T) {
	options, err := (SignalMatchRequest{}).MatchOptions()
	if err != nil {
		t.Fatalf("empty match should be valid: %v", err)
	}
	if len(options) != 0 {
		t.Fatalf("options = %d, want 0", len(options))
	}
}

func TestSubscriptionCloseIsIdempotentForEmptySubscription(t *testing.T) {
	sub := &Subscription{}
	if err := sub.Close(t.Context()); err != nil {
		t.Fatalf("first close: %v", err)
	}
	if err := sub.Close(t.Context()); err != nil {
		t.Fatalf("second close: %v", err)
	}
}
