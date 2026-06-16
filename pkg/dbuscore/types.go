package dbuscore

import (
	"fmt"
	"math"

	godbus "github.com/godbus/dbus/v5"
)

// TypedValue carries an explicit D-Bus signature for JavaScript values whose
// D-Bus type cannot be inferred safely from the JavaScript representation.
type TypedValue struct {
	Signature string
	Value     any
}

// NewTypedValue validates signature and returns a typed value wrapper.
func NewTypedValue(signature string, value any) (TypedValue, error) {
	if _, err := godbus.ParseSignature(signature); err != nil {
		return TypedValue{}, fmt.Errorf("dbus: invalid signature %q: %w", signature, err)
	}
	return TypedValue{Signature: signature, Value: value}, nil
}

func U32(value uint32) TypedValue { return TypedValue{Signature: "u", Value: value} }

func I32(value int32) TypedValue { return TypedValue{Signature: "i", Value: value} }

func ObjectPath(value string) (TypedValue, error) {
	path := godbus.ObjectPath(value)
	if !path.IsValid() {
		return TypedValue{}, fmt.Errorf("dbus: invalid object path %q", value)
	}
	return TypedValue{Signature: "o", Value: path}, nil
}

func Signature(value string) (TypedValue, error) {
	sig, err := godbus.ParseSignature(value)
	if err != nil {
		return TypedValue{}, fmt.Errorf("dbus: invalid signature %q: %w", value, err)
	}
	return TypedValue{Signature: "g", Value: sig}, nil
}

func Variant(signature string, value any) (TypedValue, error) {
	inner, err := NewTypedValue(signature, value)
	if err != nil {
		return TypedValue{}, err
	}
	return TypedValue{Signature: "v", Value: inner}, nil
}

func Uint32FromFloat64(value float64) (uint32, error) {
	if math.Trunc(value) != value || value < 0 || value > math.MaxUint32 {
		return 0, fmt.Errorf("dbus: value %v is not a uint32", value)
	}
	return uint32(value), nil
}

func Int32FromFloat64(value float64) (int32, error) {
	if math.Trunc(value) != value || value < math.MinInt32 || value > math.MaxInt32 {
		return 0, fmt.Errorf("dbus: value %v is not an int32", value)
	}
	return int32(value), nil
}
