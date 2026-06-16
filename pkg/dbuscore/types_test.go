package dbuscore

import "testing"

func TestTypedValueConstructors(t *testing.T) {
	if got := U32(42); got.Signature != "u" || got.Value != uint32(42) {
		t.Fatalf("U32 = %#v", got)
	}
	if got := I32(-7); got.Signature != "i" || got.Value != int32(-7) {
		t.Fatalf("I32 = %#v", got)
	}
	if _, err := ObjectPath("/com/example/App1"); err != nil {
		t.Fatalf("valid object path: %v", err)
	}
	if _, err := ObjectPath("not/a/path"); err == nil {
		t.Fatalf("expected invalid object path error")
	}
	if _, err := Signature("a{sv}"); err != nil {
		t.Fatalf("valid signature: %v", err)
	}
	if _, err := Signature("{"); err == nil {
		t.Fatalf("expected invalid signature error")
	}
}

func TestIntegerBounds(t *testing.T) {
	if _, err := Uint32FromFloat64(-1); err == nil {
		t.Fatalf("expected negative uint32 error")
	}
	if _, err := Uint32FromFloat64(1.5); err == nil {
		t.Fatalf("expected fractional uint32 error")
	}
	if _, err := Int32FromFloat64(1.5); err == nil {
		t.Fatalf("expected fractional int32 error")
	}
}
