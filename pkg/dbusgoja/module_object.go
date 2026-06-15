package dbusgoja

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/go-go-golems/go-go-goja/pkg/runtimebridge"
)

func NewModuleObject(vm *goja.Runtime) (*goja.Object, error) {
	if vm == nil {
		return nil, fmt.Errorf("dbus: nil runtime")
	}
	services, ok := runtimebridge.Lookup(vm)
	if !ok || services.Owner == nil {
		return nil, fmt.Errorf("dbus: module requires go-go-goja runtime services")
	}

	registry := newResourceRegistry(services.Lifetime())

	obj := vm.NewObject()
	if err := exportTypedHelpers(vm, obj); err != nil {
		return nil, err
	}
	if err := exportBusBuilders(vm, services, registry, obj); err != nil {
		return nil, err
	}
	return obj, nil
}
