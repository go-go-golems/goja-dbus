package dbusgoja

import (
	"context"
	"fmt"

	"github.com/dop251/goja"
	eventsmodule "github.com/go-go-golems/go-go-goja/modules/events"
	"github.com/go-go-golems/go-go-goja/pkg/runtimebridge"
	"github.com/go-go-golems/goja-dbus/pkg/dbuscore"
	godbus "github.com/godbus/dbus/v5"
)

type signalBuilder struct {
	vm       *goja.Runtime
	services runtimebridge.RuntimeServices
	bus      *dbuscore.Bus
	req      dbuscore.SignalMatchRequest
}

func newSignalBuilder(vm *goja.Runtime, services runtimebridge.RuntimeServices, bus *dbuscore.Bus, req dbuscore.SignalMatchRequest) goja.Value {
	builder := signalBuilder{vm: vm, services: services, bus: bus, req: req}
	return builder.toObject()
}

func (s signalBuilder) toObject() *goja.Object {
	obj := s.vm.NewObject()
	_ = obj.Set("sender", func(sender string) goja.Value {
		next := s
		next.req.Sender = sender
		return next.toObject()
	})
	_ = obj.Set("path", func(path string) goja.Value {
		objectPath := godbus.ObjectPath(path)
		if !objectPath.IsValid() {
			panic(s.vm.NewGoError(fmt.Errorf("dbus: invalid signal object path %q", path)))
		}
		next := s
		next.req.Path = objectPath
		return next.toObject()
	})
	_ = obj.Set("interface", func(iface string) goja.Value {
		next := s
		next.req.Interface = iface
		return next.toObject()
	})
	_ = obj.Set("member", func(member string) goja.Value {
		next := s
		next.req.Member = member
		return next.toObject()
	})
	_ = obj.Set("listen", func(call goja.FunctionCall) goja.Value {
		emitter, _, ok := eventsmodule.FromValue(call.Argument(0))
		if !ok {
			panic(s.vm.NewGoError(fmt.Errorf("dbus: listen requires an events.EventEmitter")))
		}
		req := s.req
		return promise(s.vm, s.services, "dbus.signals.listen", func(ctx context.Context) (any, error) {
			return s.bus.Listen(ctx, req, func(eventCtx context.Context, payload dbuscore.SignalPayload) error {
				return s.services.PostWithCustomContext(eventCtx, "dbus.signal.emit", func(_ context.Context, ownerVM *goja.Runtime) {
					_, err := emitter.Emit("signal", signalPayloadToValue(ownerVM, payload))
					if err != nil {
						_, _ = emitter.Emit("error", ownerVM.NewGoError(err))
					}
				})
			})
		}, func(vm *goja.Runtime, result any) (goja.Value, error) {
			sub, ok := result.(*dbuscore.Subscription)
			if !ok {
				return nil, fmt.Errorf("dbus: unexpected subscription result %T", result)
			}
			return subscriptionToValue(vm, s.services, sub), nil
		})
	})
	return obj
}

func signalPayloadToValue(vm *goja.Runtime, payload dbuscore.SignalPayload) goja.Value {
	obj := vm.NewObject()
	_ = obj.Set("sender", payload.Sender)
	_ = obj.Set("path", payload.Path)
	_ = obj.Set("name", payload.Name)
	_ = obj.Set("body", payload.Body)
	return obj
}

func subscriptionToValue(vm *goja.Runtime, services runtimebridge.RuntimeServices, sub *dbuscore.Subscription) goja.Value {
	obj := vm.NewObject()
	_ = obj.Set("close", func() goja.Value {
		return promise(vm, services, "dbus.signals.close", func(ctx context.Context) (any, error) {
			return nil, sub.Close(ctx)
		}, func(*goja.Runtime, any) (goja.Value, error) {
			return goja.Undefined(), nil
		})
	})
	return obj
}
