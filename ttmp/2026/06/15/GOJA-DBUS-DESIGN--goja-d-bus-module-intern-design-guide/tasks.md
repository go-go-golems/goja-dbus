# Tasks

## Phase 0 — Planning, docs, and repo hygiene

- [x] Create docmgr ticket `GOJA-DBUS-DESIGN`.
- [x] Import `/tmp/dbus.md` into `sources/01-dbus.md`.
- [x] Write the intern-facing design and implementation guide.
- [x] Upload the guide bundle to reMarkable.
- [x] Add a phase/task execution plan to this ticket.
- [ ] Commit the ticket documentation baseline.

## Phase 1 — Repository normalization

- [ ] Rename module path from `github.com/go-go-golems/XXX` to `github.com/go-go-golems/goja-dbus`.
- [ ] Rename placeholder command `cmd/XXX` to `cmd/goja-dbus-demo`.
- [ ] Update template package/logcopter names from `XXX` to `goja-dbus`.
- [ ] Update Makefile release/install/logcopter paths.
- [ ] Replace template README with a concise project overview.
- [ ] Run `gofmt` and `GOWORK=off go test ./...`.
- [ ] Commit repository normalization.

## Phase 2 — Native module skeleton and typed scalar helpers

- [ ] Add `pkg/dbuscore` for pure Go policy and typed D-Bus values.
- [ ] Add `pkg/dbusgoja` for Goja-facing exports and value conversion.
- [ ] Add `pkg/modules/dbus` implementing `modules.NativeModule` and `modules.TypeScriptDeclarer`.
- [ ] Add runtime integration test proving `require("dbus")` loads when the module package is blank-imported.
- [ ] Add unit tests for scalar helpers (`u32`, `i32`, `path`, `signature`, `variant`).
- [ ] Run `gofmt` and `GOWORK=off go test ./...`.
- [ ] Commit module skeleton and helpers.

## Phase 3 — Session bus connect and method calls

- [ ] Add `dbuscore.Bus` connection lifecycle for session/system/address connections.
- [ ] Add Go-side `Policy` checks for connect and call operations.
- [ ] Add immutable JavaScript builders for `session().timeout().policy().connect()`.
- [ ] Add remote call builders for `destination().object().interface().method().in().out().call()`.
- [ ] Marshal scalar inputs and unmarshal scalar/empty replies.
- [ ] Add opt-in integration test for `org.freedesktop.DBus.GetId` guarded by `GOJA_DBUS_INTEGRATION=1`.
- [ ] Run always-on tests and document integration-test command.
- [ ] Commit connect/call support.

## Phase 4 — Signals and cleanup handles

- [ ] Add D-Bus signal match request types in `dbuscore`.
- [ ] Add subscription lifecycle with `AddMatchSignalContext`, `Signal`, `RemoveSignal`, and `RemoveMatchSignalContext`.
- [ ] Expose EventEmitter-based `signals().listen(emitter)` in `dbusgoja`.
- [ ] Ensure subscription goroutines exit on close and runtime shutdown.
- [ ] Add unit tests for match option construction and close idempotency where possible.
- [ ] Commit signal support.

## Phase 5 — Service export design checkpoint

- [ ] Re-read the design doc section on service export and decide whether to use method tables or lower-level godbus handlers.
- [ ] Add a focused implementation note before coding service export.
- [ ] Implement the smallest echo-service export if the design checkpoint is clear.
- [ ] Add opt-in integration test for calling the exported echo service.
- [ ] Commit service export or commit the checkpoint note if implementation is deferred.

## Phase 6 — Documentation, declarations, and hardening

- [ ] Expand README examples for GetId, notification, properties, signals, and service export status.
- [ ] Complete TypeScript declarations for public APIs.
- [ ] Add error-code mapping and policy denial tests.
- [ ] Run `GOWORK=off go test ./... -count=1` and targeted race tests for signal/service code.
- [ ] Update the diary, changelog, and final handoff notes.
- [ ] Commit final docs and hardening changes.
