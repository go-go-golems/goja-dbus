package dbusgoja

import (
	"context"
	"fmt"
	"time"

	"github.com/dop251/goja"
	"github.com/go-go-golems/go-go-goja/pkg/runtimebridge"
	"github.com/go-go-golems/goja-dbus/pkg/dbuscore"
	godbus "github.com/godbus/dbus/v5"
)

type busBuilder struct {
	vm       *goja.Runtime
	services runtimebridge.RuntimeServices
	opts     dbuscore.ConnectOptions
}

func newBusBuilder(vm *goja.Runtime, services runtimebridge.RuntimeServices, opts dbuscore.ConnectOptions) *goja.Object {
	builder := busBuilder{vm: vm, services: services, opts: opts}
	return builder.toObject()
}

func (b busBuilder) toObject() *goja.Object {
	obj := b.vm.NewObject()
	_ = obj.Set("timeout", func(ms int64) goja.Value {
		next := b
		next.opts.Timeout = time.Duration(ms) * time.Millisecond
		return next.toObject()
	})
	_ = obj.Set("policy", func(call goja.FunctionCall) goja.Value {
		next := b
		policy, err := decodePolicy(b.vm, call.Argument(0), next.opts.Policy)
		if err != nil {
			panic(b.vm.NewGoError(err))
		}
		next.opts.Policy = policy
		return next.toObject()
	})
	_ = obj.Set("connect", func() goja.Value {
		return promise(b.vm, b.services, "dbus.connect", func(ctx context.Context) (any, error) {
			return dbuscore.Connect(ctx, b.opts)
		}, func(vm *goja.Runtime, result any) (goja.Value, error) {
			bus, ok := result.(*dbuscore.Bus)
			if !ok {
				return nil, fmt.Errorf("dbus: unexpected connect result %T", result)
			}
			return newBusObject(vm, b.services, bus), nil
		})
	})
	return obj
}

func newBusObject(vm *goja.Runtime, services runtimebridge.RuntimeServices, bus *dbuscore.Bus) goja.Value {
	obj := vm.NewObject()
	_ = obj.Set("close", func() goja.Value {
		return promise(vm, services, "dbus.bus.close", func(ctx context.Context) (any, error) {
			return nil, bus.Close(ctx)
		}, func(*goja.Runtime, any) (goja.Value, error) {
			return goja.Undefined(), nil
		})
	})
	_ = obj.Set("destination", func(destination string) goja.Value {
		return newDestinationObject(vm, services, bus, destination)
	})
	_ = obj.Set("signals", func() goja.Value {
		return newSignalBuilder(vm, services, bus, dbuscore.SignalMatchRequest{})
	})
	return obj
}

func newDestinationObject(vm *goja.Runtime, services runtimebridge.RuntimeServices, bus *dbuscore.Bus, destination string) goja.Value {
	obj := vm.NewObject()
	_ = obj.Set("object", func(path string) goja.Value {
		objectPath := godbus.ObjectPath(path)
		if !objectPath.IsValid() {
			panic(vm.NewGoError(fmt.Errorf("dbus: invalid object path %q", path)))
		}
		return newRemoteObject(vm, services, bus, destination, objectPath)
	})
	return obj
}

func newRemoteObject(vm *goja.Runtime, services runtimebridge.RuntimeServices, bus *dbuscore.Bus, destination string, path godbus.ObjectPath) goja.Value {
	obj := vm.NewObject()
	_ = obj.Set("interface", func(iface string) goja.Value {
		return newInterfaceObject(vm, services, bus, destination, path, iface)
	})
	return obj
}

func newInterfaceObject(vm *goja.Runtime, services runtimebridge.RuntimeServices, bus *dbuscore.Bus, destination string, path godbus.ObjectPath, iface string) goja.Value {
	obj := vm.NewObject()
	_ = obj.Set("method", func(member string) goja.Value {
		return newMethodBuilder(vm, services, bus, dbuscore.MethodCallRequest{
			Destination: destination,
			Path:        path,
			Interface:   iface,
			Member:      member,
		})
	})
	return obj
}

type methodBuilder struct {
	vm       *goja.Runtime
	services runtimebridge.RuntimeServices
	bus      *dbuscore.Bus
	req      dbuscore.MethodCallRequest
}

func newMethodBuilder(vm *goja.Runtime, services runtimebridge.RuntimeServices, bus *dbuscore.Bus, req dbuscore.MethodCallRequest) goja.Value {
	builder := methodBuilder{vm: vm, services: services, bus: bus, req: req}
	return builder.toObject()
}

func (m methodBuilder) toObject() *goja.Object {
	obj := m.vm.NewObject()
	_ = obj.Set("in", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			panic(m.vm.NewGoError(fmt.Errorf("dbus: method.in requires signature and value")))
		}
		value, err := decodeInputValue(m.vm, call.Argument(1))
		if err != nil {
			panic(m.vm.NewGoError(err))
		}
		next := m
		next.req.Inputs = append(append([]dbuscore.Arg(nil), m.req.Inputs...), dbuscore.Arg{
			Signature: call.Argument(0).String(),
			Value:     value,
		})
		return next.toObject()
	})
	_ = obj.Set("out", func(signature string) goja.Value {
		next := m
		next.req.OutputSignature = signature
		return next.toObject()
	})
	_ = obj.Set("timeout", func(ms int64) goja.Value {
		next := m
		next.req.Timeout = time.Duration(ms) * time.Millisecond
		return next.toObject()
	})
	_ = obj.Set("call", func() goja.Value {
		req := m.req
		return promise(m.vm, m.services, "dbus.method.call", func(ctx context.Context) (any, error) {
			return m.bus.Call(ctx, req)
		}, func(vm *goja.Runtime, result any) (goja.Value, error) {
			return vm.ToValue(result), nil
		})
	})
	return obj
}
