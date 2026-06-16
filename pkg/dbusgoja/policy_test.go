package dbusgoja

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/go-go-golems/goja-dbus/pkg/dbuscore"
)

func TestDecodePolicyExplicitEmptyAllowCallDeniesAll(t *testing.T) {
	vm := goja.New()
	value := vm.ToValue(map[string]any{"allowCall": []any{}})
	policy, err := decodePolicy(vm, value, dbuscore.DefaultPolicy())
	if err != nil {
		t.Fatalf("decode policy: %v", err)
	}
	if !policy.AllowCallSet {
		t.Fatalf("AllowCallSet = false, want true for explicit allowCall")
	}
	if len(policy.AllowCall) != 0 {
		t.Fatalf("AllowCall = %#v, want empty", policy.AllowCall)
	}
	req := dbuscore.MethodCallRequest{Destination: "org.freedesktop.DBus", Interface: "org.freedesktop.DBus", Member: "GetId"}
	if err := policy.CheckCall(req); err == nil {
		t.Fatalf("explicit empty allowCall should deny calls")
	}
}

func TestDecodePolicyAllowAddressBus(t *testing.T) {
	vm := goja.New()
	value := vm.ToValue(map[string]any{"allowAddressBus": true})
	policy, err := decodePolicy(vm, value, dbuscore.DefaultPolicy())
	if err != nil {
		t.Fatalf("decode policy: %v", err)
	}
	if !policy.AllowAddressBus {
		t.Fatalf("AllowAddressBus = false, want true")
	}
}
