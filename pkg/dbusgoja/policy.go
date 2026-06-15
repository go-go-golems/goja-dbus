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
	if v := obj.Get("allowSystemBus"); !goja.IsUndefined(v) && !goja.IsNull(v) {
		policy.AllowSystemBus = v.ToBoolean()
	}
	if v := obj.Get("denySystemBus"); !goja.IsUndefined(v) && !goja.IsNull(v) && v.ToBoolean() {
		policy.AllowSystemBus = false
	}
	if v := obj.Get("allowSessionBus"); !goja.IsUndefined(v) && !goja.IsNull(v) {
		policy.AllowSessionBus = v.ToBoolean()
	}
	if v := obj.Get("allowCall"); !goja.IsUndefined(v) && !goja.IsNull(v) {
		arr := v.ToObject(vm)
		length := int(arr.Get("length").ToInteger())
		calls := make([]string, 0, length)
		for i := 0; i < length; i++ {
			item := arr.Get(fmt.Sprintf("%d", i))
			if !goja.IsUndefined(item) && !goja.IsNull(item) {
				calls = append(calls, item.String())
			}
		}
		policy.AllowCall = calls
	}
	return policy, nil
}
