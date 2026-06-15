---
Title: "goja-dbus user guide"
Slug: "user-guide"
Short: "How to use the goja-dbus module from JavaScript and bundled jsverbs."
Topics:
- goja
- dbus
- javascript
Commands:
- eval
- run
- verbs
Flags: []
IsTopLevel: true
IsTemplate: false
ShowPerDefault: true
SectionType: GeneralTopic
---

This guide describes how JavaScript code should use the `dbus` module. The module follows one rule: JavaScript describes D-Bus intent, and Go owns execution, policy checks, type marshaling, cleanup, and Goja runtime scheduling.

The current implementation supports client-side method calls, explicit typed values, EventEmitter-based signal subscriptions, and runtime-owned cleanup. JavaScript-backed D-Bus service export remains deferred until the codec and callback-dispatch model are stronger.

## Import the module

Use CommonJS `require` inside the generated xgoja runtime:

```js
const dbus = require("dbus");
```

The xgoja spec selects the `dbus` runtime module from the goja-dbus provider, so the generated binary can resolve this import in `eval`, `run`, and bundled jsverbs.

## Use explicit typed values

D-Bus has a precise type system, while JavaScript has broad `Number`, array, and object values. Use helper wrappers when the D-Bus signature matters:

```js
const count = dbus.u32(7);
const path = dbus.path("/com/example/App1");
const hint = dbus.variant("s", "hello");
const names = dbus.array("as", ["one", "two"]);
const hints = dbus.dict("a{sv}", { urgency: dbus.variant("u", dbus.u32(1)) });
const pair = dbus.struct("(su)", ["count", dbus.u32(7)]);
```

Typed helper objects contain a `signature` and `value` field. They are primarily for the adapter and examples; application code should not mutate them.

## Connect to a bus

Use a builder to choose the bus and set optional policy or timeout values:

```js
const bus = await dbus.session().timeout(2000).connect();
```

The default policy allows the session bus and denies the system bus. This keeps safe examples runnable without accidentally exposing privileged host APIs.

## Call a remote method

Build calls by selecting a destination, object path, interface, and method. Inputs and outputs use D-Bus signatures:

```js
const id = await bus
  .destination("org.freedesktop.DBus")
  .object("/org/freedesktop/DBus")
  .interface("org.freedesktop.DBus")
  .method("GetId")
  .out("s")
  .call();
```

Always close buses when you are done:

```js
await bus.close();
```

The runtime also tracks buses and closes them when the runtime lifetime ends, but explicit close keeps scripts predictable.

## Listen for signals

Signal subscriptions use the Go-native EventEmitter from `require("events")`. This keeps listener registration in JavaScript while Go schedules delivery back onto the runtime owner:

```js
const EventEmitter = require("events");
const emitter = new EventEmitter();

emitter.on("signal", (signal) => {
  console.log(signal.sender, signal.path, signal.name, signal.body);
});

const sub = await bus
  .signals()
  .interface("org.freedesktop.DBus.Properties")
  .member("PropertiesChanged")
  .listen(emitter);

await sub.close();
```

The payload has `sender`, `path`, `name`, and `body` fields. Body values are normalized where the adapter has explicit support.

## Bundled example verbs

The generated binary includes jsverbs that demonstrate safe behavior:

```bash
./dist/goja-dbus verbs examples typed-values
./dist/goja-dbus verbs examples denied-system-bus
./dist/goja-dbus verbs examples get-id-script
```

Use these examples to verify that the binary contains the module and docs before trying host-specific D-Bus calls.

## Troubleshooting

| Problem | Cause | Solution |
| --- | --- | --- |
| A value marshals with the wrong type | JavaScript value did not carry an explicit D-Bus signature | Use `dbus.u32`, `dbus.variant`, `dbus.array`, `dbus.dict`, or `dbus.struct` |
| A Promise rejects with `ERR_DBUS` | Policy, connection, marshaling, or D-Bus call failed | Inspect `err.message`; start with session-bus examples |
| Signal listener does not run | No matching D-Bus signal arrived or the emitter was not an `events.EventEmitter` | Verify the match filters and use `const EventEmitter = require("events")` |
| Runtime exits with open resources | Explicit close was skipped | Current runtime cleanup closes tracked buses on shutdown, but scripts should still call `close()` for deterministic behavior |

## See Also

- `xgoja help getting-started`
- `xgoja help api-reference`
- `xgoja help jsverbs-example-overview`
