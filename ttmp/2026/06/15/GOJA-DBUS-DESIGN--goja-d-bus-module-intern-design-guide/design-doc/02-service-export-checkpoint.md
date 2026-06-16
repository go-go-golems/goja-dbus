---
Title: Service Export Checkpoint
Ticket: GOJA-DBUS-DESIGN
Status: active
Topics:
    - goja
    - dbus
    - design
DocType: design-doc
Intent: long-term
Owners: []
RelatedFiles: []
ExternalSources: []
Summary: "Checkpoint decision for deferring JavaScript-backed D-Bus service export until runtime cleanup and compound codecs are stronger."
LastUpdated: 2026-06-15T18:45:00-04:00
WhatFor: "Use this before implementing service export callbacks."
WhenToUse: "When deciding whether the next implementation step should be D-Bus service export or lifecycle/codec hardening."
---

# Service Export Checkpoint

## Executive Summary

Phase 5 should stop at a design checkpoint for now. The repository now has typed helpers, Promise-based client connections, scalar method calls, and EventEmitter-based signal subscriptions. JavaScript-backed D-Bus service export is the next large feature, but it crosses the hardest runtime boundary: external D-Bus clients call into Go on godbus-owned goroutines, and those calls must synchronously produce D-Bus replies while JavaScript handlers can only run on the go-go-goja runtime owner.

The recommendation is to defer service export implementation until two prerequisites are addressed:

- runtime-owned cleanup for open buses and signal subscriptions without relying only on explicit JavaScript `close()` calls;
- compound signature support for arrays, dictionaries, structs, and property dictionaries.

## Problem Statement

A JavaScript-backed exported service would let scripts write APIs such as:

```js
await bus
  .service("com.example.Echo1")
  .object("/com/example/Echo1")
  .interface("com.example.Echo1.Echo")
  .method("Echo")
    .in("s", "message")
    .out("s", "reply")
    .handle(({ message }) => `echo: ${message}`)
  .export();
```

This is useful, but it is riskier than client method calls and signal subscriptions because the direction reverses. Instead of JavaScript initiating a call and receiving a Promise, an external process initiates a D-Bus call and expects a timely D-Bus method return or D-Bus error. The Go adapter must invoke JavaScript on the owner thread, wait for the result, marshal it, and return it to godbus without deadlocking the owner or leaking goroutines.

## Current Implementation State

Implemented:

- `pkg/dbuscore.Bus` connection and `Call` support.
- Scalar codec for `s`, `u`, `i`, `o`, `g`, and `v`.
- JavaScript client builders through `destination().object().interface().method().call()`.
- Signal subscriptions through `bus.signals().listen(emitter)`.
- Explicit close handles for buses and subscriptions.

Not yet implemented:

- automatic runtime shutdown cleanup for D-Bus resources;
- arrays, structs, dictionaries, and `a{sv}` property dictionaries;
- D-Bus `org.freedesktop.DBus.Properties` export;
- D-Bus `org.freedesktop.DBus.Introspectable` export;
- JavaScript-backed service method dispatch.

## Options Considered

### Option A: Implement service export immediately with godbus method tables

This would use `ExportMethodTable` or `Export` and have Go callback methods bridge into JavaScript.

Pros:

- Fastest path to an echo-service demo.
- Uses existing godbus export APIs.

Cons:

- Callback signatures become awkward for dynamic JavaScript method specs.
- Waiting for JavaScript handler results from a D-Bus goroutine can deadlock if not designed carefully.
- Error mapping and Promise-returning handlers need a clear contract first.
- Properties and introspection would likely be bolted on later.

### Option B: Implement only a hard-coded echo service

This would prove external calls can reach JavaScript for one method.

Pros:

- Small demo.
- Useful for learning godbus export behavior.

Cons:

- Risks creating throwaway code that does not generalize.
- Does not solve lifecycle, introspection, or properties.
- Might encourage committing a compatibility shape that later needs replacement.

### Option C: Defer service export and harden lifecycle/codecs first

This keeps the current client/signal API stable and addresses foundations before exported services.

Pros:

- Avoids the most dangerous runtime boundary until the cleanup model is clear.
- Compound codecs are needed for realistic service APIs anyway.
- Keeps the implementation aligned with the design guide's warning that service callbacks are highest risk.

Cons:

- No echo service demo in this phase.
- Requires a future checkpoint before service export resumes.

## Decision

Choose **Option C: defer service export implementation for now**.

## Rationale

The current module is still a plain `modules.NativeModule`. That loader can access runtime services, but it does not receive `RuntimeModuleRegistrationContext.AddCloser`. Phase 4 therefore left a known lifecycle gap: signal subscriptions close correctly when JavaScript calls `close()`, but runtime shutdown cleanup is not automatic yet. Service export would add more long-lived resources, owned names, exported object paths, and JavaScript callbacks. Adding those before solving runtime cleanup would compound the gap.

The codec is also intentionally scalar. Real D-Bus services need compound values for common property dictionaries and signal payloads. Implementing service export before `a{sv}`, arrays, and structs would either make the service API too limited or force rushed codec work inside service export.

## Implementation Plan Before Reopening Service Export

1. Add a runtime-aware registration path or closer registry for this external module.
2. Track all opened buses and subscriptions in a runtime-scoped registry.
3. Close tracked resources automatically when the runtime closes.
4. Add compound signature support for arrays, dictionaries, structs, and variants containing compound values.
5. Add properties client support as an intermediate step before properties service export.
6. Revisit service export with a small design for callback dispatch, Promise-returning handlers, D-Bus error mapping, and timeout behavior.

## Future Service Export Pseudocode

```go
type ServiceSpec struct {
    Name    string
    Objects []ObjectSpec
}

type JSMethodDispatcher struct {
    services runtimebridge.RuntimeServices
    handler  goja.Callable
    timeout  time.Duration
}

func (d *JSMethodDispatcher) Dispatch(ctx context.Context, call MethodCall) (MethodReturn, error) {
    if d.timeout > 0 {
        var cancel context.CancelFunc
        ctx, cancel = context.WithTimeout(ctx, d.timeout)
        defer cancel()
    }

    ret, err := d.services.CallWithCustomContext(ctx, "dbus.service.dispatch", func(_ context.Context, vm *goja.Runtime) (any, error) {
        value, err := d.handler(goja.Undefined(), methodCallToValue(vm, call))
        if err != nil {
            return nil, err
        }
        return awaitPromiseIfNeeded(vm, value)
    })
    if err != nil {
        return MethodReturn{}, mapToDBusError(err)
    }
    return marshalMethodReturn(call.OutputSignature, ret)
}
```

## Status

Accepted for the current implementation pass: service export is deferred, and Phase 5 is represented by this checkpoint note.
