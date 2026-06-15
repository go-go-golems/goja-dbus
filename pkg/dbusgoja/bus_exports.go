package dbusgoja

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/go-go-golems/go-go-goja/pkg/runtimebridge"
	"github.com/go-go-golems/goja-dbus/pkg/dbuscore"
)

func exportBusBuilders(vm *goja.Runtime, services runtimebridge.RuntimeServices, registry *resourceRegistry, target *goja.Object) error {
	if err := target.Set("session", func() goja.Value {
		return newBusBuilder(vm, services, registry, dbuscore.ConnectOptions{
			Kind:   dbuscore.BusSession,
			Policy: dbuscore.DefaultPolicy(),
		})
	}); err != nil {
		return fmt.Errorf("dbus: export session: %w", err)
	}

	if err := target.Set("system", func() goja.Value {
		return newBusBuilder(vm, services, registry, dbuscore.ConnectOptions{
			Kind:   dbuscore.BusSystem,
			Policy: dbuscore.DefaultPolicy(),
		})
	}); err != nil {
		return fmt.Errorf("dbus: export system: %w", err)
	}

	if err := target.Set("connect", func(address string) goja.Value {
		return newBusBuilder(vm, services, registry, dbuscore.ConnectOptions{
			Kind:    dbuscore.BusAddress,
			Address: address,
			Policy:  dbuscore.DefaultPolicy(),
		})
	}); err != nil {
		return fmt.Errorf("dbus: export connect: %w", err)
	}

	return nil
}
