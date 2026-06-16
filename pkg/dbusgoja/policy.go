package dbusgoja

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/go-go-golems/goja-dbus/pkg/dbuscore"
)

func decodePolicy(vm *goja.Runtime, value goja.Value, base dbuscore.Policy) (dbuscore.Policy, error) {
	if value == nil || goja.IsUndefined(value) || goja.IsNull(value) {
		return base, nil
	}
	obj := value.ToObject(vm)
	policy := base
	if v := obj.Get("allowSystemBus"); isPolicyValueSet(v) {
		policy.AllowSystemBus = v.ToBoolean()
	}
	if v := obj.Get("denySystemBus"); isPolicyValueSet(v) && v.ToBoolean() {
		policy.AllowSystemBus = false
	}
	if v := obj.Get("allowSessionBus"); isPolicyValueSet(v) {
		policy.AllowSessionBus = v.ToBoolean()
	}
	if v := obj.Get("allowAddressBus"); isPolicyValueSet(v) {
		policy.AllowAddressBus = v.ToBoolean()
	}
	if v := obj.Get("allowCall"); isPolicyValueSet(v) {
		arr := v.ToObject(vm)
		length := int(arr.Get("length").ToInteger())
		calls := make([]string, 0, length)
		for i := 0; i < length; i++ {
			item := arr.Get(fmt.Sprintf("%d", i))
			if isPolicyValueSet(item) {
				calls = append(calls, item.String())
			}
		}
		policy.AllowCall = calls
		policy.AllowCallSet = true
	}
	return policy, nil
}

func isPolicyValueSet(value goja.Value) bool {
	return value != nil && !goja.IsUndefined(value) && !goja.IsNull(value)
}
