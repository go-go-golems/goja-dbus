# Tasks

## Phase 0 — Planning, docs, and repo hygiene

- [x] Create docmgr ticket `GOJA-DBUS-DESIGN`.
- [x] Import `/tmp/dbus.md` into `sources/01-dbus.md`.
- [x] Write the intern-facing design and implementation guide.
- [x] Upload the guide bundle to reMarkable.
- [x] Add a phase/task execution plan to this ticket.
- [x] Commit the ticket documentation baseline (`981ab07`).

## Phase 1 — Repository normalization

- [x] Rename module path from `github.com/go-go-golems/XXX` to `github.com/go-go-golems/goja-dbus`.
- [x] Rename placeholder command `cmd/XXX` to `cmd/goja-dbus-demo`.
- [x] Update template package/logcopter names from `XXX` to `goja-dbus`.
- [x] Update Makefile release/install/logcopter paths.
- [x] Replace template README with a concise project overview.
- [x] Run `gofmt` and `GOWORK=off go test ./...`.
- [x] Commit repository normalization (`1d42a91`).

## Phase 2 — Native module skeleton and typed scalar helpers

- [x] Add `pkg/dbuscore` for pure Go policy and typed D-Bus values.
- [x] Add `pkg/dbusgoja` for Goja-facing exports and value conversion.
- [x] Add `pkg/modules/dbus` implementing `modules.NativeModule` and `modules.TypeScriptDeclarer`.
- [x] Add runtime integration test proving `require("dbus")` loads when the module package is blank-imported.
- [x] Add unit tests for scalar helpers (`u32`, `i32`, `path`, `signature`, `variant`).
- [x] Run `gofmt` and `GOWORK=off go test ./...`.
- [x] Commit module skeleton and helpers (`7927235`).

## Phase 3 — Session bus connect and method calls

- [x] Add `dbuscore.Bus` connection lifecycle for session/system/address connections.
- [x] Add Go-side `Policy` checks for connect and call operations.
- [x] Add immutable JavaScript builders for `session().timeout().policy().connect()`.
- [x] Add remote call builders for `destination().object().interface().method().in().out().call()`.
- [x] Marshal scalar inputs and unmarshal scalar/empty replies.
- [x] Add opt-in integration test for `org.freedesktop.DBus.GetId` guarded by `GOJA_DBUS_INTEGRATION=1`.
- [x] Run always-on tests and document integration-test command.
- [x] Commit connect/call support (`7ffee55`).

## Phase 4 — Signals and cleanup handles

- [x] Add D-Bus signal match request types in `dbuscore`.
- [x] Add subscription lifecycle with `AddMatchSignalContext`, `Signal`, `RemoveSignal`, and `RemoveMatchSignalContext`.
- [x] Expose EventEmitter-based `signals().listen(emitter)` in `dbusgoja`.
- [ ] Ensure subscription goroutines exit on runtime shutdown without requiring explicit bus/subscription close.
- [x] Add unit tests for match option construction and close idempotency where possible.
- [x] Commit signal support (`a2b3d5c`).

## Phase 5 — Service export design checkpoint

- [x] Re-read the design doc section on service export and decide whether to use method tables or lower-level godbus handlers.
- [x] Add a focused implementation note before coding service export.
- [ ] Implement the smallest echo-service export if the design checkpoint is clear. Deferred by `design-doc/02-service-export-checkpoint.md`.
- [ ] Add opt-in integration test for calling the exported echo service. Deferred with service export.
- [x] Commit service export or commit the checkpoint note if implementation is deferred (`c5acc52`).

## Phase 6 — Documentation, declarations, and hardening

- [ ] Expand README examples for GetId, notification, properties, signals, and service export status.
- [ ] Complete TypeScript declarations for public APIs.
- [ ] Add error-code mapping and policy denial tests.
- [ ] Run `GOWORK=off go test ./... -count=1` and targeted race tests for signal/service code.
- [ ] Update the diary, changelog, and final handoff notes.
- [ ] Commit final docs and hardening changes.
