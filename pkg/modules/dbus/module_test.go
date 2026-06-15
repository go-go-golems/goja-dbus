package dbusmod_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/dop251/goja"
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

func TestDBusSystemConnectDeniedByDefault(t *testing.T) {
	rt := newDBusRuntime(t)

	_, err := rt.Owner.Call(context.Background(), "dbus.system.denied.setup", func(_ context.Context, vm *goja.Runtime) (any, error) {
		_, runErr := vm.RunString(`
globalThis.dbusState = { done: false, error: "", code: "" };
const dbus = require("dbus");
dbus.system().connect()
  .then(() => { globalThis.dbusState.done = true; })
  .catch((err) => { globalThis.dbusState.error = String(err); globalThis.dbusState.code = err.code || ""; });
`)
		return nil, runErr
	})
	if err != nil {
		t.Fatalf("setup: %v", err)
	}

	state := waitDBusState(t, rt)
	if state.Error == "" {
		t.Fatalf("expected system bus denial, state=%+v", state)
	}
	if state.Code != "ERR_DBUS" {
		t.Fatalf("error code = %q, want ERR_DBUS", state.Code)
	}
}

func TestDBusGetIdIntegration(t *testing.T) {
	if os.Getenv("GOJA_DBUS_INTEGRATION") != "1" {
		t.Skip("set GOJA_DBUS_INTEGRATION=1 to run against a real session bus")
	}
	rt := newDBusRuntime(t)

	_, err := rt.Owner.Call(context.Background(), "dbus.getid.setup", func(_ context.Context, vm *goja.Runtime) (any, error) {
		_, runErr := vm.RunString(`
globalThis.dbusState = { done: false, error: "", id: "" };
const dbus = require("dbus");
dbus.session().timeout(2000).connect()
  .then((bus) => bus
    .destination("org.freedesktop.DBus")
    .object("/org/freedesktop/DBus")
    .interface("org.freedesktop.DBus")
    .method("GetId")
    .out("s")
    .call()
    .then((id) => bus.close().then(() => id)))
  .then((id) => { globalThis.dbusState.done = true; globalThis.dbusState.id = String(id); })
  .catch((err) => { globalThis.dbusState.error = String(err); });
`)
		return nil, runErr
	})
	if err != nil {
		t.Fatalf("setup: %v", err)
	}

	state := waitDBusState(t, rt)
	if state.Error != "" {
		t.Fatalf("unexpected error: %s", state.Error)
	}
	if state.ID == "" {
		t.Fatalf("expected bus id, state=%+v", state)
	}
}

type dbusState struct {
	Done  bool   `json:"done"`
	Error string `json:"error"`
	ID    string `json:"id"`
	Code  string `json:"code"`
}

func newDBusRuntime(t *testing.T) *engine.Runtime {
	t.Helper()
	factory, err := engine.NewRuntimeFactoryBuilder().UseModuleMiddleware(engine.MiddlewareOnly("dbus")).Build()
	if err != nil {
		t.Fatalf("build factory: %v", err)
	}
	rt, err := factory.NewRuntime(engine.WithStartupContext(context.Background()), engine.WithLifetimeContext(context.Background()))
	if err != nil {
		t.Fatalf("new runtime: %v", err)
	}
	t.Cleanup(func() { _ = rt.Close(context.Background()) })
	return rt
}

func waitDBusState(t *testing.T, rt *engine.Runtime) dbusState {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	var state dbusState
	for time.Now().Before(deadline) {
		state = readDBusState(t, rt)
		if state.Done || state.Error != "" {
			return state
		}
		time.Sleep(10 * time.Millisecond)
	}
	return state
}

func readDBusState(t *testing.T, rt *engine.Runtime) dbusState {
	t.Helper()
	ret, err := rt.Owner.Call(context.Background(), "dbus.state.read", func(_ context.Context, vm *goja.Runtime) (any, error) {
		value, runErr := vm.RunString(`JSON.stringify(globalThis.dbusState || { done: false, error: "", id: "" })`)
		if runErr != nil {
			return nil, runErr
		}
		return value.String(), nil
	})
	if err != nil {
		t.Fatalf("read state: %v", err)
	}
	var state dbusState
	if raw, ok := ret.(string); ok && raw != "" {
		if err := json.Unmarshal([]byte(raw), &state); err != nil {
			t.Fatalf("unmarshal state: %v", err)
		}
	}
	return state
}
