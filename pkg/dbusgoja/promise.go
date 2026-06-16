package dbusgoja

import (
	"context"

	"github.com/dop251/goja"
	"github.com/go-go-golems/go-go-goja/pkg/runtimebridge"
)

func promise(
	vm *goja.Runtime,
	services runtimebridge.RuntimeServices,
	op string,
	work func(context.Context) (any, error),
	toValue func(*goja.Runtime, any) (goja.Value, error),
) goja.Value {
	p, resolve, reject := vm.NewPromise()
	callCtx := runtimebridge.CurrentOwnerContext(vm)

	go func() {
		result, err := work(callCtx)
		_ = services.PostWithCustomContext(callCtx, op+".settle", func(_ context.Context, ownerVM *goja.Runtime) {
			if err != nil {
				_ = reject(dbusError(ownerVM, err))
				return
			}
			value, convErr := toValue(ownerVM, result)
			if convErr != nil {
				_ = reject(dbusError(ownerVM, convErr))
				return
			}
			_ = resolve(value)
		})
	}()

	return vm.ToValue(p)
}
