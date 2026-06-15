package dbuscore

import "testing"

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
