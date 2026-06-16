package dbuscore

import (
	"testing"

	godbus "github.com/godbus/dbus/v5"
)

func TestMarshalScalars(t *testing.T) {
	got, err := Marshal("s", "hello")
	if err != nil || got != "hello" {
		t.Fatalf("Marshal string = %#v, %v", got, err)
	}
	got, err = Marshal("u", U32(42))
	if err != nil || got != uint32(42) {
		t.Fatalf("Marshal u32 typed = %#v, %v", got, err)
	}
	got, err = Marshal("i", float64(-7))
	if err != nil || got != int32(-7) {
		t.Fatalf("Marshal i32 float = %#v, %v", got, err)
	}
	path, err := Marshal("o", "/com/example/App1")
	if err != nil || path == nil {
		t.Fatalf("Marshal object path = %#v, %v", path, err)
	}
	if _, err := Marshal("v", "plain"); err == nil {
		t.Fatalf("expected plain variant error")
	}
}

func TestMarshalRejectsMismatchedTypedValue(t *testing.T) {
	if _, err := Marshal("i", U32(42)); err == nil {
		t.Fatalf("expected signature mismatch error")
	}
}

func TestMarshalCompoundValues(t *testing.T) {
	got, err := Marshal("as", []any{"one", "two"})
	if err != nil {
		t.Fatalf("Marshal as: %v", err)
	}
	strings, ok := got.([]string)
	if !ok || len(strings) != 2 || strings[1] != "two" {
		t.Fatalf("as = %#v", got)
	}

	got, err = Marshal("a{sv}", map[string]any{"name": mustVariant(t, "s", "demo")})
	if err != nil {
		t.Fatalf("Marshal a{sv}: %v", err)
	}
	variants, ok := got.(map[string]godbus.Variant)
	if !ok || variants["name"].Value() != "demo" {
		t.Fatalf("a{sv} = %#v", got)
	}

	got, err = Marshal("(su)", []any{"count", uint32(7)})
	if err != nil {
		t.Fatalf("Marshal struct: %v", err)
	}
	items, ok := got.([]any)
	if !ok || len(items) != 2 || items[0] != "count" || items[1] != uint32(7) {
		t.Fatalf("struct = %#v", got)
	}
}

func TestUnmarshalNormalizesSingleValue(t *testing.T) {
	got, err := Unmarshal("s", []any{"hello"})
	if err != nil || got != "hello" {
		t.Fatalf("Unmarshal = %#v, %v", got, err)
	}
	got, err = Unmarshal("", nil)
	if err != nil || got != nil {
		t.Fatalf("Unmarshal empty = %#v, %v", got, err)
	}
}

func mustVariant(t *testing.T, signature string, value any) TypedValue {
	t.Helper()
	v, err := Variant(signature, value)
	if err != nil {
		t.Fatalf("variant: %v", err)
	}
	return v
}
