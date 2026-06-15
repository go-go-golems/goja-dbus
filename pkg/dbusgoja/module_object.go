package dbusgoja

import (
	"fmt"

	"github.com/dop251/goja"
)

func NewModuleObject(vm *goja.Runtime) (*goja.Object, error) {
	if vm == nil {
		return nil, fmt.Errorf("dbus: nil runtime")
	}
	obj := vm.NewObject()
	if err := exportTypedHelpers(vm, obj); err != nil {
		return nil, err
	}
	return obj, nil
}
