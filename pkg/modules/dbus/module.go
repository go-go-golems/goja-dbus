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
The dbus module exposes explicit D-Bus typed value helpers.

Functions:
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
