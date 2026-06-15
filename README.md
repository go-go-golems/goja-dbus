# goja-dbus

`goja-dbus` is a native module for [go-go-goja](https://github.com/go-go-golems/go-go-goja) that exposes D-Bus client capabilities to JavaScript through `require("dbus")`.

The implementation goal is strict and boring on the Go side:

- JavaScript describes bus intent, destinations, object paths, interfaces, signatures, handlers, and lifecycle.
- Go owns D-Bus execution, marshaling, policy enforcement, cleanup, and Goja runtime scheduling.
- All JavaScript callbacks, `goja.Value` creation, EventEmitter delivery, and Promise settlement happen on the go-go-goja runtime owner.

See the docmgr ticket at `ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/` for the detailed intern-facing design and implementation guide.

## Current status

Implemented:

- `require("dbus")` native module registration.
- Typed helpers: `u32`, `i32`, `path`, `signature`, and `variant`.
- Promise-based session/system/address bus builders.
- Default-denied system bus policy.
- Remote method-call builders for scalar signatures.
- EventEmitter-based signal subscription builders.

Deferred:

- automatic runtime-shutdown cleanup for open D-Bus resources;
- compound D-Bus signatures such as arrays, dictionaries, and structs;
- JavaScript-backed D-Bus service export.

## JavaScript examples

### Call `org.freedesktop.DBus.GetId`

```js
const dbus = require("dbus");

const bus = await dbus.session().timeout(2000).connect();
try {
  const id = await bus
    .destination("org.freedesktop.DBus")
    .object("/org/freedesktop/DBus")
    .interface("org.freedesktop.DBus")
    .method("GetId")
    .out("s")
    .call();

  console.log("bus id:", id);
} finally {
  await bus.close();
}
```

### Use explicit typed values

```js
const dbus = require("dbus");

const count = dbus.u32(42);
const objectPath = dbus.path("/com/example/App1");
const options = dbus.variant("s", "hello");

console.log(count.signature, count.value);      // "u", 42
console.log(objectPath.signature, objectPath.value);
console.log(options.signature, options.value.signature);
```

### Listen for signals with EventEmitter

```js
const dbus = require("dbus");
const EventEmitter = require("events");

const bus = await dbus.session().connect();
const emitter = new EventEmitter();

emitter.on("signal", (signal) => {
  console.log(signal.sender, signal.path, signal.name, signal.body);
});

emitter.on("error", (err) => {
  console.error(err.code, err.message);
});

const sub = await bus
  .signals()
  .interface("org.freedesktop.DBus.Properties")
  .member("PropertiesChanged")
  .listen(emitter);

// Later:
await sub.close();
await bus.close();
```

## Go embedding sketch

External applications must blank-import the module package before creating a go-go-goja runtime so the module registers itself with the default registry.

```go
package main

import (
    "context"

    "github.com/go-go-golems/go-go-goja/pkg/engine"
    _ "github.com/go-go-golems/goja-dbus/pkg/modules/dbus"
)

func main() {
    factory, err := engine.NewRuntimeFactoryBuilder().
        UseModuleMiddleware(engine.MiddlewareOnly("dbus", "events")).
        Build()
    if err != nil {
        panic(err)
    }

    rt, err := factory.NewRuntime(
        engine.WithStartupContext(context.Background()),
        engine.WithLifetimeContext(context.Background()),
    )
    if err != nil {
        panic(err)
    }
    defer rt.Close(context.Background())

    // Run JavaScript that calls require("dbus").
}
```

## Development

```bash
GOWORK=off go test ./... -count=1
GOWORK=off go generate ./...
```

Run the opt-in real D-Bus integration test on a machine with a working session bus:

```bash
GOJA_DBUS_INTEGRATION=1 GOWORK=off go test ./pkg/modules/dbus -run TestDBusGetIdIntegration -count=1
```

Run the placeholder demo command:

```bash
go run ./cmd/goja-dbus-demo
```
