package dbuscore

import (
	"fmt"
	"strings"

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
	case "as":
		return marshalStringArray(value)
	case "au":
		return marshalUint32Array(value)
	case "ai":
		return marshalInt32Array(value)
	case "ao":
		return marshalObjectPathArray(value)
	case "av":
		return marshalVariantArray(value)
	case "a{sv}":
		return marshalStringVariantMap(value)
	default:
		if strings.HasPrefix(signature, "(") && strings.HasSuffix(signature, ")") {
			return marshalStruct(signature, value)
		}
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

func marshalStringArray(value any) ([]string, error) {
	items, err := toAnySlice(value)
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, len(items))
	for _, item := range items {
		v, ok := item.(string)
		if !ok {
			return nil, fmt.Errorf("dbus: expected string array item, got %T", item)
		}
		out = append(out, v)
	}
	return out, nil
}

func marshalUint32Array(value any) ([]uint32, error) {
	items, err := toAnySlice(value)
	if err != nil {
		return nil, err
	}
	out := make([]uint32, 0, len(items))
	for _, item := range items {
		v, err := marshalUint32(item)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, nil
}

func marshalInt32Array(value any) ([]int32, error) {
	items, err := toAnySlice(value)
	if err != nil {
		return nil, err
	}
	out := make([]int32, 0, len(items))
	for _, item := range items {
		v, err := marshalInt32(item)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, nil
}

func marshalObjectPathArray(value any) ([]godbus.ObjectPath, error) {
	items, err := toAnySlice(value)
	if err != nil {
		return nil, err
	}
	out := make([]godbus.ObjectPath, 0, len(items))
	for _, item := range items {
		v, err := marshalObjectPath(item)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, nil
}

func marshalVariantArray(value any) ([]godbus.Variant, error) {
	items, err := toAnySlice(value)
	if err != nil {
		return nil, err
	}
	out := make([]godbus.Variant, 0, len(items))
	for _, item := range items {
		typed, ok := item.(TypedValue)
		if !ok || typed.Signature != "v" {
			return nil, fmt.Errorf("dbus: av items must be variants")
		}
		v, err := typedToDBus("v", typed)
		if err != nil {
			return nil, err
		}
		out = append(out, v.(godbus.Variant))
	}
	return out, nil
}

func marshalStringVariantMap(value any) (map[string]godbus.Variant, error) {
	items, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("dbus: expected object/map for a{sv}, got %T", value)
	}
	out := make(map[string]godbus.Variant, len(items))
	for key, item := range items {
		typed, ok := item.(TypedValue)
		if !ok || typed.Signature != "v" {
			return nil, fmt.Errorf("dbus: a{sv} value for %q must be a variant", key)
		}
		v, err := typedToDBus("v", typed)
		if err != nil {
			return nil, err
		}
		out[key] = v.(godbus.Variant)
	}
	return out, nil
}

func marshalStruct(signature string, value any) ([]any, error) {
	items, err := toAnySlice(value)
	if err != nil {
		return nil, err
	}
	inner := strings.TrimSuffix(strings.TrimPrefix(signature, "("), ")")
	sigs, err := splitFlatSignatures(inner)
	if err != nil {
		return nil, err
	}
	if len(sigs) != len(items) {
		return nil, fmt.Errorf("dbus: struct %s expects %d values, got %d", signature, len(sigs), len(items))
	}
	out := make([]any, 0, len(items))
	for i, sig := range sigs {
		v, err := Marshal(sig, items[i])
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, nil
}

func toAnySlice(value any) ([]any, error) {
	switch v := value.(type) {
	case []any:
		return v, nil
	case []string:
		out := make([]any, len(v))
		for i, item := range v {
			out[i] = item
		}
		return out, nil
	case []uint32:
		out := make([]any, len(v))
		for i, item := range v {
			out[i] = item
		}
		return out, nil
	case []int32:
		out := make([]any, len(v))
		for i, item := range v {
			out[i] = item
		}
		return out, nil
	default:
		return nil, fmt.Errorf("dbus: expected array, got %T", value)
	}
}

func splitFlatSignatures(s string) ([]string, error) {
	var out []string
	for i := 0; i < len(s); {
		switch s[i] {
		case 's', 'u', 'i', 'o', 'g', 'v':
			out = append(out, s[i:i+1])
			i++
		case 'a':
			if i+1 >= len(s) {
				return nil, fmt.Errorf("dbus: incomplete array signature")
			}
			if s[i+1] == '{' {
				end := strings.IndexByte(s[i+1:], '}')
				if end < 0 {
					return nil, fmt.Errorf("dbus: incomplete dict signature")
				}
				end = i + 1 + end
				out = append(out, s[i:end+1])
				i = end + 1
			} else {
				out = append(out, s[i:i+2])
				i += 2
			}
		default:
			return nil, fmt.Errorf("dbus: unsupported flat signature token %q", s[i])
		}
	}
	return out, nil
}

func normalizeDBusValue(value any) any {
	switch v := value.(type) {
	case godbus.ObjectPath:
		return string(v)
	case godbus.Signature:
		return v.String()
	case godbus.Variant:
		return normalizeDBusValue(v.Value())
	case []any:
		out := make([]any, 0, len(v))
		for _, item := range v {
			out = append(out, normalizeDBusValue(item))
		}
		return out
	case []string:
		out := make([]any, 0, len(v))
		for _, item := range v {
			out = append(out, item)
		}
		return out
	case []uint32:
		out := make([]any, 0, len(v))
		for _, item := range v {
			out = append(out, item)
		}
		return out
	case []int32:
		out := make([]any, 0, len(v))
		for _, item := range v {
			out = append(out, item)
		}
		return out
	case map[string]godbus.Variant:
		out := make(map[string]any, len(v))
		for key, item := range v {
			out[key] = normalizeDBusValue(item)
		}
		return out
	default:
		return v
	}
}
