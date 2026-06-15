package dbuscore

import "testing"

func TestPolicyConnectDefaults(t *testing.T) {
	policy := DefaultPolicy()
	if err := policy.CheckConnect(BusSession); err != nil {
		t.Fatalf("session should be allowed: %v", err)
	}
	if err := policy.CheckConnect(BusSystem); err == nil {
		t.Fatalf("system should be denied by default")
	}
}

func TestPolicyCallPatterns(t *testing.T) {
	policy := Policy{AllowSessionBus: true, AllowCall: []string{"org.freedesktop.*"}}
	req := MethodCallRequest{Destination: "org.freedesktop.DBus", Interface: "org.freedesktop.DBus", Member: "GetId"}
	if err := policy.CheckCall(req); err != nil {
		t.Fatalf("call should be allowed: %v", err)
	}
	req.Destination = "com.example.Blocked"
	req.Interface = "com.example.Blocked"
	if err := policy.CheckCall(req); err == nil {
		t.Fatalf("call should be denied")
	}
}
