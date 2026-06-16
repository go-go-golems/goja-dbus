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
    - Path: .goreleaser.yaml
      Note: |-
        renamed release metadata away from demo command
        release packaging for goja-dbus CLI
    - Path: Makefile
      Note: |-
        updated xgoja target paths for commit ea0a7a5
        install now uses generated goja-dbus binary
        goreleaser check and snapshot helpers
    - Path: cmd/goja-dbus/main.go
      Note: checked-in GoReleaser entrypoint for goja-dbus CLI
    - Path: cmd/goja-dbus/xgoja.yaml
      Note: |-
        moved xgoja binary spec for commit ea0a7a5
        canonical goja-dbus CLI spec after removing demo command
    - Path: pkg/dbuscore/policy.go
      Note: 'PR #1 policy hardening'
    - Path: pkg/dbusgoja/policy.go
      Note: JavaScript policy decoding fixes
    - Path: pkg/modules/dbus/module_test.go
      Note: regression coverage for explicit address denial
    - Path: pkg/xgoja/provider/docs/help/01-getting-started.md
      Note: |-
        bundled Glazed help docs for commit 50ad0ab
        updated bundled build instructions for commit ea0a7a5
    - Path: pkg/xgoja/provider/provider.go
      Note: xgoja provider registration for commit 50ad0ab
    - Path: pkg/xgoja/provider/verbs/examples.js
      Note: bundled jsverb examples for commit 50ad0ab
    - Path: ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/design-doc/01-goja-d-bus-module-intern-design-and-implementation-guide.md
      Note: Primary deliverable produced during the investigation
    - Path: ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/sources/01-dbus.md
      Note: Imported source material referenced by the diary
    - Path: xgoja.yaml
      Note: generated binary build spec for commit 50ad0ab
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

## Step 8: Runtime-Owned D-Bus Resource Cleanup

I closed the Phase 4 lifecycle gap by adding runtime-owned cleanup for D-Bus resources created through the JavaScript module. The Goja adapter now creates a per-module resource registry tied to the runtime lifetime context; when the runtime closes, that registry closes all tracked buses, and bus close cascades to tracked signal subscriptions.

This makes explicit `bus.close()` and `subscription.close()` still useful, but no longer the only cleanup path. It also removes runtime-shutdown cleanup from the README's deferred list.

### Prompt Context

**User prompt (verbatim):** "continue"

**Assistant interpretation:** Continue the previously defined implementation plan from the remaining open tasks.

**Inferred user intent:** Keep progressing through deferred foundational work with commits and diary updates.

**Commit (code):** `3d3f7b841dede1177a477733cc8da65589fdda22` — "Close D-Bus resources on runtime shutdown"

### What I did

- Added `pkg/dbusgoja/resources.go` with a `resourceRegistry` tied to `RuntimeServices.Lifetime()`.
- The registry tracks connected `*dbuscore.Bus` values and closes them when the runtime lifetime context is canceled.
- Updated `NewModuleObject` to create the registry and pass it into bus builders.
- Updated bus builders so successful `connect()` registers a bus, and explicit `bus.close()` unregisters it before closing.
- Updated `dbuscore.Bus` to track subscriptions.
- Updated `dbuscore.Bus.Close` to close tracked subscriptions before closing the underlying D-Bus connection.
- Added `dbuscore.Bus.Closed()` for lifecycle tests.
- Added tests for bus close cascading to subscriptions and registry close-on-lifetime-cancel behavior.
- Removed runtime cleanup from the README deferred list.
- Ran `GOWORK=off go test ./... -count=1`.

### Why

- Phase 4 left an explicit open task: signal subscription goroutines should exit on runtime shutdown even if JavaScript forgets to call close.
- Service export remains deferred partly because long-lived resources need a safe cleanup model. This step strengthens that foundation.

### What worked

- `GOWORK=off go test ./... -count=1` passed.
- The registry tests confirmed runtime lifetime cancellation closes tracked buses.
- The bus test confirmed `Bus.Close` cancels tracked subscriptions.

### What didn't work

- The first idea for registry testing tried to infer closure by calling `Bus.Call` on a zero-value bus, but a zero-value bus already reports `dbus: bus is closed` because it has no connection. That could not distinguish registry cleanup from an unusable zero-value bus.
- Fix: added a small `Bus.Closed()` method and tested the explicit closed flag instead.

### What I learned

- A plain `modules.NativeModule` can still perform runtime cleanup if its loader obtains `RuntimeServices.Lifetime()` and starts a non-Goja goroutine that only closes plain Go resources.
- Tracking subscriptions inside `dbuscore.Bus` is simpler than asking the JavaScript adapter to remember every subscription independently.

### What was tricky to build

- `Bus.Close` needed to copy subscriptions while holding the bus mutex, release the mutex, then close subscriptions. Calling `Subscription.Close` while still holding the bus mutex would deadlock because subscription close unregisters itself from the bus.
- Explicit `bus.close()` must unregister the bus from the registry so the later runtime cleanup goroutine does not attempt to close it again. The close methods are idempotent, but unregistering keeps the registry accurate.

### What warrants a second pair of eyes

- The registry goroutine starts when the module is required. This is acceptable because it only waits on runtime lifetime and closes Go resources, but reviewers should confirm there is no leak if a runtime never requires `dbus`.
- The cleanup path uses `context.Background()` during lifetime cancellation. A future refinement might use a bounded timeout context for close operations.

### What should be done in the future

- Add compound signature codecs now that runtime cleanup is stronger.
- Revisit the service export checkpoint after compound codecs and perhaps a tighter close timeout policy.

### Code review instructions

- Review `pkg/dbusgoja/resources.go` for runtime-lifetime cleanup behavior.
- Review `pkg/dbuscore/bus.go` and `pkg/dbuscore/signals.go` for subscription tracking and close ordering.
- Validate with `GOWORK=off go test ./... -count=1`.

### Technical details

Commands:

```bash
cd goja-dbus
gofmt -w pkg/dbuscore pkg/dbusgoja
GOWORK=off go test ./... -count=1
git commit -m "Close D-Bus resources on runtime shutdown"
```

## Step 9: Compound D-Bus Typed Helpers and Codec Foundations

I added the first compound D-Bus typed helper and codec foundation. JavaScript now has `dbus.array(signature, values)`, `dbus.dict(signature, values)`, and `dbus.struct(signature, values)` helpers in addition to scalar helpers, and the core codec can marshal common arrays, `a{sv}` dictionaries, and flat structs.

This is still a foundation, not complete D-Bus signature coverage. It supports the common subset needed for many notification/property-style calls, while keeping unsupported signatures explicit errors rather than silently guessing.

### Prompt Context

**User prompt (verbatim):** (same as Step 8)

**Assistant interpretation:** Continue from the remaining foundational tasks after runtime cleanup.

**Inferred user intent:** Progress toward the prerequisites identified in the service export checkpoint, especially compound codec support.

**Commit (code):** `833cb85e40c7def04ecba848aef87a149d5fcc7a` — "Add compound D-Bus typed helpers"

### What I did

- Added core marshaling support for:
  - `as` string arrays;
  - `au` uint32 arrays;
  - `ai` int32 arrays;
  - `ao` object-path arrays;
  - `av` variant arrays;
  - `a{sv}` string-to-variant dictionaries;
  - flat structs such as `(su)`.
- Added recursive JavaScript decoding for arrays, plain objects, and nested `DBusTypedValue` objects.
- Added JavaScript helpers:
  - `dbus.array(signature, values)`;
  - `dbus.dict(signature, values)`;
  - `dbus.struct(signature, values)`.
- Updated TypeScript declarations for the new helpers.
- Updated README examples and the current/deferred status.
- Added core codec tests for arrays, dictionaries, and structs.
- Extended the runtime module test to assert helper object shapes for array/dict/struct.
- Ran `GOWORK=off go test ./... -count=1`.

### Why

- The service export checkpoint identified compound codecs as a prerequisite for properties, notifications, and eventually service export.
- Common D-Bus APIs use `as` and `a{sv}` frequently; supporting them makes the client API substantially more useful.

### What worked

- Full tests passed after fixing the JavaScript decoder recursion issue.
- The helper API preserves explicit signatures while giving JavaScript users ergonomic wrappers.
- Nested typed values now survive JavaScript object/array decoding, which is required for `a{sv}` values like `{ urgency: dbus.variant("u", dbus.u32(1)) }`.

### What didn't work

- The first compile failed because the TypeScript declaration raw string slice missed a trailing comma after the new `struct` declaration:
  - Error: `pkg/modules/dbus/module.go:94:79: missing ',' before newline in composite literal`
  - Fix: added the missing comma.
- The first decoder implementation panicked on missing object properties because `obj.Get(typedMarker)` can return nil:
  - Symptom: `panic: runtime error: invalid memory address or nil pointer dereference`
  - Fix: check for nil/undefined before calling `ToBoolean()`.
- The next decoder implementation caused a stack overflow because primitive strings were converted to objects and recursively decoded through enumerable string indices:
  - Symptom: `fatal error: stack overflow` with repeated `decodeJSValue` frames.
  - Fix: only treat `ClassName() == "Array"` as an array and `ClassName() == "Object"` as a plain object.

### What I learned

- Goja `ToObject` is too broad for generic decoding because primitives also become wrapper objects. Decoder code must check object class before recursing.
- Preserving nested typed values is necessary before dictionary helpers are useful; otherwise `dbus.variant(...)` inside a JavaScript object degrades into a plain exported map.
- A deliberately limited signature subset is safer than pretending to support all D-Bus signatures dynamically.

### What was tricky to build

- D-Bus marshaling depends on Go concrete types. For arrays, returning `[]string`, `[]uint32`, `[]int32`, `[]dbus.ObjectPath`, or `[]dbus.Variant` is safer than returning `[]any` and hoping godbus infers the intended element type.
- Flat struct support is represented as `[]any` for now. That is enough for internal codec shape tests, but it may require revisiting with real godbus integration before relying on struct inputs heavily.
- The recursive JS decoder must handle typed helper objects, arrays, and plain objects without walking into primitive wrapper objects or function objects.

### What warrants a second pair of eyes

- `marshalStruct` should be reviewed before using it in real D-Bus calls; godbus may require concrete struct representations for some cases.
- The supported signature subset should be documented clearly for users before adding notification examples that depend on `as` and `a{sv}`.
- `splitFlatSignatures` intentionally supports only flat signatures and should not be mistaken for a full D-Bus signature parser.

### What should be done in the future

- Add real integration coverage for a method that accepts `as` and `a{sv}`.
- Add properties client helpers now that `a{sv}` is available.
- Revisit service export after properties and compound codec integration tests.

### Code review instructions

- Review `pkg/dbuscore/codec.go` for supported signature behavior and explicit unsupported cases.
- Review `pkg/dbusgoja/decode.go` for recursive decoding safety.
- Review `pkg/dbusgoja/typed_values.go` for helper export behavior.
- Validate with `GOWORK=off go test ./... -count=1`.

### Technical details

Commands:

```bash
cd goja-dbus
gofmt -w pkg/dbuscore pkg/dbusgoja pkg/modules/dbus
GOWORK=off go test ./... -count=1
git commit -m "Add compound D-Bus typed helpers"
```

## Step 10: xgoja Binary with Bundled Help and Example Verbs

I added a goja-dbus xgoja provider so the project can build a generated binary that carries the native `dbus` module, user-facing Glazed help pages, and bundled JavaScript example verbs. This turns the module from a Go embedding-only package into a runnable documentation and smoke-test host.

The generated host now has `getting-started`, `user-guide`, and `api-reference` help pages, plus example verbs for typed helper shapes, default system-bus policy denial, and a printable `GetId` script for machines with a session bus.

### Prompt Context

**User prompt (verbatim):** "Create an `xgoja`-built binary for `goja-dbus` that bundles documentation (`getting-started`, `user-guide`, `api-reference`) in Glazed help entry format and exposes various verbs as examples. Use `xgoja help --all` to find relevant help pages."

**Assistant interpretation:** Add an xgoja build spec and provider package for goja-dbus, ship help docs and jsverb examples through provider sources, build the generated binary, and validate the bundled commands.

**Inferred user intent:** Make goja-dbus easy to explore as a standalone xgoja-generated CLI, not only as a native module embedded by another Go program.

**Commit (code):** `50ad0abd0bdb49c86bc77b523416f55ac6fa543d` — "Add xgoja provider with bundled DBus docs"

### What I did

- Added `pkg/xgoja/provider/provider.go`, which registers provider package `dbus` with:
  - the existing native `dbus` module;
  - a provider help source named `docs`;
  - a provider jsverb source named `examples`.
- Added Glazed help entries:
  - `pkg/xgoja/provider/docs/help/01-getting-started.md`;
  - `pkg/xgoja/provider/docs/help/02-user-guide.md`;
  - `pkg/xgoja/provider/docs/help/03-api-reference.md`.
- Added bundled example verbs in `pkg/xgoja/provider/verbs/examples.js`:
  - `typed-values`;
  - `denied-system-bus`;
  - `get-id-script`.
- Added `xgoja.yaml` selecting:
  - the goja-dbus provider module `dbus`;
  - core `events` and `timer` modules;
  - provider help docs;
  - provider jsverbs;
  - `eval`, `run`, and `verbs` builtin commands.
- Added Makefile targets:
  - `make xgoja-doctor`;
  - `make xgoja-build`.
- Updated `README.md` with build and usage instructions for `dist/goja-dbus-xgoja`.
- Ran validation and smoke tests.

### Why

- xgoja provider sources are the idiomatic way to ship a native module, provider-owned help docs, and provider-owned jsverbs into a generated binary.
- The help docs provide a user-facing layer over the lower-level design ticket and README.
- Example verbs make the binary useful even on hosts without a working session bus: `typed-values` and `denied-system-bus` are safe always-on examples.

### What worked

- `GOWORK=off go test ./... -count=1` passed.
- `make xgoja-doctor` passed schema, module-resolution, and source-plan checks.
- `make xgoja-build` produced `dist/goja-dbus-xgoja`.
- The generated binary rendered all three help pages:
  - `./dist/goja-dbus-xgoja help getting-started`;
  - `./dist/goja-dbus-xgoja help user-guide`;
  - `./dist/goja-dbus-xgoja help api-reference`.
- The generated binary ran bundled example verbs:
  - `./dist/goja-dbus-xgoja verbs examples typed-values`;
  - `./dist/goja-dbus-xgoja verbs examples denied-system-bus`;
  - `./dist/goja-dbus-xgoja verbs examples get-id-script`.

### What didn't work

- Running `GOWORK=off go run ../go-go-goja/cmd/xgoja doctor -f xgoja.yaml` from inside `goja-dbus` initially failed because the `goja-dbus` module did not have all xgoja command transitive `go.sum` entries:
  - `missing go.sum entry for module providing package github.com/charmbracelet/glamour`
  - `missing go.sum entry for module providing package github.com/evanw/esbuild/pkg/api`
  - plus related Glazed, bobatea, tree-sitter, and protobuf/grpc entries.
- Fix: make the repeatable targets invoke `go run ./cmd/xgoja` from the sibling `../go-go-goja` module, while pointing `-f` and `--output` back at `goja-dbus`.
- The first generated binary could not load provider help docs:
  - `load help source docs: unknown provider help source dbus.docs`
- Fix: align the goja-dbus provider package ID with the xgoja provider id by registering `PackageID = "dbus"`.
- The next generated binary could not resolve core modules:
  - `runtime references unknown provider module core.events`
- Fix: use the core provider package id `go-go-goja-core` in `xgoja.yaml` instead of a local alias `core`.
- The first jsverb metadata used `output: "object"`, which the scanner rejected:
  - `examples.js#deniedSystemBus has unsupported output mode "object"`
- Fix: omit the output mode for structured object/table output and keep only `output: "text"` for the script-printing verb.
- The generated binary's help command does not support `help --all`, even though the source xgoja command used during discovery did. I updated the bundled getting-started page to use `help`, `help getting-started`, `help user-guide`, and `help api-reference` for the generated binary.

### What I learned

- In xgoja provider references, the provider id used in `runtime.modules` and provider source selectors must match the package id registered by the provider, not merely a convenient YAML alias.
- Provider-shipped help and jsverb sources are a good fit for this module because they keep documentation/examples versioned with the native provider.
- jsverb `output` is intentionally limited; plain structured object returns should not specify `output: "object"`.

### What was tricky to build

- The most subtle issue was provider identity. The generated host reported unknown provider sources/modules at runtime, not at build time, because the spec could compile while still referencing provider ids that did not match registered package ids. I fixed this by aligning `PackageID = "dbus"` and using `go-go-goja-core` for the first-party core provider.
- Build ergonomics were also tricky because invoking another module's `go run` from the goja-dbus module caused missing `go.sum` failures. The Makefile now runs xgoja from the sibling module and writes the binary to `../goja-dbus/dist/goja-dbus-xgoja`.
- The help documentation had to distinguish between the source `xgoja` tool, which was used for discovery and supports `help --all`, and the generated goja-dbus binary, whose help command lists topics without `--all`.

### What warrants a second pair of eyes

- Review `xgoja.yaml` provider ids and explicit local replacements before using this spec outside the current sibling-checkout workspace.
- Review whether the generated binary should eventually be built by an installed/released `xgoja` binary rather than by `cd ../go-go-goja && go run ./cmd/xgoja`.
- Review the public help page claims against future API changes, especially once properties helpers and real compound integration tests land.

### What should be done in the future

- Add a CI smoke target that builds `dist/goja-dbus-xgoja` and runs help/verb smoke tests when a sibling or released xgoja is available.
- Add properties client helpers and then update the xgoja help docs and examples with real `org.freedesktop.DBus.Properties` usage.
- Consider adding an opt-in xgoja integration verb/test that runs the real `GetId` call when a session bus is present.

### Code review instructions

- Start with `pkg/xgoja/provider/provider.go` to verify provider registration, embedded filesystems, and module loader wiring.
- Review `xgoja.yaml` for provider ids, selected runtime modules, sources, commands, and output artifact.
- Review `pkg/xgoja/provider/docs/help/*.md` for Glazed help frontmatter and user-facing accuracy.
- Review `pkg/xgoja/provider/verbs/examples.js` for safe examples and jsverb metadata.
- Validate with:
  - `GOWORK=off go test ./... -count=1`
  - `make xgoja-doctor`
  - `make xgoja-build`
  - `./dist/goja-dbus-xgoja help getting-started`
  - `./dist/goja-dbus-xgoja help user-guide`
  - `./dist/goja-dbus-xgoja help api-reference`
  - `./dist/goja-dbus-xgoja verbs examples typed-values`
  - `./dist/goja-dbus-xgoja verbs examples denied-system-bus`
  - `./dist/goja-dbus-xgoja verbs examples get-id-script`

### Technical details

Commands:

```bash
cd goja-dbus
GOWORK=off go test ./... -count=1
make xgoja-doctor
make xgoja-build
./dist/goja-dbus-xgoja help getting-started
./dist/goja-dbus-xgoja help user-guide
./dist/goja-dbus-xgoja help api-reference
./dist/goja-dbus-xgoja verbs examples typed-values
./dist/goja-dbus-xgoja verbs examples denied-system-bus
./dist/goja-dbus-xgoja verbs examples get-id-script
git commit -m "Add xgoja provider with bundled DBus docs"
```

## Step 11: Move xgoja Build Spec Under `cmd/goja-dbus`

I moved the generated binary's xgoja spec from the repository root into `cmd/goja-dbus/xgoja.yaml`. This puts the binary configuration next to the command namespace that owns it while keeping the generated executable output in `dist/goja-dbus-xgoja`.

The move required adjusting relative provider replacement paths and Makefile targets so the sibling-checkout build flow still works from the repository root.

### Prompt Context

**User prompt (verbatim):** "mv the xgoja.yaml binary to be in ./goja-dbus/cmd/goja-dbus/"

**Assistant interpretation:** Move the xgoja build specification for the goja-dbus binary under `cmd/goja-dbus/` and update all references/validation commands accordingly.

**Inferred user intent:** Keep binary-specific xgoja configuration under the command directory rather than at repository root.

**Commit (code/docs):** `ea0a7a5fa3df0040a03929670b6c9c1a612af491` — "Move xgoja spec under cmd/goja-dbus"

### What I did

- Moved `xgoja.yaml` to `cmd/goja-dbus/xgoja.yaml`.
- Updated the spec's relative local replacements:
  - `github.com/go-go-golems/go-go-goja` now resolves through `../../../go-go-goja`;
  - `github.com/go-go-golems/goja-dbus` now resolves through `../..`.
- Updated Makefile targets to use `../goja-dbus/cmd/goja-dbus/xgoja.yaml`.
- Updated README and bundled getting-started help text to reference `cmd/goja-dbus/xgoja.yaml`.
- Revalidated the xgoja doctor/build flow and smoke-tested generated help plus a bundled verb.

### Why

- The xgoja spec describes a generated binary, so placing it under `cmd/goja-dbus/` better matches normal Go project layout.
- Keeping the Makefile targets stable preserves the existing root-level developer workflow.

### What worked

- `make xgoja-doctor` passed with the moved spec.
- `make xgoja-build` produced `dist/goja-dbus-xgoja`.
- `./dist/goja-dbus-xgoja help getting-started` rendered successfully.
- `./dist/goja-dbus-xgoja verbs examples denied-system-bus` ran successfully.
- `GOWORK=off go test ./... -count=1` passed.

### What didn't work

- My first automated replacement accidentally rewrote the core provider replacement path as `../.../../../go-go-goja`.
- Fix: corrected it to `../../../go-go-goja`, which `xgoja doctor` then resolved to the sibling go-go-goja checkout.

### What I learned

- The xgoja spec's `module.replace` paths are resolved relative to the spec file path, so moving the spec requires recalculating local replacements.
- The root Makefile can hide that relocation from day-to-day users by passing the new `-f cmd/goja-dbus/xgoja.yaml` path.

### What was tricky to build

- The path relocation is small but easy to get wrong because the command invoking xgoja runs from the sibling `go-go-goja` module while the spec file lives inside `goja-dbus/cmd/goja-dbus`. The validation step was necessary to prove both local replacements still resolve correctly.

### What warrants a second pair of eyes

- Review whether the generated binary name should eventually be shortened from `goja-dbus-xgoja` to `goja-dbus` now that the spec lives under `cmd/goja-dbus`.
- Review whether a future checked-in generated host should live in this command directory or whether xgoja generation should remain purely build-time.

### What should be done in the future

- If the command becomes the primary CLI, consider renaming the artifact from `dist/goja-dbus-xgoja` to `dist/goja-dbus` and updating release packaging.

### Code review instructions

- Review `cmd/goja-dbus/xgoja.yaml` for relative `module.replace` paths.
- Review `Makefile` for root-level `xgoja-doctor` and `xgoja-build` target paths.
- Review README and `pkg/xgoja/provider/docs/help/01-getting-started.md` for stale root `xgoja.yaml` references.
- Validate with:
  - `make xgoja-doctor`
  - `make xgoja-build`
  - `GOWORK=off go test ./... -count=1`

### Technical details

Commands:

```bash
cd goja-dbus
mkdir -p cmd/goja-dbus
mv xgoja.yaml cmd/goja-dbus/xgoja.yaml
make xgoja-doctor
make xgoja-build
./dist/goja-dbus-xgoja help getting-started
./dist/goja-dbus-xgoja verbs examples denied-system-bus
GOWORK=off go test ./... -count=1
git commit -m "Move xgoja spec under cmd/goja-dbus"
```

## Step 12: Rename Generated xgoja Artifact to `goja-dbus`

I renamed the generated xgoja host from `goja-dbus-xgoja` to `goja-dbus`. The xgoja spec now describes the command name, generated module, binary artifact, and TypeScript declaration artifact with the shorter project command name.

This follows the previous relocation of the spec into `cmd/goja-dbus/`: the command directory and generated binary now use the same name.

### Prompt Context

**User prompt (verbatim):** "goja-debus"

**Assistant interpretation:** The user started to specify the desired binary/name but corrected it in the next prompt.

**Inferred user intent:** Rename the generated xgoja artifact to the canonical project command name.

**User prompt (verbatim):** "goja-dbus"

**Assistant interpretation:** Use `goja-dbus` as the generated binary/artifact name instead of `goja-dbus-xgoja`.

**Inferred user intent:** Align the generated command name with the repository/module name.

**Commit (code/docs):** `e9f91364bcf219111cd9330e4acf4d68e118e961` — "Rename xgoja artifact to goja-dbus"

### What I did

- Updated `cmd/goja-dbus/xgoja.yaml`:
  - `name: goja-dbus`;
  - `app.name: goja-dbus`;
  - generated module `xgoja.generated/goja-dbus`;
  - binary output `dist/goja-dbus`;
  - DTS output `dist/goja-dbus.d.ts`.
- Updated `make xgoja-build` to write `dist/goja-dbus`.
- Updated README and bundled getting-started help commands from `./dist/goja-dbus-xgoja` to `./dist/goja-dbus`.
- Rebuilt and smoke-tested the renamed binary.

### Why

- The generated host now lives under `cmd/goja-dbus`, so the shorter artifact name is clearer and better matches normal CLI naming.
- Keeping `xgoja` out of the user-facing binary name makes the tool feel like the goja-dbus CLI rather than an implementation detail.

### What worked

- `make xgoja-doctor` passed.
- `make xgoja-build` produced `dist/goja-dbus`.
- `./dist/goja-dbus help getting-started` rendered successfully.
- `./dist/goja-dbus verbs examples typed-values` ran successfully.
- `GOWORK=off go test ./... -count=1` passed.

### What didn't work

- N/A. The rename was mechanical and validated cleanly.

### What I learned

- The xgoja artifact name is controlled in several places: top-level spec name, app name, generated Go module path, artifact output paths, Makefile overrides, and user-facing docs.

### What was tricky to build

- The main risk was stale command examples in bundled help/README. I used repository search for `goja-dbus-xgoja` and validated the rendered help after replacement.

### What warrants a second pair of eyes

- Review whether the older placeholder `cmd/goja-dbus-demo` should stay, be removed, or be renamed now that `cmd/goja-dbus/xgoja.yaml` is the main command configuration.

### What should be done in the future

- Decide how release packaging should handle both the generated xgoja binary and the existing demo command.

### Code review instructions

- Review `cmd/goja-dbus/xgoja.yaml` for name/module/output changes.
- Review README and bundled getting-started help for updated commands.
- Validate with:
  - `make xgoja-doctor`
  - `make xgoja-build`
  - `./dist/goja-dbus help getting-started`
  - `./dist/goja-dbus verbs examples typed-values`
  - `GOWORK=off go test ./... -count=1`

### Technical details

Commands:

```bash
cd goja-dbus
make xgoja-doctor
make xgoja-build
./dist/goja-dbus help getting-started
./dist/goja-dbus verbs examples typed-values
GOWORK=off go test ./... -count=1
git commit -m "Rename xgoja artifact to goja-dbus"
```

## Step 13: Remove the Placeholder Demo Command

I removed the old placeholder `cmd/goja-dbus-demo` command so the project now has a single canonical goja-dbus CLI identity. The xgoja spec under `cmd/goja-dbus/xgoja.yaml` is now the command configuration, and `make install` installs the generated `dist/goja-dbus` binary.

This also removes the previous split identity where the repository was `goja-dbus` but the template-era executable was `goja-dbus-demo`.

### Prompt Context

**User prompt (verbatim):** "remove it, all goja-dbus"

**Assistant interpretation:** Remove the remaining placeholder demo command and update build/install/release metadata so the command identity is consistently `goja-dbus`.

**Inferred user intent:** Avoid having both `goja-dbus-demo` and `goja-dbus`; make `goja-dbus` the only user-facing CLI name.

**Commit (code/docs):** `b492511693b2c151b7e1683bee1f3d36c7340da6` — "Remove placeholder goja-dbus demo command"

### What I did

- Deleted `cmd/goja-dbus-demo/main.go`.
- Updated `Makefile install` to depend on `xgoja-build` and install `dist/goja-dbus`.
- Updated `.goreleaser.yaml` names/descriptions away from `goja-dbus-demo` toward `goja-dbus`.
- Removed the README section that ran the placeholder demo command.
- Updated the design guide package tree to show `cmd/goja-dbus/xgoja.yaml` as the canonical CLI spec.
- Updated ticket tasks to reflect removal of the placeholder command.

### Why

- The generated xgoja CLI is now the real user-facing binary.
- Keeping the empty demo command around would confuse users and release tooling.

### What worked

- `make xgoja-doctor` passed.
- `make xgoja-build` produced `dist/goja-dbus`.
- `./dist/goja-dbus help getting-started` rendered successfully.
- `./dist/goja-dbus verbs examples denied-system-bus` ran successfully.
- `GOWORK=off go test ./... -count=1` passed and no longer lists `cmd/goja-dbus-demo`.

### What didn't work

- N/A. The deletion and references update validated cleanly.

### What I learned

- Once the xgoja-generated binary became `goja-dbus`, the empty demo command no longer served a useful purpose.
- The install target should install the generated binary rather than compiling a placeholder Go command.

### What was tricky to build

- The main subtlety is release packaging: `.goreleaser.yaml` no longer names `goja-dbus-demo`, but xgoja-generated release packaging still needs a dedicated follow-up because the command is generated from `cmd/goja-dbus/xgoja.yaml` rather than a checked-in `main.go` package.

### What warrants a second pair of eyes

- Review `.goreleaser.yaml` before the next release. It now uses `goja-dbus` naming, but the generated xgoja binary release flow may need a proper GoReleaser `dir`/generated-module setup.

### What should be done in the future

- Design the release pipeline for the generated xgoja host explicitly instead of relying on the old template GoReleaser build configuration.

### Code review instructions

- Review the removal of `cmd/goja-dbus-demo/main.go`.
- Review `Makefile install` to confirm it installs `dist/goja-dbus`.
- Review `.goreleaser.yaml` for remaining release-readiness work.
- Validate with:
  - `make xgoja-doctor`
  - `make xgoja-build`
  - `./dist/goja-dbus help getting-started`
  - `./dist/goja-dbus verbs examples denied-system-bus`
  - `GOWORK=off go test ./... -count=1`

### Technical details

Commands:

```bash
cd goja-dbus
git rm -r cmd/goja-dbus-demo
make xgoja-doctor
make xgoja-build
./dist/goja-dbus help getting-started
./dist/goja-dbus verbs examples denied-system-bus
GOWORK=off go test ./... -count=1
git commit -m "Remove placeholder goja-dbus demo command"
```

## Step 14: Wire Release Packaging Around the `goja-dbus` CLI

I wired release packaging to the canonical `goja-dbus` CLI by adding a checked-in `cmd/goja-dbus/main.go` host that mirrors the xgoja runtime plan. GoReleaser can now build a normal Go main package, while the xgoja spec remains available for generated-host validation and local xgoja builds.

This resolves the release gap left after removing the placeholder demo command: release tooling no longer points at an empty or missing `cmd/goja-dbus-demo`; it points at the real `goja-dbus` command.

### Prompt Context

**User prompt (verbatim):** "go ahead"

**Assistant interpretation:** Continue with the previously identified next step: wire release packaging around the generated/canonical `goja-dbus` binary.

**Inferred user intent:** Make release packaging work for `goja-dbus` rather than just documenting that it still needs review.

**Commit (code):** `8b2b6e68296b9b1eb4595f065f6283b56cb8023c` — "Wire GoReleaser to goja-dbus CLI"

**Commit (docs):** `97639afc17ec08fb013b63d0e99a0410b6fb5e14` — "Document GoReleaser snapshot workflow"

### What I did

- Added `cmd/goja-dbus/main.go`, a checked-in xgoja host for the `goja-dbus` runtime plan.
- Kept `cmd/goja-dbus/xgoja.yaml` as the source xgoja spec for doctor/build workflows.
- Updated `.goreleaser.yaml` so builds target `./cmd/goja-dbus` and produce `goja-dbus`.
- Replaced deprecated GoReleaser `snapshot.name_template` with `snapshot.version_template`.
- Replaced deprecated `brews` configuration with `homebrew_casks` and a `goja-dbus` binary entry.
- Ran `go mod tidy`, which pulled in the transitive CLI/runtime dependencies required by the checked-in host.
- Committed generated logcopter package stubs produced by `go generate ./...` so the GoReleaser before hook is reproducible.
- Added Makefile helpers:
  - `make goreleaser-check`;
  - `make goreleaser-snapshot`.
- Documented the snapshot workflow in README.

### Why

- GoReleaser builds Go main packages; it does not directly consume `xgoja.yaml` as a binary artifact pipeline.
- A checked-in host keeps release packaging simple: `go build ./cmd/goja-dbus` is now the release entrypoint.
- The xgoja spec still validates the intended runtime shape and can regenerate a temporary host for comparison during development.

### What worked

- `GOWORK=off go test ./... -count=1` passed.
- `GOWORK=off go build -o /tmp/goja-dbus-checkedin ./cmd/goja-dbus` produced a working binary.
- `/tmp/goja-dbus-checkedin help api-reference` rendered bundled provider docs.
- `GOWORK=off goreleaser check` passed after removing deprecated config fields.
- `GOWORK=off goreleaser release --skip=sign --snapshot --clean --single-target` succeeded.
- The snapshot produced:
  - `dist/linux_amd64/goja-dbus-linux_linux_amd64_v1/goja-dbus`;
  - `dist/linux_amd64/goja-dbus_0.0.1-next_linux_amd64.tar.gz`;
  - `dist/linux_amd64/goja-dbus_0.0.1-next_linux_amd64.deb`;
  - `dist/linux_amd64/goja-dbus_0.0.1-next_linux_amd64.rpm`.
- The GoReleaser-built binary ran a bundled example verb successfully:
  - `./dist/linux_amd64/goja-dbus-linux_linux_amd64_v1/goja-dbus verbs examples denied-system-bus`.

### What didn't work

- `GOWORK=off goreleaser check` initially failed because the config was valid but used deprecated fields:
  - `snapshot.name_template should not be used anymore`;
  - `brews is being phased out in favor of homebrew_casks`.
- Fix: changed `snapshot.name_template` to `snapshot.version_template` and converted `brews` to `homebrew_casks` with `binaries: [goja-dbus]`.
- The first `GOWORK=off go test ./... -count=1` after adding `cmd/goja-dbus/main.go` failed with:
  - `go: updates to go.mod needed; to update it: go mod tidy`.
- Fix: ran `GOWORK=off go mod tidy` and committed the resulting dependency changes.

### What I learned

- xgoja's temporary build host is easy to use locally, but release automation is simpler when a stable `main.go` exists in the repository.
- GoReleaser v2.13 treats deprecations as a failed `check`, so release config needs to stay ahead of deprecated keys even when the old fields still parse.
- The checked-in host and xgoja spec now duplicate the runtime plan; future changes should update both or introduce a generation/check workflow.

### What was tricky to build

- The main design tradeoff was duplication versus release simplicity. A checked-in `main.go` duplicates the xgoja runtime JSON, but it lets GoReleaser build with ordinary Go semantics and avoids requiring a generated temporary module during release.
- The GoReleaser before hook runs `go generate ./...`, which produced logcopter stubs in several packages. Committing them keeps repeated release runs stable.

### What warrants a second pair of eyes

- Review the duplicated runtime plan in `cmd/goja-dbus/main.go` against `cmd/goja-dbus/xgoja.yaml`; drift is the main ongoing risk.
- Review the Homebrew cask conversion before publishing to the tap because cask/formula conventions may differ from prior GoReleaser `brews` behavior.
- Review whether the release pipeline should eventually generate `cmd/goja-dbus/main.go` from `xgoja.yaml` instead of keeping it manually synchronized.

### What should be done in the future

- Add a check that compares the checked-in `cmd/goja-dbus/main.go` embedded runtime plan with the xgoja spec output.
- Decide whether to publish as a Homebrew cask long-term or restore a formula-style workflow if GoReleaser guidance changes.

### Code review instructions

- Review `cmd/goja-dbus/main.go` first; it is the GoReleaser entrypoint.
- Review `.goreleaser.yaml` for build targets, snapshot config, Homebrew cask config, and package metadata.
- Review `Makefile` for `goreleaser-check` and `goreleaser-snapshot` helpers.
- Validate with:
  - `GOWORK=off go test ./... -count=1`
  - `GOWORK=off go build -o /tmp/goja-dbus-checkedin ./cmd/goja-dbus`
  - `GOWORK=off goreleaser check`
  - `GOWORK=off goreleaser release --skip=sign --snapshot --clean --single-target`

### Technical details

Commands:

```bash
cd goja-dbus
GOWORK=off go mod tidy
GOWORK=off go test ./... -count=1
GOWORK=off go build -o /tmp/goja-dbus-checkedin ./cmd/goja-dbus
/tmp/goja-dbus-checkedin help api-reference
GOWORK=off goreleaser check
GOWORK=off goreleaser release --skip=sign --snapshot --clean --single-target
./dist/linux_amd64/goja-dbus-linux_linux_amd64_v1/goja-dbus verbs examples denied-system-bus
git commit -m "Wire GoReleaser to goja-dbus CLI"
git commit -m "Document GoReleaser snapshot workflow"
```

## Step 15: Address PR #1 Policy and Release Review Findings

I addressed the automated review findings on PR #1. The main code changes harden the policy boundary: explicit D-Bus addresses are no longer allowed by the default session-bus policy, and an explicitly empty JavaScript `allowCall: []` policy now means deny-all rather than falling back to unrestricted calls.

The committed local `replace github.com/go-go-golems/go-go-goja => ../go-go-goja` issue had already been resolved before this step; I verified the current `go.mod` has no local replace and that the `GOWORK=off` validation path still works.

### Prompt Context

**User prompt (verbatim):** "Address code review issues: https://github.com/go-go-golems/goja-dbus/pull/1"

**Assistant interpretation:** Inspect PR #1 review comments, implement the requested fixes, validate locally, and leave the branch ready for another review pass.

**Inferred user intent:** Resolve the blocking automated review comments so the PR can proceed.

**Commit (code):** `2078db801d11f9fa6e32e535eed0ac59117efedd` — "Harden DBus policy defaults"

### What I did

- Fetched PR #1 review comments with `gh pr view` and `gh api repos/go-go-golems/goja-dbus/pulls/1/comments`.
- Added `Policy.AllowAddressBus` so explicit-address connections require their own opt-in instead of piggybacking on session/system permissions.
- Updated `DefaultPolicy()` so it still allows session-bus connections and wildcard method calls, but does not allow explicit-address connections.
- Added `Policy.AllowCallSet` so an explicitly provided empty `allowCall: []` list is distinguishable from an omitted allow-call policy.
- Updated JavaScript policy decoding to set `AllowCallSet` when `allowCall` is present and to decode `allowAddressBus`.
- Hardened `decodePolicy` against nil values returned by `Object.Get`.
- Updated TypeScript declarations and API docs for `allowAddressBus` and explicit empty `allowCall` semantics.
- Added tests for:
  - default denial of explicit-address connections;
  - explicit opt-in for address connections;
  - explicit empty allow-call deny-all behavior in core policy;
  - JavaScript policy decoding for `allowCall: []` and `allowAddressBus`;
  - JavaScript `dbus.connect("unix:path=/run/dbus/system_bus_socket").connect()` denial by default.

### Why

- Without a separate address policy, callers could use `dbus.connect("unix:path=/run/dbus/system_bus_socket")` to bypass the default system-bus denial.
- Without an explicit allow-call presence bit, JavaScript callers who tried to narrow a policy with `allowCall: []` accidentally got unrestricted method calls.
- Removing local module replacements is necessary for clean CI, downstream checkouts, GoReleaser, and `go install` style builds.

### What worked

- `GOWORK=off go test ./... -count=1` passed.
- The pre-commit hook also ran `GOWORK=off go test ./...` and `GOWORK=off golangci-lint run -v`; both passed.
- `make xgoja-doctor` passed.
- `make xgoja-build` produced `dist/goja-dbus`.
- `./dist/goja-dbus verbs examples denied-system-bus` ran successfully.
- `make goreleaser-check` passed.
- `GOWORK=off goreleaser build --snapshot --clean --single-target` passed.
- `go.mod` currently has no local `replace` directive for `go-go-goja`.

### What didn't work

- The first new `dbusgoja` policy decoder test panicked:
  - `panic: runtime error: invalid memory address or nil pointer dereference`
  - Location: `pkg/dbusgoja/policy.go:17`
- Root cause: `obj.Get("allowSystemBus")` can return nil in this Goja version/path, and the decoder passed that nil value into `goja.IsUndefined` / `goja.IsNull` checks.
- Fix: added `isPolicyValueSet(value goja.Value) bool`, which checks `value != nil` before testing undefined/null.

### What I learned

- Policy structs need to preserve the difference between omitted fields and explicitly empty fields when JavaScript can use them to narrow permissions.
- Treating arbitrary D-Bus addresses as equivalent to session/system bus permissions is too permissive; address connections are a distinct capability.
- The Goja adapter should consistently nil-check object properties before calling methods on `goja.Value`.

### What was tricky to build

- The tricky part was maintaining the current default wildcard call policy while making explicit `allowCall: []` deny all calls. A plain `[]string` cannot represent that distinction, so `AllowCallSet` records whether the field was intentionally supplied.
- Explicit-address policy is conservative: it does not try to classify raw addresses as session or system bus addresses. Instead, it requires `allowAddressBus: true` before any direct address connection can proceed.

### What warrants a second pair of eyes

- Review whether `allowAddressBus` is the desired long-term host policy shape or whether future hosts should classify addresses more precisely.
- Review whether `AllowCallSet` should remain exported or become internal if a future policy builder API replaces direct struct construction.
- Review docs so users understand that omitted `allowCall` inherits the existing call policy while explicit `allowCall: []` denies all method calls.

### What should be done in the future

- Add host-level maximum policy intersection before untrusted JavaScript can widen permissions.
- Consider subcodes such as `ERR_DBUS_POLICY` so policy denials can be distinguished from transport errors.

### Code review instructions

- Start with `pkg/dbuscore/policy.go` and `pkg/dbuscore/policy_test.go` for the policy semantics.
- Review `pkg/dbusgoja/policy.go` and `pkg/dbusgoja/policy_test.go` for JavaScript decoding behavior.
- Review `pkg/modules/dbus/module_test.go` for the default explicit-address denial regression test.
- Review `pkg/modules/dbus/module.go` and `pkg/xgoja/provider/docs/help/03-api-reference.md` for public API docs.
- Validate with:
  - `GOWORK=off go test ./... -count=1`
  - `make xgoja-doctor`
  - `make xgoja-build`
  - `make goreleaser-check`
  - `GOWORK=off goreleaser build --snapshot --clean --single-target`

### Technical details

Commands:

```bash
cd goja-dbus
gh pr view 1 --json title,body,comments,reviews,latestReviews,files,url
gh api repos/go-go-golems/goja-dbus/pulls/1/comments --paginate
GOWORK=off go test ./... -count=1
make xgoja-doctor
make xgoja-build
./dist/goja-dbus verbs examples denied-system-bus
make goreleaser-check
GOWORK=off goreleaser build --snapshot --clean --single-target
git commit -m "Harden DBus policy defaults"
```
