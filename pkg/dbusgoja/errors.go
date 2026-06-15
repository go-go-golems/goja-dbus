package dbusgoja

import "github.com/dop251/goja"

func dbusError(vm *goja.Runtime, err error) *goja.Object {
	obj := vm.NewGoError(err)
	_ = obj.Set("name", "DBusError")
	_ = obj.Set("code", "ERR_DBUS")
	return obj
}
