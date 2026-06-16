---
Title: Imported D-Bus API Proposal
Ticket: GOJA-DBUS-DESIGN
Status: active
Topics:
    - goja
    - dbus
    - design
DocType: reference
Intent: long-term
Owners: []
RelatedFiles: []
ExternalSources:
    - /tmp/dbus.md
Summary: "Imported source proposal for a Goja D-Bus JavaScript API and Go-side architecture."
LastUpdated: 2026-06-15T17:50:00-04:00
WhatFor: "Reference the user-provided proposal that shaped the intern-facing design guide."
WhenToUse: "When comparing implementation choices against the original requested D-Bus API direction."
---

A good shape is: **JS defines intent; Go owns execution**. JS should describe bus, destination, path, interface, signatures, handlers, and lifecycle. Go should enforce policy, marshal D-Bus types, isolate the Goja runtime, and schedule all callbacks back onto one JS event loop.

Goja is a good fit for this because it lets Go pass values into JS and export JS values back to Go, but a `goja.Runtime` is not goroutine-safe, and object values cannot be passed between runtimes. It also does not provide browser/Node timers by itself; the host must provide an event loop, or use `goja_nodejs/eventloop`. ([GitHub][1]) D-Bus itself is naturally asynchronous IPC, and `godbus/dbus/v5` supports connecting to buses, calling methods, emitting/receiving signals, and exporting Go values as D-Bus services. ([Go Packages][2])

## The JS API I would want

```js
const bus = await dbus
  .session()
  .timeout(5000)
  .policy({
    allowOwn: ["com.example.*"],
    allowCall: ["org.freedesktop.*", "org.bluez.*"],
    allowSystemBus: false
  })
  .connect();
```

Core objects:

```js
dbus.session() / dbus.system()
dbus.connect(address)

bus.destination(name)
bus.object(path)
bus.interface(name)

remote.method(name)
  .in(signature, value)
  .out(signature)
  .timeout(ms)
  .call()

bus.signals()
  .sender(name)
  .path(path)
  .interface(name)
  .member(name)
  .on(fn)
  .listen()

bus.service(name)
  .object(path)
  .interface(name)
  .method(name)
  .property(name, signature)
  .signal(name)
  .export()
```

Type helpers matter because JS has only one `Number`, while D-Bus distinguishes `int32`, `uint32`, `int64`, `double`, `variant`, `object path`, signatures, arrays, structs, and dictionaries. `godbus/dbus/v5` has conversion rules for Go primitives, arrays, maps, structs, variants, object paths, signatures, and Unix FDs, but the JS side should still make ambiguous types explicit. ([Go Packages][2])

```js
dbus.u32(42)
dbus.i32(-1)
dbus.u64("18446744073709551615")
dbus.path("/com/example/App1")
dbus.signature("a{sv}")
dbus.variant("s", "hello")
dbus.dict("sv", { urgency: dbus.variant("y", 1) })
dbus.array("s", ["one", "two"])
dbus.struct("ssu", ["name", "state", dbus.u32(7)])
```

## Example 1: call a simple method

```js
const bus = await dbus.session().connect();

const name = await bus
  .destination("org.freedesktop.DBus")
  .object("/org/freedesktop/DBus")
  .interface("org.freedesktop.DBus")
  .method("GetId")
  .out("s")
  .call();

console.log("bus id:", name);
```

## Example 2: desktop notification

```js
const bus = await dbus.session().connect();

const id = await bus
  .destination("org.freedesktop.Notifications")
  .object("/org/freedesktop/Notifications")
  .interface("org.freedesktop.Notifications")
  .method("Notify")
  .in("s", "goja-dbus-demo")          // app_name
  .in("u", dbus.u32(0))               // replaces_id
  .in("s", "")                        // app_icon
  .in("s", "Hello from Goja")         // summary
  .in("s", "This came from a JS script hosted inside Go.")
  .in("as", [])                       // actions
  .in("a{sv}", {
    urgency: dbus.variant("y", 1)
  })
  .in("i", dbus.i32(3000))            // timeout ms
  .out("u")
  .call();

console.log("notification id:", id);
```

## Example 3: reuse a remote interface

```js
const bus = await dbus.session().connect();

const notifications = bus
  .destination("org.freedesktop.Notifications")
  .object("/org/freedesktop/Notifications")
  .interface("org.freedesktop.Notifications");

const info = await notifications
  .method("GetServerInformation")
  .out("(ssss)")
  .call();

console.log({
  name: info[0],
  vendor: info[1],
  version: info[2],
  specVersion: info[3],
});
```

## Example 4: read a property

Expose the standard properties API as a convenience. D-Bus API guidelines recommend using `org.freedesktop.DBus.Properties` for object state instead of custom `GetX`/`SetX` methods. ([dbus.freedesktop.org][3])

```js
const bus = await dbus.system().connect();

const hostname = await bus
  .destination("org.freedesktop.hostname1")
  .object("/org/freedesktop/hostname1")
  .properties("org.freedesktop.hostname1")
  .get("Hostname", "s");

console.log("hostname:", hostname);
```

Equivalent raw call:

```js
const hostname = await bus
  .destination("org.freedesktop.hostname1")
  .object("/org/freedesktop/hostname1")
  .interface("org.freedesktop.DBus.Properties")
  .method("Get")
  .in("s", "org.freedesktop.hostname1")
  .in("s", "Hostname")
  .out("v")
  .call();
```

## Example 5: listen to a signal

D-Bus signals should be scheduled into the JS event loop, never called directly from a Go D-Bus goroutine.

```js
const bus = await dbus.system().connect();

const sub = await bus
  .signals()
  .sender("org.freedesktop.login1")
  .path("/org/freedesktop/login1")
  .interface("org.freedesktop.login1.Manager")
  .member("PrepareForSleep")
  .on((signal) => {
    const goingToSleep = signal.body[0];
    console.log("sleep state changed:", goingToSleep);
  })
  .listen();

process.onExit(() => sub.close());
```

A more ergonomic signal body decoder:

```js
await bus
  .signals()
  .interface("org.freedesktop.DBus.Properties")
  .member("PropertiesChanged")
  .decode("(sa{sv}as)", ([iface, changed, invalidated], meta) => {
    console.log("changed:", meta.path, iface, changed, invalidated);
  })
  .listen();
```

## Example 6: expose an echo service

D-Bus public APIs should version service names, interface names, and object paths from the start. The freedesktop guidelines explicitly recommend version numbers in service name, interface name, and object path so incompatible versions can coexist. ([dbus.freedesktop.org][3])

```js
const bus = await dbus.session().connect();

await bus
  .service("com.example.Echo1")
  .object("/com/example/Echo1")
  .interface("com.example.Echo1.Echo")
  .method("Echo")
    .in("s", "message")
    .out("s", "reply")
    .handle(({ message }) => {
      return `echo: ${message}`;
    })
  .export();

console.log("Echo service exported");
```

Client:

```js
const bus = await dbus.session().connect();

const reply = await bus
  .destination("com.example.Echo1")
  .object("/com/example/Echo1")
  .interface("com.example.Echo1.Echo")
  .method("Echo")
  .in("s", "hi")
  .out("s")
  .call();

console.log(reply);
```

## Example 7: expose a counter with properties and signals

```js
let count = 0;

const bus = await dbus.session().connect();

const counter = await bus
  .service("com.example.Counter1")
  .object("/com/example/Counter1")
  .interface("com.example.Counter1.Counter")
  .property("Count", "u")
    .read(() => dbus.u32(count))
    .emitsChanged()
  .method("Add")
    .in("u", "delta")
    .out("u", "count")
    .handle(({ delta }) => {
      count += Number(delta);
      counter.properties.set("Count", dbus.u32(count));
      counter.emit("Changed", dbus.u32(count));
      return dbus.u32(count);
    })
  .method("Reset")
    .out("u", "count")
    .handle(() => {
      count = 0;
      counter.properties.set("Count", dbus.u32(count));
      counter.emit("Changed", dbus.u32(count));
      return dbus.u32(count);
    })
  .signal("Changed")
    .arg("u", "count")
  .export();
```

The property API should map to `org.freedesktop.DBus.Properties`. `godbus/dbus/v5/prop` provides a `Properties` helper for implementing that interface, including property change behavior. ([Go Packages][4])

## Example 8: expose multiple objects

```js
const bus = await dbus.session().connect();

const service = bus.service("com.example.Sensors1");

function exposeSensor(id, initialValue) {
  let value = initialValue;

  return service
    .object(`/com/example/Sensors1/${id}`)
    .interface("com.example.Sensors1.Sensor")
    .property("Id", "s")
      .read(() => id)
      .constant()
    .property("Value", "d")
      .read(() => value)
      .emitsChanged()
    .method("Read")
      .out("d", "value")
      .handle(() => value)
    .method("SetFakeValue")
      .in("d", "value")
      .handle(({ value: next }) => {
        value = Number(next);
        this.properties.set("Value", value);
        this.emit("ValueChanged", value);
      })
    .signal("ValueChanged")
      .arg("d", "value")
    .export();
}

await exposeSensor("cpu", 47.2);
await exposeSensor("gpu", 50.1);
await exposeSensor("battery", 91.0);
```

For a real variable-size object tree, add `ObjectManager` support. The D-Bus guidelines recommend `org.freedesktop.DBus.ObjectManager` when clients are expected to care about most or all objects in a tree. ([dbus.freedesktop.org][3])

Possible DSL:

```js
await bus
  .service("com.example.Sensors1")
  .objectManager("/com/example/Sensors1")
  .export();
```

## Example 9: proxy one D-Bus service into another

```js
const bus = await dbus.session().connect();

const source = bus
  .destination("org.freedesktop.Notifications")
  .object("/org/freedesktop/Notifications")
  .interface("org.freedesktop.Notifications");

await bus
  .service("com.example.NotificationProxy1")
  .object("/com/example/NotificationProxy1")
  .interface("com.example.NotificationProxy1.Proxy")
  .method("NotifySimple")
    .in("s", "title")
    .in("s", "body")
    .out("u", "id")
    .handle(async ({ title, body }) => {
      return await source
        .method("Notify")
        .in("s", "notification-proxy")
        .in("u", dbus.u32(0))
        .in("s", "")
        .in("s", title)
        .in("s", body)
        .in("as", [])
        .in("a{sv}", {})
        .in("i", dbus.i32(3000))
        .out("u")
        .call();
    })
  .export();
```

## Example 10: policy-limited automation script

```js
const bus = await dbus
  .session()
  .timeout(2000)
  .policy({
    allowCall: [
      "org.freedesktop.Notifications",
      "org.mpris.MediaPlayer2.*"
    ],
    denySystemBus: true,
    denyOwn: ["*"]
  })
  .connect();

await bus
  .destination("org.freedesktop.Notifications")
  .object("/org/freedesktop/Notifications")
  .interface("org.freedesktop.Notifications")
  .method("Notify")
  .in("s", "restricted-script")
  .in("u", dbus.u32(0))
  .in("s", "")
  .in("s", "Sandboxed")
  .in("s", "This script can notify, but cannot own names or use system bus.")
  .in("as", [])
  .in("a{sv}", {})
  .in("i", dbus.i32(2000))
  .out("u")
  .call();
```

## Go-side architecture

Use one Goja runtime per script/plugin, with one event loop. Do not call JS callbacks directly from D-Bus goroutines.

```go
type RuntimeHost struct {
    vm     *goja.Runtime
    loop   *eventloop.EventLoop
    buses  *BusRegistry
    policy Policy
}

type Bus struct {
    conn   *dbus.Conn
    kind   BusKind
    host   *RuntimeHost
}

type RemoteBuilder struct {
    bus         *Bus
    destination string
    path        dbus.ObjectPath
    iface       string
}

type MethodCallBuilder struct {
    remote  *RemoteBuilder
    member  string
    inputs  []Arg
    output  string
    timeout time.Duration
}

type SignalBuilder struct {
    bus     *Bus
    match   MatchRule
    handler goja.Callable
}

type ServiceBuilder struct {
    bus     *Bus
    name    string
    objects []*ObjectBuilder
}
```

Promise helper pattern:

```go
func (h *RuntimeHost) promise(
    rt *goja.Runtime,
    work func(ctx context.Context) (any, error),
) goja.Value {
    p, resolve, reject := rt.NewPromise()

    go func() {
        result, err := work(context.Background())

        h.loop.RunOnLoop(func(rt *goja.Runtime) {
            if err != nil {
                _ = reject(rt.NewGoError(err))
                return
            }
            _ = resolve(result)
        })
    }()

    return rt.ToValue(p)
}
```

For method calls:

```go
func (m *MethodCallBuilder) Call(rt *goja.Runtime) goja.Value {
    return m.remote.bus.host.promise(rt, func(ctx context.Context) (any, error) {
        obj := m.remote.bus.conn.Object(
            m.remote.destination,
            m.remote.path,
        )

        method := m.remote.iface + "." + m.member

        call := obj.CallWithContext(
            ctx,
            method,
            0,
            marshalArgs(m.inputs)...,
        )

        if call.Err != nil {
            return nil, call.Err
        }

        return unmarshalReply(m.output, call.Body)
    })
}
```

For signal delivery:

```go
func (s *SignalBuilder) dispatch(sig *dbus.Signal) {
    payload := convertSignal(sig)

    s.bus.host.loop.RunOnLoop(func(rt *goja.Runtime) {
        _, err := s.handler(
            goja.Undefined(),
            rt.ToValue(payload),
        )
        if err != nil {
            // Route to host logger / unhandled callback handler.
        }
    })
}
```

## API design rules I would enforce

1. **Everything async returns a Promise.** D-Bus IPC can fail, time out, or be delayed; D-Bus guidelines emphasize that method calls are asynchronous and clients must handle errors. ([dbus.freedesktop.org][3])

2. **Signatures are explicit at boundaries.** JS numbers are ambiguous. Make callers write `"u"`, `"i"`, `"x"`, `"t"`, `"d"`, `"v"`, `"a{sv}"`, etc.

3. **Builders are immutable until execution.** This makes policies, tracing, dry-run introspection, and error messages much easier.

4. **Callbacks are always scheduled onto the Goja loop.** A Goja runtime can only be used by one goroutine at a time. ([GitHub][1])

5. **Expose standard D-Bus interfaces automatically.** At minimum: `Introspectable`, `Properties`, optionally `ObjectManager`.

6. **Sandbox policy is Go-side, not JS-side.** JS should not be able to bypass bus restrictions by mutating the API object.

7. **Support cancellation and cleanup.** Every `.listen()` and `.export()` should return a handle with `.close()`.

8. **Use versioned names.** Example: `com.example.Counter1`, `com.example.Counter1.Counter`, `/com/example/Counter1`.

## Nice final developer feel

The ideal script should read like this:

```js
const bus = await dbus.session().connect();

const player = bus
  .destination("org.mpris.MediaPlayer2.spotify")
  .object("/org/mpris/MediaPlayer2")
  .interface("org.mpris.MediaPlayer2.Player");

await player.method("PlayPause").call();

await bus
  .signals()
  .sender("org.mpris.MediaPlayer2.spotify")
  .interface("org.freedesktop.DBus.Properties")
  .member("PropertiesChanged")
  .decode("(sa{sv}as", ([iface, changed]) => {
    if (iface === "org.mpris.MediaPlayer2.Player") {
      console.log("player changed:", changed);
    }
  })
  .listen();
```

The Go implementation should feel strict and boring underneath: one runtime, one event loop, explicit type conversion, explicit policy, no cross-goroutine JS values, and D-Bus errors mapped cleanly into rejected promises.

[1]: https://github.com/dop251/goja "GitHub - dop251/goja: ECMAScript/JavaScript engine in pure Go · GitHub"
[2]: https://pkg.go.dev/github.com/godbus/dbus/v5 "dbus package - github.com/godbus/dbus/v5 - Go Packages"
[3]: https://dbus.freedesktop.org/doc/dbus-api-design.html "D-Bus API Design Guidelines"
[4]: https://pkg.go.dev/github.com/godbus/dbus/v5/prop "prop package - github.com/godbus/dbus/v5/prop - Go Packages"

