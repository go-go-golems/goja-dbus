package dbusgoja

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/go-go-golems/goja-dbus/pkg/dbuscore"
)

func decodeInputValue(vm *goja.Runtime, value goja.Value) (any, error) {
	if value == nil || goja.IsUndefined(value) || goja.IsNull(value) {
		return nil, nil
	}
	if obj := value.ToObject(vm); obj != nil && obj.Get(typedMarker).ToBoolean() {
		return decodeTypedValue(vm, obj)
	}
	return value.Export(), nil
}

func decodeTypedValue(vm *goja.Runtime, obj *goja.Object) (dbuscore.TypedValue, error) {
	signatureValue := obj.Get("signature")
	if goja.IsUndefined(signatureValue) || goja.IsNull(signatureValue) {
		return dbuscore.TypedValue{}, fmt.Errorf("dbus: typed value is missing signature")
	}
	signature := signatureValue.String()
	raw := obj.Get("value")
	var value any
	if rawObj := raw.ToObject(vm); rawObj != nil && rawObj.Get(typedMarker).ToBoolean() {
		inner, err := decodeTypedValue(vm, rawObj)
		if err != nil {
			return dbuscore.TypedValue{}, err
		}
		value = inner
	} else {
		value = raw.Export()
	}
	return dbuscore.TypedValue{Signature: signature, Value: value}, nil
}
