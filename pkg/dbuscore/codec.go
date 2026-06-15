package dbuscore

import (
	"fmt"

	godbus "github.com/godbus/dbus/v5"
)

func Marshal(signature string, value any) (any, error) {
	if _, err := godbus.ParseSignature(signature); err != nil {
		return nil, fmt.Errorf("dbus: invalid input signature %q: %w", signature, err)
	}
	if typed, ok := value.(TypedValue); ok {
		return typedToDBus(signature, typed)
	}

	switch signature {
	case "s":
		v, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("dbus: expected string for signature s, got %T", value)
		}
		return v, nil
	case "u":
		return marshalUint32(value)
	case "i":
		return marshalInt32(value)
	case "o":
		return marshalObjectPath(value)
	case "g":
		return marshalSignature(value)
	case "v":
		return nil, fmt.Errorf("dbus: variant values must use dbus.variant(signature, value)")
	default:
		return nil, fmt.Errorf("dbus: unsupported input signature %q", signature)
	}
}

func Unmarshal(signature string, body []any) (any, error) {
	if signature != "" {
		if _, err := godbus.ParseSignature(signature); err != nil {
			return nil, fmt.Errorf("dbus: invalid output signature %q: %w", signature, err)
		}
	}
	if len(body) == 0 {
		return nil, nil
	}
	if len(body) == 1 {
		return normalizeDBusValue(body[0]), nil
	}
	out := make([]any, 0, len(body))
	for _, value := range body {
		out = append(out, normalizeDBusValue(value))
	}
	return out, nil
}

func typedToDBus(expected string, typed TypedValue) (any, error) {
	if expected != typed.Signature {
		return nil, fmt.Errorf("dbus: typed value signature %q does not match expected %q", typed.Signature, expected)
	}
	if expected == "v" {
		inner, ok := typed.Value.(TypedValue)
		if !ok {
			return nil, fmt.Errorf("dbus: variant payload must be a typed value")
		}
		value, err := Marshal(inner.Signature, inner.Value)
		if err != nil {
			return nil, err
		}
		parsed := godbus.ParseSignatureMust(inner.Signature)
		return godbus.MakeVariantWithSignature(value, parsed), nil
	}
	return Marshal(typed.Signature, typed.Value)
}

func marshalUint32(value any) (uint32, error) {
	switch v := value.(type) {
	case uint32:
		return v, nil
	case int:
		return Uint32FromFloat64(float64(v))
	case int64:
		return Uint32FromFloat64(float64(v))
	case float64:
		return Uint32FromFloat64(v)
	default:
		return 0, fmt.Errorf("dbus: expected uint32-compatible value, got %T", value)
	}
}

func marshalInt32(value any) (int32, error) {
	switch v := value.(type) {
	case int32:
		return v, nil
	case int:
		return Int32FromFloat64(float64(v))
	case int64:
		return Int32FromFloat64(float64(v))
	case float64:
		return Int32FromFloat64(v)
	default:
		return 0, fmt.Errorf("dbus: expected int32-compatible value, got %T", value)
	}
}

func marshalObjectPath(value any) (godbus.ObjectPath, error) {
	switch v := value.(type) {
	case godbus.ObjectPath:
		if !v.IsValid() {
			return "", fmt.Errorf("dbus: invalid object path %q", v)
		}
		return v, nil
	case string:
		path := godbus.ObjectPath(v)
		if !path.IsValid() {
			return "", fmt.Errorf("dbus: invalid object path %q", v)
		}
		return path, nil
	default:
		return "", fmt.Errorf("dbus: expected object path string, got %T", value)
	}
}

func marshalSignature(value any) (godbus.Signature, error) {
	switch v := value.(type) {
	case godbus.Signature:
		return v, nil
	case string:
		sig, err := godbus.ParseSignature(v)
		if err != nil {
			return godbus.Signature{}, fmt.Errorf("dbus: invalid signature %q: %w", v, err)
		}
		return sig, nil
	default:
		return godbus.Signature{}, fmt.Errorf("dbus: expected signature string, got %T", value)
	}
}

func normalizeDBusValue(value any) any {
	switch v := value.(type) {
	case godbus.ObjectPath:
		return string(v)
	case godbus.Signature:
		return v.String()
	case godbus.Variant:
		return normalizeDBusValue(v.Value())
	default:
		return v
	}
}
