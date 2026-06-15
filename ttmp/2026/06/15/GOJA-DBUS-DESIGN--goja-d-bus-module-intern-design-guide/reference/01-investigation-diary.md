---
Title: Investigation Diary
Ticket: GOJA-DBUS-DESIGN
Status: active
Topics:
    - goja
    - dbus
    - design
DocType: reference
Intent: long-term
Owners: []
RelatedFiles:
    - Path: goja-dbus/ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/design-doc/01-goja-d-bus-module-intern-design-and-implementation-guide.md
      Note: Primary deliverable produced during the investigation
    - Path: goja-dbus/ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/sources/01-dbus.md
      Note: Imported source material referenced by the diary
ExternalSources:
    - ../sources/01-dbus.md
Summary: Chronological notes for creating the Goja D-Bus design ticket and intern guide.
LastUpdated: 2026-06-15T17:45:00-04:00
WhatFor: Resume or review the investigation behind the Goja D-Bus module design guide.
WhenToUse: When continuing implementation, reviewing evidence, or checking what commands and failures occurred during ticket creation.
---


# Diary

## Goal

Capture the creation of the `GOJA-DBUS-DESIGN` docmgr ticket, the import of `/tmp/dbus.md`, the evidence gathered from `./go-go-goja/`, and the resulting intern-facing design and implementation guide.

## Step 1: Created the ticket, imported the source, and wrote the intern guide

I created a new docmgr ticket for the Goja D-Bus module design work, copied the provided `/tmp/dbus.md` into the ticket's `sources/` directory, and used the existing `go-go-goja` documentation and code as evidence for the design. The resulting design doc explains the runtime model, D-Bus concepts, JavaScript API shape, Go package layout, async/signal/service-export flows, pseudocode, diagrams, decision records, implementation phases, and validation strategy.

The main architectural conclusion is that this module should follow the rule from the imported source: JavaScript defines intent, while Go owns execution. That maps cleanly to go-go-goja's existing native-module, runtimebridge, and connected-EventEmitter patterns.

### Prompt Context

**User prompt (verbatim):**

```text
Create a new docmgr ticket and import /tmp/dbus.md into its sources directory and then Create  a detailed analysis / design / implementation guide that is for a new intern, explaining all the parts of the system needed to understand what it is, with prose paragraphs and bullet point sand pseudocode and diagrams and api references and file references. It should be very clear and technical. Store in the ticket and the nupload to remarkable.


Look at ./go-go-goja/ for all the docs needed.
```

**Assistant interpretation:** Create a docmgr ticket, preserve `/tmp/dbus.md` as a ticket source, research `./go-go-goja/`, write a detailed intern-oriented design/implementation guide in the ticket, then upload the resulting document bundle to reMarkable.

**Inferred user intent:** Produce durable, reviewable, and portable documentation that an intern can use to implement a go-go-goja D-Bus native module safely and idiomatically.

**Commit (code):** N/A — documentation-only work in a docmgr ticket; no Git commit was requested or created.

### What I did

- Ran `docmgr status --summary-only` and confirmed there were no existing tickets in this workspace.
- Created ticket `GOJA-DBUS-DESIGN` with title `Goja D-Bus Module Intern Design Guide`.
- Added the primary design document: `design-doc/01-goja-d-bus-module-intern-design-and-implementation-guide.md`.
- Added this diary document: `reference/01-investigation-diary.md`.
- Copied `/tmp/dbus.md` to `goja-dbus/ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/sources/01-dbus.md`.
- Read key go-go-goja documentation and source files:
  - `go-go-goja/pkg/doc/01-introduction.md`
  - `go-go-goja/pkg/doc/02-creating-modules.md`
  - `go-go-goja/pkg/doc/03-async-patterns.md`
  - `go-go-goja/pkg/doc/16-nodejs-primitives.md`
  - `go-go-goja/pkg/doc/17-connected-eventemitters-developer-guide.md`
  - `go-go-goja/modules/common.go`
  - `go-go-goja/pkg/engine/factory.go`
  - `go-go-goja/pkg/engine/runtime.go`
  - `go-go-goja/pkg/runtimebridge/runtimebridge.go`
  - `go-go-goja/pkg/jsevents/manager.go`
  - `go-go-goja/pkg/jsevents/fswatch.go`
- Inspected the current `goja-dbus` template state via `goja-dbus/go.mod` and `goja-dbus/README.md`.
- Used a temporary Go module under `/tmp` to inspect `github.com/godbus/dbus/v5` API documentation with `go doc`.
- Wrote the primary design guide with Mermaid diagrams, pseudocode, API references, file references, decision records, and phased implementation guidance.
- Ran `docmgr doctor --ticket GOJA-DBUS-DESIGN --stale-after 30` and fixed vocabulary/source frontmatter issues until it passed.
- Uploaded the ticket bundle to reMarkable at `/ai/2026/06/15/GOJA-DBUS-DESIGN` with document name `GOJA DBUS DESIGN GUIDE`.

### Why

- The imported `/tmp/dbus.md` contained the desired JavaScript API and a first-pass architecture, but it needed to be converted into a durable implementation guide tied to the actual go-go-goja runtime model.
- `go-go-goja` has strict runtime-owner rules that materially affect every D-Bus method call, signal callback, Promise settlement, and service export.
- D-Bus is host-access IPC, so the design needed explicit policy and lifecycle guidance rather than a simple binding sketch.

### What worked

- `docmgr ticket create-ticket` created the full ticket workspace, including `sources/`, `design-doc/`, `reference/`, `tasks.md`, and `changelog.md`.
- `/tmp/dbus.md` copied cleanly into the ticket's `sources/` directory.
- The `go-go-goja` docs provided strong evidence for the native module contract, Promise scheduling pattern, runtime contexts, connected EventEmitter pattern, and sandboxing model.
- A temporary Go module successfully fetched `github.com/godbus/dbus/v5` and allowed `go doc` inspection of the core D-Bus APIs.
- The final reMarkable upload succeeded with: `OK: uploaded GOJA DBUS DESIGN GUIDE.pdf -> /ai/2026/06/15/GOJA-DBUS-DESIGN`.

### What didn't work

- Kagi web search failed because the local surf socket was unavailable:
  - Command/tool: `kagi_web_search("github.com/godbus/dbus/v5 Go package CallWithContext ExportSubtree Emit Signal MatchOption pkg.go.dev prop Properties")`
  - Error: `dial unix /home/manuel/snap/chromium/common/surf-cli/surf.sock: connect: no such file or directory`
- Running `go doc` from the workspace initially failed because `go.work` is out of sync with module Go versions:
  - Command: `cd goja-dbus && go doc github.com/godbus/dbus/v5.Conn`
  - Error included: `go: module ../go-go-goja listed in go.work file requires go >= 1.26.1, but go.work lists go 1.25; to update it: go work use`
- `GOWORK=off go doc github.com/godbus/dbus/v5.Conn` from `goja-dbus` still failed because the dependency is not yet in the module:
  - Error: `doc: cannot find package "github.com/godbus/dbus/v5.Conn"`
- A first temporary `go doc` loop exited early when asking for a non-existent symbol `WithMatchArg0`; rerunning for `prop` and `introspect` separately succeeded.
- The first real reMarkable upload attempt failed during PDF rendering:
  - Mermaid error: `Parse error on line 2 ... require("dbus") ... got 'PS'`
  - Pandoc/LaTeX error: `Undefined control sequence` at the prompt text that contained literal backslash-n sequences.
  - Fix applied: simplified the Mermaid edge label and replaced the literal backslash-n prompt text with a fenced text block containing real line breaks.

### What I learned

- The current `goja-dbus` repository is still a Go template and must be normalized before implementation.
- go-go-goja already has most of the runtime primitives needed for D-Bus: `NativeModule`, runtime services, owner-thread scheduling, runtime closers, and connected EventEmitters.
- D-Bus signals map well onto the connected EventEmitter pattern because they are long-lived Go-side resources that must deliver events to JavaScript safely.
- The design should require explicit D-Bus signatures and typed helpers because JavaScript values are not precise enough for D-Bus marshaling.

### What was tricky to build

- The main tricky part was aligning the imported source's raw event-loop pseudocode with the actual go-go-goja runtimebridge pattern. The source used `loop.RunOnLoop`, but the repository docs recommend `runtimebridge.RuntimeServices` because it preserves owner-thread safety and context/lifetime behavior.
- Another subtle area was service export: D-Bus calls arrive on Go/godbus goroutines, but JavaScript handlers must run on the Goja owner. The guide explicitly calls this out as the highest-risk runtime boundary and recommends implementing method export only after client calls and signals are stable.
- The workspace's `go.work` version mismatch made direct `go doc` inspection fail, so I used a temporary module under `/tmp` to inspect godbus APIs without changing repository files.

### What warrants a second pair of eyes

- The proposed service export callback model should be reviewed by someone familiar with both `godbus` export internals and go-go-goja owner scheduling.
- The policy model should be reviewed before implementation to ensure system-bus and service-name ownership defaults are safe.
- The codec plan should be reviewed for D-Bus signature edge cases, especially variants, dictionaries, structs, object paths, and 64-bit integer handling.

### What should be done in the future

- Normalize `goja-dbus/go.mod` away from the template module path.
- Decide whether the module will live as a separate repository or inside `go-go-goja/modules`.
- Implement Phase 1 from the design doc: module skeleton, `require("dbus")`, type helpers, and a runtime integration test.
- Fix or regenerate the workspace `go.work` if future commands need to run across all three modules.

### Code review instructions

- Start with the primary design doc:
  - `/home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/design-doc/01-goja-d-bus-module-intern-design-and-implementation-guide.md`
- Compare its API shape against the imported source:
  - `/home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/sources/01-dbus.md`
- Validate that the design is consistent with:
  - `go-go-goja/pkg/doc/03-async-patterns.md`
  - `go-go-goja/pkg/doc/17-connected-eventemitters-developer-guide.md`
  - `go-go-goja/pkg/runtimebridge/runtimebridge.go`
- Run `docmgr doctor --ticket GOJA-DBUS-DESIGN --stale-after 30` after any future edits.

### Technical details

Important commands run during this step:

```bash
docmgr status --summary-only
docmgr ticket create-ticket --ticket GOJA-DBUS-DESIGN --title "Goja D-Bus Module Intern Design Guide" --topics goja,dbus,design
docmgr doc add --ticket GOJA-DBUS-DESIGN --doc-type design-doc --title "Goja D-Bus Module Intern Design and Implementation Guide"
docmgr doc add --ticket GOJA-DBUS-DESIGN --doc-type reference --title "Investigation Diary"
cp /tmp/dbus.md goja-dbus/ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/sources/01-dbus.md
docmgr doctor --ticket GOJA-DBUS-DESIGN --stale-after 30
remarquee upload bundle ... --name "GOJA DBUS DESIGN GUIDE" --remote-dir "/ai/2026/06/15/GOJA-DBUS-DESIGN" --toc-depth 2 --non-interactive
```

Temporary godbus API inspection used this pattern:

```bash
TMP=$(mktemp -d /tmp/go-doc-dbus-XXXX)
cd "$TMP"
go mod init tmpdbusdoc
go get github.com/godbus/dbus/v5@latest
go doc github.com/godbus/dbus/v5.Conn
go doc github.com/godbus/dbus/v5.Object
go doc github.com/godbus/dbus/v5/prop
go doc github.com/godbus/dbus/v5/introspect
```

## Step 2: Phase 1 Repository Normalization

I normalized the repository so future implementation work no longer carries the template identity. This changed the Go module path, demo command path, release metadata, logcopter area names, and README so the project now consistently presents itself as `goja-dbus`.

This step intentionally did not add the D-Bus module yet. It creates a clean baseline for later commits so module code, tests, and documentation refer to the correct import path and command paths from the start.

### Prompt Context

**User prompt (verbatim):** "Create detailed phasesa nd for each phase, tasks, and then execute them one by one, and commit atappropriate intervals, and keep a detailed d iary."

**Assistant interpretation:** Convert the design guide into concrete implementation phases, update ticket tasks, then execute each phase sequentially with focused commits and diary updates.

**Inferred user intent:** Move from planning into implementation while preserving reviewable commits and a detailed continuation trail.

**Commit (code):** `1d42a9151ffba42e35614d297131e09955c8c3ba` — "Normalize goja-dbus repository"

### What I did

- Updated `goja-dbus/go.mod` from `github.com/go-go-golems/XXX` to `github.com/go-go-golems/goja-dbus`.
- Renamed `cmd/XXX/main.go` to `cmd/goja-dbus-demo/main.go`.
- Updated `Makefile` paths for logcopter, release module path, install binary name, and build target.
- Updated `.goreleaser.yaml` project/build/binary/homepage/package metadata from `XXX` to `goja-dbus` / `goja-dbus-demo`.
- Updated `logcopter_generate.go` and `pkg/logcopter.go` area names.
- Replaced the template README with a concise project overview and development commands.
- Ran `gofmt` on changed Go files.
- Ran `GOWORK=off go mod tidy`.
- Ran `GOWORK=off go test ./...`.
- Committed the code changes.

### Why

- The repository was still using a template module path and command name, which would make later package imports and tests misleading.
- Normalizing the repository before adding the module keeps later diffs focused on actual D-Bus functionality.

### What worked

- `GOWORK=off go test ./...` passed after the rename:
  - `? github.com/go-go-golems/goja-dbus [no test files]`
  - `? github.com/go-go-golems/goja-dbus/cmd/goja-dbus-demo [no test files]`
  - `? github.com/go-go-golems/goja-dbus/pkg [no test files]`
- Git detected the command rename as a clean rename from `cmd/XXX/main.go` to `cmd/goja-dbus-demo/main.go`.

### What didn't work

- My first staging command tried to add the old path after `git mv`:
  - Command: `git add .goreleaser.yaml Makefile README.md go.mod logcopter_generate.go pkg/logcopter.go cmd/goja-dbus-demo/main.go cmd/XXX/main.go`
  - Error: `fatal: pathspec 'cmd/XXX/main.go' did not match any files`
- I corrected this by staging the `cmd` directory instead:
  - Command: `git add .goreleaser.yaml Makefile README.md go.mod logcopter_generate.go pkg/logcopter.go cmd`

### What I learned

- The template had placeholder references not only in `go.mod` and README, but also in release, Makefile, and logcopter configuration.
- `AGENT.md` still contains generic examples using `XXX/YYY/FOOO`, but those are instructional placeholders rather than project identity references, so I left them unchanged.

### What was tricky to build

- The root `logcopter_generate.go` package name could not use the hyphenated module name. I used `package gojadbus` while keeping the module path and logcopter area as `goja-dbus`.
- The release binary name and module name are intentionally different: the repository/module is `goja-dbus`, while the placeholder executable became `goja-dbus-demo` because the real deliverable is a native module, not primarily a CLI.

### What warrants a second pair of eyes

- The `.goreleaser.yaml` metadata is mechanically normalized but may need product decisions before an actual release.
- The install target uses `which goja-dbus-demo`; if the binary is not installed yet, that target may still need a better default destination in a future release-readiness pass.

### What should be done in the future

- Add the real module packages under `pkg/dbuscore`, `pkg/dbusgoja`, and `pkg/modules/dbus`.
- Revisit `.goreleaser.yaml` once the project has a real CLI or decide to remove CLI release packaging if this remains library-only.

### Code review instructions

- Start with `go.mod` to confirm the module path.
- Review `Makefile`, `.goreleaser.yaml`, and `logcopter_generate.go` for remaining template references.
- Validate with:
  - `GOWORK=off go test ./...`
  - `rg -n 'XXX|go-go-golems/XXX|go-go-golems\.XXX' -S . --glob '!ttmp/**' --glob '!go.sum'`

### Technical details

Commands:

```bash
cd goja-dbus
git mv cmd/XXX/main.go cmd/goja-dbus-demo/main.go
GOWORK=off go mod tidy
GOWORK=off go test ./...
git commit -m "Normalize goja-dbus repository"
```

## Step 3: Phase 2 Native Module Skeleton and Typed Helpers

I added the first real `dbus` native module implementation. This phase does not connect to D-Bus yet; it establishes the package boundaries, module registration, TypeScript declaration hook, JavaScript export object, and explicit typed-value helpers that later method-call and property code will consume.

The key implementation choice was to keep D-Bus type validation in `pkg/dbuscore` and JavaScript object construction in `pkg/dbusgoja`. That preserves the design-doc boundary where core D-Bus logic remains testable without Goja, while the adapter owns lowerCamel JavaScript object shape and Goja exceptions.

### Prompt Context

**User prompt (verbatim):** (same as Step 2)

**Assistant interpretation:** Continue executing the documented implementation phases with focused code commits and diary updates.

**Inferred user intent:** Build the D-Bus module incrementally while keeping each phase reviewable and validated.

**Commit (code):** `7927235e43d6c9246dd9fa6ef6193433b5e497db` — "Add dbus native module skeleton"

### What I did

- Added `pkg/dbuscore/types.go` with `TypedValue`, scalar constructors, object path validation, signature validation, variant wrapping, and integer-bound helpers.
- Added `pkg/dbuscore/policy.go` with an intentionally small Phase 2 policy placeholder.
- Added `pkg/dbusgoja/typed_values.go` to expose JavaScript helpers:
  - `dbus.u32(value)`
  - `dbus.i32(value)`
  - `dbus.path(value)`
  - `dbus.signature(value)`
  - `dbus.variant(signature, value)`
- Added `pkg/dbusgoja/module_object.go` to construct the JavaScript export object.
- Added `pkg/modules/dbus/module.go` implementing `modules.NativeModule` and `modules.TypeScriptDeclarer`.
- Added unit tests in `pkg/dbuscore/types_test.go`.
- Added runtime integration tests in `pkg/modules/dbus/module_test.go` that blank-import the module and `require("dbus")` through go-go-goja.
- Added dependencies on `go-go-goja`, `goja`, and `godbus/dbus/v5`; added a local replace to `../go-go-goja` for workspace development.
- Ran `gofmt` and `GOWORK=off go test ./... -count=1`.

### Why

- Later D-Bus calls need explicit typed values because JavaScript numbers and objects do not carry enough D-Bus type information.
- The module needs to prove it can register with go-go-goja before adding connection and method-call complexity.
- The TypeScript declaration hook should exist from the start so the public API stays documented as it grows.

### What worked

- The module loads via `require("dbus")` when `pkg/modules/dbus` is blank-imported before runtime creation.
- JavaScript helper objects expose lowerCamel properties: `__dbusTyped`, `signature`, and `value`.
- `GOWORK=off go test ./... -count=1` passed after fixing the object path/signature JavaScript presentation issue.

### What didn't work

- The first `TestRequireDBusTypedHelpers` run failed:
  - Command: `GOWORK=off go test ./... -count=1`
  - Error: `run script: Error: bad path at <eval>:8:67(69)`
- Cause: `dbuscore.ObjectPath` stored a `godbus.ObjectPath`, and the Goja object exposed that Go-specific value directly instead of a JavaScript string.
- Fix: `typedValueToObject` now converts `godbus.ObjectPath` to `string(v)` and `godbus.Signature` to `v.String()` for JavaScript-facing `value` fields while preserving typed validation in core.

### What I learned

- For JavaScript-facing helper results, Go-backed domain values should be normalized into plain JavaScript-readable fields unless the design explicitly chooses reflected Go objects.
- A separate adapter layer pays off immediately: the pure core can preserve D-Bus-aware values, while `dbusgoja` controls what JavaScript sees.
- The go-go-goja engine can include this external module in tests with `UseModuleMiddleware(engine.MiddlewareOnly("dbus"))` as long as the module package is blank-imported first.

### What was tricky to build

- The module is outside the `go-go-goja` repository, so tests must explicitly blank-import `github.com/go-go-golems/goja-dbus/pkg/modules/dbus`; otherwise the module's `init()` registration never runs.
- `go.mod` needs a local `replace github.com/go-go-golems/go-go-goja => ../go-go-goja` so `GOWORK=off` tests can use the checked-out go-go-goja implementation.
- The JS helper functions use `goja.FunctionCall` where range validation depends on numeric conversion and errors must be thrown with `panic(vm.NewGoError(err))`.

### What warrants a second pair of eyes

- The `DBusTypedValue` JavaScript shape currently uses a visible `__dbusTyped` marker. This is simple and testable, but later codec code should decide whether the marker should become non-enumerable or internal.
- The local `replace` is useful for development but may need a release strategy once go-go-goja publishes the required version.
- The Phase 2 `Policy` type is intentionally incomplete and must be expanded before system bus or service ownership features are exposed.

### What should be done in the future

- Add decoding from JavaScript `DBusTypedValue` objects back into `dbuscore.TypedValue` for method-call inputs.
- Add session bus connection builders and Promise-based method calls.
- Expand TypeScript declarations as builders are added.

### Code review instructions

- Start with `pkg/modules/dbus/module.go` to confirm the `NativeModule` and TypeScript declaration shape.
- Review `pkg/dbuscore/types.go` for D-Bus validation behavior.
- Review `pkg/dbusgoja/typed_values.go` for JavaScript-facing object shape and thrown errors.
- Validate with:
  - `GOWORK=off go test ./... -count=1`

### Technical details

Commands:

```bash
cd goja-dbus
gofmt -w pkg/dbuscore pkg/dbusgoja pkg/modules/dbus
GOWORK=off go test ./... -count=1
git commit -m "Add dbus native module skeleton"
```

## Step 4: Phase 3 Session Bus Connect and Method Call Builders

I added the first Promise-based D-Bus client path. JavaScript can now create bus builders with `dbus.session()`, `dbus.system()`, or `dbus.connect(address)`, apply timeouts and simple policy options, connect asynchronously, and build remote method calls through `destination().object().interface().method().in().out().call()`.

This phase also adds the core connection, policy, and scalar codec layer needed for real method calls. The always-on tests do not require a running D-Bus daemon; the real `org.freedesktop.DBus.GetId` test is present but guarded behind `GOJA_DBUS_INTEGRATION=1`.

### Prompt Context

**User prompt (verbatim):** (same as Step 2)

**Assistant interpretation:** Continue executing the documented phases sequentially, committing implementation milestones and recording validation/failures.

**Inferred user intent:** Incrementally turn the design guide into working code without losing the reasoning trail.

**Commit (code):** `7ffee5514db236c03e5bf034e0fe2e8697589fd7` — "Add D-Bus connect and method call builders"

### What I did

- Added `pkg/dbuscore/bus.go` with `BusKind`, `ConnectOptions`, `Bus`, connection lifecycle, `Close`, and `Call`.
- Added `pkg/dbuscore/codec.go` with scalar marshaling/unmarshaling for strings, uint32, int32, object paths, signatures, and variants.
- Expanded `pkg/dbuscore/policy.go` with connect checks, call checks, simple exact/suffix-wildcard pattern matching, and default session-bus-only policy.
- Added `pkg/dbusgoja/promise.go` to settle asynchronous work back on the runtime owner through `runtimebridge.RuntimeServices`.
- Added `pkg/dbusgoja/bus_exports.go` for top-level `session`, `system`, and `connect` exports.
- Added `pkg/dbusgoja/builders.go` for immutable bus, destination, object, interface, and method builders.
- Added `pkg/dbusgoja/decode.go` to decode JavaScript `DBusTypedValue` objects back into core typed values.
- Added `pkg/dbusgoja/policy.go` to decode lowerCamel JavaScript policy objects.
- Updated `pkg/modules/dbus/module.go` docs and TypeScript declarations to include builders.
- Added always-on tests for codec and policy behavior.
- Added a Promise rejection test proving `dbus.system().connect()` is denied by default policy without needing a system bus.
- Added an opt-in integration test for `org.freedesktop.DBus.GetId`, guarded by `GOJA_DBUS_INTEGRATION=1`.
- Ran `GOWORK=off go test ./... -count=1` successfully.

### Why

- This is the first phase where the module becomes more than helper functions: it now has a D-Bus client API shape that matches the design guide.
- Policy denial is tested before real bus access so the security boundary is exercised in normal CI.
- Real D-Bus integration is opt-in to avoid making CI depend on a desktop/session bus.

### What worked

- Always-on tests passed:
  - `ok github.com/go-go-golems/goja-dbus/pkg/dbuscore`
  - `ok github.com/go-go-golems/goja-dbus/pkg/modules/dbus`
- Promise rejection for default-denied system bus access works without contacting a real bus.
- The code now validates object paths before creating remote object builders.

### What didn't work

- The first Phase 3 test run failed at build time because `Policy` contained a slice and was compared directly to a zero struct:
  - Error: `invalid operation: policy == (Policy{}) (struct containing []string cannot be compared)`
  - Fix: added `Policy.IsZero()` and used that in `Connect`.
- The next test run failed because `module_object.go` imported `dbuscore` but no longer used it after the builder export refactor:
  - Error: `"github.com/go-go-golems/goja-dbus/pkg/dbuscore" imported and not used`
  - Fix: removed the unused import.
- Adding Promise tests required additional imports in `module_test.go` (`encoding/json`, `os`, `time`, and `goja`); the initial compile surfaced these missing imports.

### What I learned

- A default-denied system bus test is a useful no-daemon test case because it exercises the Promise path, policy path, and runtime-owner settlement path.
- The builder layer should convert JavaScript inputs to plain Go values during the owner-thread `.in(...)` call, before any goroutine starts.
- Keeping the real bus test opt-in makes Phase 3 safe to run in headless CI while still documenting how to validate against a real D-Bus daemon.

### What was tricky to build

- The Promise helper must create the Promise on the owner call, do blocking work in a goroutine, and then use `PostWithCustomContext` for settlement. This follows go-go-goja's async pattern and avoids touching JavaScript values from the goroutine.
- The method builder has to copy builder state on each fluent method to preserve immutability. The implementation copies input slices when appending `.in(...)` arguments.
- The adapter has to decode JavaScript typed helper objects immediately. Retaining arbitrary `goja.Value` values inside a builder that may later execute in a goroutine would violate the runtime-owner rule.

### What warrants a second pair of eyes

- The current `Policy.AllowCall` matcher is intentionally simple. It supports exact matches and suffix `*` prefix matches, but not a complete glob language.
- Runtime-owned bus cleanup is currently explicit via `bus.close()`; Phase 4 should revisit runtime shutdown closers when subscriptions are added.
- The codec only supports a scalar subset plus variants. Arrays, structs, and dictionaries remain future work.

### What should be done in the future

- Run `GOJA_DBUS_INTEGRATION=1 GOWORK=off go test ./pkg/modules/dbus -run TestDBusGetIdIntegration -count=1` on a machine with a working session bus.
- Implement signal subscriptions with close handles.
- Expand the codec for arrays, dictionaries, and structs before implementing properties and service export.

### Code review instructions

- Start with `pkg/dbusgoja/builders.go` for the JavaScript API flow.
- Review `pkg/dbusgoja/promise.go` for owner-thread settlement.
- Review `pkg/dbuscore/bus.go` for D-Bus connection/call behavior and context/timeout handling.
- Review `pkg/dbuscore/codec.go` for scalar conversion assumptions.
- Validate with:
  - `GOWORK=off go test ./... -count=1`
  - Optional: `GOJA_DBUS_INTEGRATION=1 GOWORK=off go test ./pkg/modules/dbus -run TestDBusGetIdIntegration -count=1`

### Technical details

Commands:

```bash
cd goja-dbus
gofmt -w pkg/dbuscore pkg/dbusgoja pkg/modules/dbus
GOWORK=off go test ./... -count=1
git commit -m "Add D-Bus connect and method call builders"
```

## Step 5: Phase 4 Signal Subscription Builders

I added the first D-Bus signal subscription path. JavaScript can now call `bus.signals()` after connecting, add match filters, and call `listen(emitter)` with a Go-native `events.EventEmitter`; the Go side installs the D-Bus match rule, receives signals on a channel, and schedules `emitter.emit("signal", payload)` back onto the runtime owner.

This phase completes explicit close handling for subscriptions, but it deliberately leaves one lifecycle follow-up open: automatic subscription cleanup on runtime shutdown without an explicit `bus.close()` or subscription `close()` still needs a runtime-aware registrar or another closer hook design.

### Prompt Context

**User prompt (verbatim):** (same as Step 2)

**Assistant interpretation:** Continue executing implementation phases with commits and a detailed diary, while documenting incomplete lifecycle edges honestly.

**Inferred user intent:** Add signal support incrementally without hiding remaining cleanup work.

**Commit (code):** `a2b3d5c8d07f547046cf409c69d8c0f9ebfdcd36` — "Add D-Bus signal subscription builders"

### What I did

- Added `pkg/dbuscore/signals.go` with:
  - `SignalMatchRequest`
  - `SignalPayload`
  - `SignalSink`
  - `Subscription`
  - `Bus.Listen`
  - `SignalMatchRequest.MatchOptions`
  - idempotent `Subscription.Close`
- `Bus.Listen` now pairs:
  - `Conn.AddMatchSignalContext`
  - `Conn.Signal(ch)`
  - `Conn.RemoveSignal(ch)`
  - `Conn.RemoveMatchSignalContext`
- Added `pkg/dbusgoja/signals.go` with JavaScript builder methods:
  - `sender(name)`
  - `path(path)`
  - `interface(name)`
  - `member(name)`
  - `listen(emitter)`
- Updated `pkg/dbusgoja/builders.go` so connected bus objects expose `signals()`.
- Updated TypeScript declarations in `pkg/modules/dbus/module.go` with `SignalBuilder`, `SignalPayload`, and `SignalSubscription`.
- Added tests for match path validation, empty match options, and close idempotency for an empty subscription.
- Ran `GOWORK=off go test ./... -count=1` successfully.

### Why

- D-Bus signals are long-lived Go-side events and need the same owner-thread safety discipline as the design guide's connected EventEmitter pattern.
- Signal subscriptions require explicit close handles because match rules and signal channels otherwise remain registered on the connection.

### What worked

- Always-on tests pass without a real D-Bus daemon.
- Signal match construction validates object paths before installing match rules.
- `listen(emitter)` validates that the provided value is a go-go-goja `events.EventEmitter` by using `events.FromValue`.
- Signal delivery builds the JavaScript payload on the runtime owner before calling `emitter.Emit`.

### What didn't work

- I did not implement automatic cleanup on runtime shutdown in this phase. The module is currently a plain `modules.NativeModule`, whose loader does not receive `RuntimeModuleRegistrationContext.AddCloser`.
- Because of that, the Phase 4 task was narrowed in `tasks.md`: explicit close support is complete, but runtime-shutdown cleanup remains open.

### What I learned

- It is possible to use the existing `events.EventEmitter` safely without the full `jsevents.Manager` when the module schedules all `EventEmitter.Emit` calls back onto the owner thread.
- A separate `SignalSink` callback in `dbuscore` keeps the core package free of Goja while still allowing the adapter to own JavaScript delivery.

### What was tricky to build

- `events.EventEmitter` itself is not goroutine-safe. The signal goroutine must not call it directly; it only asks `runtimebridge.RuntimeServices.PostWithCustomContext` to do emission on the owner.
- The subscription close operation must cancel the goroutine and unregister both the signal channel and D-Bus match rule. Missing either half would leak delivery state.
- Runtime shutdown cleanup is trickier because it needs a registration-time closer hook, not just a module loader hook.

### What warrants a second pair of eyes

- The EventEmitter validation and retained emitter pointer should be reviewed against go-go-goja's event module invariants.
- The current implementation reports listener emission errors by emitting `error` on the same emitter; this mirrors Node-style behavior but may need a host-level unhandled-error policy.
- Runtime shutdown cleanup should be designed before relying on signals in long-running hosts.

### What should be done in the future

- Add runtime closer support, either by adding a runtime-aware registrar for `dbus` or by extending the module integration pattern.
- Add an opt-in real signal integration test against a known session-bus signal.
- Add richer signal body decoding once compound D-Bus signatures are implemented.

### Code review instructions

- Start with `pkg/dbuscore/signals.go` for match and close lifecycle.
- Review `pkg/dbusgoja/signals.go` for owner-thread EventEmitter delivery.
- Review `pkg/modules/dbus/module.go` for the public TypeScript signal shape.
- Validate with:
  - `GOWORK=off go test ./... -count=1`

### Technical details

Commands:

```bash
cd goja-dbus
gofmt -w pkg/dbuscore pkg/dbusgoja pkg/modules/dbus
GOWORK=off go test ./... -count=1
git commit -m "Add D-Bus signal subscription builders"
```

## Step 6: Phase 5 Service Export Checkpoint

I paused before implementing JavaScript-backed D-Bus service export and wrote a focused checkpoint note. The checkpoint documents why service export should wait until runtime cleanup and compound signature support are stronger.

This was an intentional scope-control decision. Client calls and signal subscriptions are now implemented, but service export would add the riskiest callback direction: external D-Bus calls would need to synchronously wait for JavaScript handler execution on the Goja owner.

### Prompt Context

**User prompt (verbatim):** (same as Step 2)

**Assistant interpretation:** Continue phase execution, but use the phase plan's checkpoint escape hatch when implementation would be premature.

**Inferred user intent:** Keep progress deliberate and reviewable rather than pushing risky service export code too early.

**Commit (code/docs):** `c5acc524bf2cb3f2d699a574d0ed62000cb07038` — "Docs: add service export checkpoint"

### What I did

- Added `design-doc/02-service-export-checkpoint.md`.
- Compared service export options:
  - immediate godbus method-table implementation;
  - hard-coded echo-service demo;
  - defer service export until cleanup and codecs are stronger.
- Chose to defer service export for the current pass.
- Updated Phase 5 tasks to mark the checkpoint complete and leave echo-service implementation/testing deferred.

### Why

- Service export is the highest-risk runtime boundary because Go/godbus goroutines would need to invoke JavaScript handlers through the runtime owner and return D-Bus replies correctly.
- The module still lacks automatic runtime-shutdown cleanup for D-Bus resources.
- The codec currently supports only scalar signatures plus variants; realistic service export needs arrays, structs, dictionaries, and `a{sv}`.

### What worked

- `docmgr doctor --ticket GOJA-DBUS-DESIGN --stale-after 30` passed after adding the checkpoint note.
- The checkpoint creates an explicit decision record instead of burying the deferral in a diary paragraph.

### What didn't work

- N/A. This phase was documentation and scope control, not code implementation.

### What I learned

- The current implementation has enough client/signal functionality to justify stopping before service export.
- A small echo-service demo would be tempting, but it could create the wrong abstraction before lifecycle and codec foundations are complete.

### What was tricky to build

- The tricky part was deciding not to build. The implementation plan allowed a checkpoint if service export was not clear, and the current runtime-cleanup gap makes immediate service export a poor next step.
- The checkpoint had to be technical enough to guide the future implementation rather than merely saying "deferred".

### What warrants a second pair of eyes

- The deferral decision should be reviewed by whoever owns the desired product scope. If a demo is more important than architecture cleanliness, a hard-coded echo-service experiment could be accepted as throwaway work.
- The future service dispatcher pseudocode should be reviewed before coding.

### What should be done in the future

- Implement runtime-owned resource cleanup.
- Add compound codecs.
- Reopen service export after those foundations are in place.

### Code review instructions

- Review `ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/design-doc/02-service-export-checkpoint.md`.
- Check that the open tasks in Phase 5 accurately reflect the deferral.
- Validate with:
  - `docmgr doctor --ticket GOJA-DBUS-DESIGN --stale-after 30`

### Technical details

Commands:

```bash
cd goja-dbus
docmgr doc add --ticket GOJA-DBUS-DESIGN --doc-type design-doc --title "Service Export Checkpoint"
docmgr doctor --ticket GOJA-DBUS-DESIGN --stale-after 30
git commit -m "Docs: add service export checkpoint"
```

## Step 7: Phase 6 Current API Documentation and Error Hardening

I finished the current implementation pass by documenting the API that actually exists now and adding a small JavaScript error-code mapping. The README now includes usage examples for GetId, typed helpers, signal subscriptions, embedding, test commands, and the current deferred status of service export.

This phase did not pretend that all originally imagined APIs are complete. Notification/properties examples and service export remain deferred because compound D-Bus codecs and runtime-owned cleanup still need more work.

### Prompt Context

**User prompt (verbatim):** (same as Step 2)

**Assistant interpretation:** Finish the current phase sequence with documentation, validation, and hardening appropriate to the implemented API surface.

**Inferred user intent:** Leave the repo in a coherent, tested, and reviewable state after several implementation commits.

**Commit (code/docs):** `de506ab6490a16231cc19ac7e529708b1481f6df` — "Document current dbus module API"

### What I did

- Expanded `README.md` with:
  - current status;
  - GetId example;
  - typed helper example;
  - EventEmitter signal example;
  - Go embedding sketch;
  - development and integration test commands.
- Added `pkg/dbusgoja/errors.go` with `dbusError`, which sets JavaScript error `name` to `DBusError` and `code` to `ERR_DBUS`.
- Updated Promise rejection in `pkg/dbusgoja/promise.go` to reject with `dbusError`.
- Updated the system-bus policy denial test to assert `err.code === "ERR_DBUS"`.
- Ran:
  - `GOWORK=off go test ./... -count=1`
  - `GOWORK=off go test -race ./pkg/modules/dbus -run TestDBusSystemConnectDeniedByDefault -count=1`
- Committed the changes.

### Why

- The README should match the current module, not the future design doc.
- JavaScript callers need a stable error code for policy denials and future D-Bus failures.
- A targeted race test gives some confidence that the Promise rejection path schedules correctly through the runtime owner.

### What worked

- Full always-on tests passed.
- Targeted race test passed:
  - `ok github.com/go-go-golems/goja-dbus/pkg/modules/dbus 1.028s`
- The current README now clearly marks service export and compound signatures as deferred.

### What didn't work

- N/A for this phase. The hardening and documentation changes passed on the first validation run.

### What I learned

- The implemented API is now coherent enough to document as a useful current subset: typed helpers, client calls, policy-denied system bus, and signal subscriptions.
- Keeping future examples out of the README avoids creating false expectations before compound codecs and service export exist.

### What was tricky to build

- Error mapping should not obscure the original Go error message. The current implementation wraps the Go error in `vm.NewGoError(err)` and then adds stable JavaScript fields.
- The README has to show `events` in the embedding middleware because signal examples need both `dbus` and `events` modules.

### What warrants a second pair of eyes

- The error taxonomy is intentionally minimal: all current adapter Promise rejections use `ERR_DBUS`. Future work may need subcodes such as `ERR_DBUS_POLICY`, `ERR_DBUS_TIMEOUT`, and `ERR_DBUS_MARSHAL`.
- README examples that use top-level `await` assume an evaluator or host that supports wrapping/awaiting scripts; embedding examples may need host-specific execution details.

### What should be done in the future

- Add compound codecs and properties support.
- Add runtime-owned cleanup.
- Reopen service export after the checkpoint prerequisites are done.
- Optionally run the real D-Bus integration test on a desktop/session-bus machine.

### Code review instructions

- Review `README.md` for accuracy against implemented APIs.
- Review `pkg/dbusgoja/errors.go` and `pkg/dbusgoja/promise.go` for error object behavior.
- Review `pkg/modules/dbus/module_test.go` for the policy-denial error-code assertion.
- Validate with:
  - `GOWORK=off go test ./... -count=1`
  - `GOWORK=off go test -race ./pkg/modules/dbus -run TestDBusSystemConnectDeniedByDefault -count=1`

### Technical details

Commands:

```bash
cd goja-dbus
gofmt -w pkg/dbusgoja/errors.go pkg/dbusgoja/promise.go pkg/modules/dbus/module_test.go
GOWORK=off go test ./... -count=1
GOWORK=off go test -race ./pkg/modules/dbus -run TestDBusSystemConnectDeniedByDefault -count=1
git commit -m "Document current dbus module API"
```
