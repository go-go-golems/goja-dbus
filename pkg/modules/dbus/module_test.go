package dbusmod_test

import (
	"context"
	"testing"

	"github.com/go-go-golems/go-go-goja/pkg/engine"
	_ "github.com/go-go-golems/goja-dbus/pkg/modules/dbus"
)

func TestRequireDBusTypedHelpers(t *testing.T) {
	factory, err := engine.NewRuntimeFactoryBuilder().UseModuleMiddleware(engine.MiddlewareOnly("dbus")).Build()
	if err != nil {
		t.Fatalf("build factory: %v", err)
	}
	rt, err := factory.NewRuntime(engine.WithStartupContext(context.Background()), engine.WithLifetimeContext(context.Background()))
	if err != nil {
		t.Fatalf("new runtime: %v", err)
	}
	defer func() { _ = rt.Close(context.Background()) }()

	_, err = rt.VM.RunString(`
const dbus = require("dbus");
const u = dbus.u32(42);
if (!u.__dbusTyped || u.signature !== "u" || u.value !== 42) throw new Error("bad u32");
const i = dbus.i32(-7);
if (!i.__dbusTyped || i.signature !== "i" || i.value !== -7) throw new Error("bad i32");
const p = dbus.path("/com/example/App1");
if (p.signature !== "o" || p.value !== "/com/example/App1") throw new Error("bad path");
const s = dbus.signature("a{sv}");
if (s.signature !== "g" || s.value !== "a{sv}") throw new Error("bad signature");
const v = dbus.variant("s", "hello");
if (v.signature !== "v" || v.value.signature !== "s" || v.value.value !== "hello") throw new Error("bad variant");
`)
	if err != nil {
		t.Fatalf("run script: %v", err)
	}
}

func TestRequireDBusHelperValidation(t *testing.T) {
	factory, err := engine.NewRuntimeFactoryBuilder().UseModuleMiddleware(engine.MiddlewareOnly("dbus")).Build()
	if err != nil {
		t.Fatalf("build factory: %v", err)
	}
	rt, err := factory.NewRuntime(engine.WithStartupContext(context.Background()), engine.WithLifetimeContext(context.Background()))
	if err != nil {
		t.Fatalf("new runtime: %v", err)
	}
	defer func() { _ = rt.Close(context.Background()) }()

	_, err = rt.VM.RunString(`
const dbus = require("dbus");
function mustThrow(fn, label) {
  let ok = false;
  try { fn(); } catch (err) { ok = true; }
  if (!ok) throw new Error("expected throw: " + label);
}
mustThrow(() => dbus.u32(-1), "u32 negative");
mustThrow(() => dbus.i32(1.5), "i32 fractional");
mustThrow(() => dbus.path("not/a/path"), "path");
mustThrow(() => dbus.signature("{"), "signature");
`)
	if err != nil {
		t.Fatalf("run script: %v", err)
	}
}
