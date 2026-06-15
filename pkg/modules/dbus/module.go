package dbusmod

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/go-go-golems/go-go-goja/modules"
	"github.com/go-go-golems/go-go-goja/pkg/tsgen/spec"
	"github.com/go-go-golems/goja-dbus/pkg/dbusgoja"
)

type module struct{}

var _ modules.NativeModule = (*module)(nil)
var _ modules.TypeScriptDeclarer = (*module)(nil)

func (m *module) Name() string { return "dbus" }

func (m *module) Doc() string {
	return `
The dbus module exposes D-Bus connection builders and explicit typed value helpers.

Functions:
  session(): Creates a session-bus builder.
  system(): Creates a system-bus builder. Denied by default policy unless enabled.
  connect(address): Creates an explicit-address bus builder.
  u32(value): Wraps an unsigned 32-bit integer.
  i32(value): Wraps a signed 32-bit integer.
  path(value): Validates and wraps a D-Bus object path.
  signature(value): Validates and wraps a D-Bus signature.
  variant(signature, value): Wraps a value as a D-Bus variant.
`
}

func (m *module) TypeScriptModule() *spec.Module {
	return &spec.Module{
		Name:        "dbus",
		Description: "D-Bus client, service, and typed value helpers.",
		RawDTS: []string{
			"interface DBusTypedValue {",
			"  readonly __dbusTyped: true;",
			"  readonly signature: string;",
			"  readonly value: any;",
			"}",
			"interface DBusPolicy {",
			"  allowSessionBus?: boolean;",
			"  allowSystemBus?: boolean;",
			"  denySystemBus?: boolean;",
			"  allowCall?: string[];",
			"}",
			"interface BusBuilder {",
			"  timeout(ms: number): BusBuilder;",
			"  policy(policy: DBusPolicy): BusBuilder;",
			"  connect(): Promise<DBusBus>;",
			"}",
			"interface DBusBus {",
			"  close(): Promise<void>;",
			"  destination(name: string): RemoteDestination;",
			"}",
			"interface RemoteDestination { object(path: string): RemoteObject; }",
			"interface RemoteObject { interface(name: string): RemoteInterface; }",
			"interface RemoteInterface { method(name: string): MethodCallBuilder; }",
			"interface MethodCallBuilder {",
			"  in(signature: string, value: any): MethodCallBuilder;",
			"  out(signature: string): MethodCallBuilder;",
			"  timeout(ms: number): MethodCallBuilder;",
			"  call(): Promise<any>;",
			"}",
			"export function session(): BusBuilder;",
			"export function system(): BusBuilder;",
			"export function connect(address: string): BusBuilder;",
			"export function u32(value: number): DBusTypedValue;",
			"export function i32(value: number): DBusTypedValue;",
			"export function path(value: string): DBusTypedValue;",
			"export function signature(value: string): DBusTypedValue;",
			"export function variant(signature: string, value: any): DBusTypedValue;",
		},
	}
}

func (m *module) Loader(vm *goja.Runtime, moduleObj *goja.Object) {
	exports := moduleObj.Get("exports").(*goja.Object)
	obj, err := dbusgoja.NewModuleObject(vm)
	if err != nil {
		panic(vm.NewGoError(fmt.Errorf("dbus module: create exports: %w", err)))
	}
	for _, key := range obj.Keys() {
		modules.SetExport(exports, m.Name(), key, obj.Get(key))
	}
}

func init() {
	modules.Register(&module{})
}
