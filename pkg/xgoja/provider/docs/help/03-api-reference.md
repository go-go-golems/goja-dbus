---
Title: "goja-dbus API reference"
Slug: "api-reference"
Short: "Reference for the current JavaScript API exposed by require(\"dbus\")."
Topics:
- goja
- dbus
- api
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

This reference lists the current JavaScript API exposed by the goja-dbus native module. It describes the implemented subset, not the full future design.

The module is available as:

```js
const dbus = require("dbus");
```

## Top-level builders

### `dbus.session(): BusBuilder`

Creates a builder for the session bus. This is the default safe bus for examples.

### `dbus.system(): BusBuilder`

Creates a builder for the system bus. The default policy denies system bus access, so `dbus.system().connect()` rejects unless policy support is explicitly widened in a future host configuration.

### `dbus.connect(address: string): BusBuilder`

Creates a builder for an explicit D-Bus address. Empty addresses are rejected. Explicit-address connections are denied by the default policy; opt in with `policy({ allowAddressBus: true })` only when the host intentionally allows direct addresses.

## `BusBuilder`

### `timeout(ms: number): BusBuilder`

Returns a new builder with a default timeout for connect-created bus operations.

### `policy(policy: DBusPolicy): BusBuilder`

Returns a new builder with policy options applied. Current fields are:

```ts
interface DBusPolicy {
  allowSessionBus?: boolean;
  allowSystemBus?: boolean;
  denySystemBus?: boolean;
  allowAddressBus?: boolean;
  allowCall?: string[];
}
```

`allowAddressBus` is separate from session/system policy so callers cannot bypass the default system-bus denial by passing a raw system bus address to `connect(address)`. `allowCall` supports exact matches and suffix `*` prefix matches. An explicitly empty `allowCall: []` list denies all method calls; omit the field to keep the builder's existing call policy.

### `connect(): Promise<DBusBus>`

Connects to the selected bus and resolves to a bus object. Rejections use JavaScript Error objects with:

```ts
err.name === "DBusError"
err.code === "ERR_DBUS"
```

## `DBusBus`

### `close(): Promise<void>`

Closes the bus and tracked subscriptions. The runtime also closes tracked buses when the runtime lifetime ends.

### `destination(name: string): RemoteDestination`

Selects a remote D-Bus destination.

### `signals(): SignalBuilder`

Creates a signal subscription builder.

## Remote method builders

```ts
bus
  .destination("org.freedesktop.DBus")
  .object("/org/freedesktop/DBus")
  .interface("org.freedesktop.DBus")
  .method("GetId")
  .out("s")
  .call();
```

### `RemoteDestination.object(path: string): RemoteObject`

Selects an object path. Invalid object paths throw immediately.

### `RemoteObject.interface(name: string): RemoteInterface`

Selects a D-Bus interface.

### `RemoteInterface.method(name: string): MethodCallBuilder`

Selects a method member.

### `MethodCallBuilder.in(signature: string, value: any): MethodCallBuilder`

Adds one input argument. Use typed helpers for ambiguous values.

### `MethodCallBuilder.out(signature: string): MethodCallBuilder`

Declares the expected output signature. The current unmarshal path validates the signature and returns a single value, an array of values, or `null`/`undefined` equivalent when the reply body is empty.

### `MethodCallBuilder.timeout(ms: number): MethodCallBuilder`

Overrides the bus default timeout for this method call.

### `MethodCallBuilder.call(): Promise<any>`

Executes the D-Bus method call.

## Typed value helpers

```ts
interface DBusTypedValue {
  readonly __dbusTyped: true;
  readonly signature: string;
  readonly value: any;
}
```

### Scalar helpers

- `dbus.u32(value: number): DBusTypedValue`
- `dbus.i32(value: number): DBusTypedValue`
- `dbus.path(value: string): DBusTypedValue`
- `dbus.signature(value: string): DBusTypedValue`
- `dbus.variant(signature: string, value: any): DBusTypedValue`

### Compound helpers

- `dbus.array(signature: string, values: any[]): DBusTypedValue`
- `dbus.dict(signature: string, values: Record<string, any>): DBusTypedValue`
- `dbus.struct(signature: string, values: any[]): DBusTypedValue`

The current codec supports common arrays (`as`, `au`, `ai`, `ao`, `av`), `a{sv}`, and flat structs such as `(su)`. Unsupported signatures fail explicitly.

## Signals

```ts
interface SignalPayload {
  sender: string;
  path: string;
  name: string;
  body: any[];
}

interface SignalSubscription {
  close(): Promise<void>;
}
```

### `SignalBuilder.sender(name: string): SignalBuilder`

Adds a sender match.

### `SignalBuilder.path(path: string): SignalBuilder`

Adds an object path match. Invalid object paths throw immediately.

### `SignalBuilder.interface(name: string): SignalBuilder`

Adds an interface match.

### `SignalBuilder.member(name: string): SignalBuilder`

Adds a signal member match.

### `SignalBuilder.listen(emitter: EventEmitter): Promise<SignalSubscription>`

Installs the match rule and emits matching signals through `emitter.emit("signal", payload)`. The emitter must come from `require("events")`.

## Troubleshooting

| Problem | Cause | Solution |
| --- | --- | --- |
| `dbus.variant("v", value)` is confusing | Variants should usually wrap a concrete inner signature, not another generic variant | Use a concrete signature such as `s`, `u`, or `a{sv}` |
| `a{sv}` rejects values | Each map value must be a D-Bus variant helper | Wrap values with `dbus.variant(signature, value)` |
| Struct marshaling fails | Only flat struct signatures are currently supported | Keep structs simple or add codec support before using nested structs |
| System bus is denied | Default policy blocks it | Use session bus or implement reviewed host policy widening |

## See Also

- `xgoja help getting-started`
- `xgoja help user-guide`
- `xgoja help xgoja-v2-reference`
