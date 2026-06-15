package dbusgoja

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/go-go-golems/goja-dbus/pkg/dbuscore"
	godbus "github.com/godbus/dbus/v5"
)

const typedMarker = "__dbusTyped"

func typedValueToObject(vm *goja.Runtime, value dbuscore.TypedValue) *goja.Object {
	obj := vm.NewObject()
	_ = obj.Set(typedMarker, true)
	_ = obj.Set("signature", value.Signature)
	_ = obj.Set("value", typedPayloadToValue(vm, value.Value))
	return obj
}

func typedPayloadToValue(vm *goja.Runtime, value any) any {
	switch v := value.(type) {
	case dbuscore.TypedValue:
		return typedValueToObject(vm, v)
	case godbus.ObjectPath:
		return string(v)
	case godbus.Signature:
		return v.String()
	case []any:
		out := make([]any, 0, len(v))
		for _, item := range v {
			out = append(out, typedPayloadToValue(vm, item))
		}
		return out
	case map[string]any:
		out := map[string]any{}
		for key, item := range v {
			out[key] = typedPayloadToValue(vm, item)
		}
		return out
	default:
		return v
	}
}

func exportTypedHelpers(vm *goja.Runtime, target *goja.Object) error {
	if err := target.Set("u32", func(call goja.FunctionCall) goja.Value {
		value, err := dbuscore.Uint32FromFloat64(call.Argument(0).ToFloat())
		if err != nil {
			panic(vm.NewGoError(err))
		}
		return typedValueToObject(vm, dbuscore.U32(value))
	}); err != nil {
		return fmt.Errorf("dbus: export u32: %w", err)
	}

	if err := target.Set("i32", func(call goja.FunctionCall) goja.Value {
		value, err := dbuscore.Int32FromFloat64(call.Argument(0).ToFloat())
		if err != nil {
			panic(vm.NewGoError(err))
		}
		return typedValueToObject(vm, dbuscore.I32(value))
	}); err != nil {
		return fmt.Errorf("dbus: export i32: %w", err)
	}

	if err := target.Set("path", func(raw string) goja.Value {
		value, err := dbuscore.ObjectPath(raw)
		if err != nil {
			panic(vm.NewGoError(err))
		}
		return typedValueToObject(vm, value)
	}); err != nil {
		return fmt.Errorf("dbus: export path: %w", err)
	}

	if err := target.Set("signature", func(raw string) goja.Value {
		value, err := dbuscore.Signature(raw)
		if err != nil {
			panic(vm.NewGoError(err))
		}
		return typedValueToObject(vm, value)
	}); err != nil {
		return fmt.Errorf("dbus: export signature: %w", err)
	}

	if err := target.Set("variant", func(call goja.FunctionCall) goja.Value {
		signature := call.Argument(0).String()
		inner, err := decodeJSValue(vm, call.Argument(1))
		if err != nil {
			panic(vm.NewGoError(err))
		}
		value, err := dbuscore.Variant(signature, inner)
		if err != nil {
			panic(vm.NewGoError(err))
		}
		return typedValueToObject(vm, value)
	}); err != nil {
		return fmt.Errorf("dbus: export variant: %w", err)
	}

	if err := target.Set("array", func(call goja.FunctionCall) goja.Value {
		signature := call.Argument(0).String()
		items, err := decodeJSValue(vm, call.Argument(1))
		if err != nil {
			panic(vm.NewGoError(err))
		}
		value, err := dbuscore.NewTypedValue(signature, items)
		if err != nil {
			panic(vm.NewGoError(err))
		}
		return typedValueToObject(vm, value)
	}); err != nil {
		return fmt.Errorf("dbus: export array: %w", err)
	}

	if err := target.Set("dict", func(call goja.FunctionCall) goja.Value {
		signature := call.Argument(0).String()
		items, err := decodeJSValue(vm, call.Argument(1))
		if err != nil {
			panic(vm.NewGoError(err))
		}
		value, err := dbuscore.NewTypedValue(signature, items)
		if err != nil {
			panic(vm.NewGoError(err))
		}
		return typedValueToObject(vm, value)
	}); err != nil {
		return fmt.Errorf("dbus: export dict: %w", err)
	}

	if err := target.Set("struct", func(call goja.FunctionCall) goja.Value {
		signature := call.Argument(0).String()
		items, err := decodeJSValue(vm, call.Argument(1))
		if err != nil {
			panic(vm.NewGoError(err))
		}
		value, err := dbuscore.NewTypedValue(signature, items)
		if err != nil {
			panic(vm.NewGoError(err))
		}
		return typedValueToObject(vm, value)
	}); err != nil {
		return fmt.Errorf("dbus: export struct: %w", err)
	}

	return nil
}
