package dbusgoja

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/go-go-golems/goja-dbus/pkg/dbuscore"
)

func decodeInputValue(vm *goja.Runtime, value goja.Value) (any, error) {
	return decodeJSValue(vm, value)
}

func decodeTypedValue(vm *goja.Runtime, obj *goja.Object) (dbuscore.TypedValue, error) {
	signatureValue := obj.Get("signature")
	if goja.IsUndefined(signatureValue) || goja.IsNull(signatureValue) {
		return dbuscore.TypedValue{}, fmt.Errorf("dbus: typed value is missing signature")
	}
	signature := signatureValue.String()
	value, err := decodeJSValue(vm, obj.Get("value"))
	if err != nil {
		return dbuscore.TypedValue{}, err
	}
	return dbuscore.TypedValue{Signature: signature, Value: value}, nil
}

func decodeJSValue(vm *goja.Runtime, value goja.Value) (any, error) {
	if value == nil || goja.IsUndefined(value) || goja.IsNull(value) {
		return nil, nil
	}
	obj := value.ToObject(vm)
	if obj != nil {
		marker := obj.Get(typedMarker)
		if marker != nil && !goja.IsUndefined(marker) && marker.ToBoolean() {
			return decodeTypedValue(vm, obj)
		}
	}
	if obj != nil && obj.ClassName() == "Array" {
		lengthValue := obj.Get("length")
		if lengthValue != nil && !goja.IsUndefined(lengthValue) {
			length := int(lengthValue.ToInteger())
			if length >= 0 {
				out := make([]any, 0, length)
				for i := 0; i < length; i++ {
					item, err := decodeJSValue(vm, obj.Get(fmt.Sprintf("%d", i)))
					if err != nil {
						return nil, err
					}
					out = append(out, item)
				}
				return out, nil
			}
		}
	}
	if obj != nil && obj.ClassName() == "Object" {
		keys := obj.Keys()
		if len(keys) > 0 {
			out := map[string]any{}
			for _, key := range keys {
				item, err := decodeJSValue(vm, obj.Get(key))
				if err != nil {
					return nil, err
				}
				out[key] = item
			}
			return out, nil
		}
	}
	return value.Export(), nil
}
