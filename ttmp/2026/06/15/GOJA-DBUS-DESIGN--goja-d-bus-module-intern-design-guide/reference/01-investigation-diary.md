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
