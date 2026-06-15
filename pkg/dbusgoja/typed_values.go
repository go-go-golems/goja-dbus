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
	switch v := value.Value.(type) {
	case dbuscore.TypedValue:
		_ = obj.Set("value", typedValueToObject(vm, v))
	case godbus.ObjectPath:
		_ = obj.Set("value", string(v))
	case godbus.Signature:
		_ = obj.Set("value", v.String())
	default:
		_ = obj.Set("value", v)
	}
	return obj
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
		value, err := dbuscore.Variant(signature, call.Argument(1).Export())
		if err != nil {
			panic(vm.NewGoError(err))
		}
		return typedValueToObject(vm, value)
	}); err != nil {
		return fmt.Errorf("dbus: export variant: %w", err)
	}

	return nil
}
