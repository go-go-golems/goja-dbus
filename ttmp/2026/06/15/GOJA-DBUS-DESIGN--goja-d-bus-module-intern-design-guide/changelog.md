# Changelog

## 2026-06-15

- Initial workspace created


## 2026-06-15

Created intern-facing Goja D-Bus design guide, imported /tmp/dbus.md into sources, and recorded investigation diary.

### Related Files

- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/design-doc/01-goja-d-bus-module-intern-design-and-implementation-guide.md — Primary design deliverable
- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/reference/01-investigation-diary.md — Diary of commands
- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/sources/01-dbus.md — Imported user-provided source


## 2026-06-15

Uploaded GOJA DBUS DESIGN GUIDE bundle to reMarkable at /ai/2026/06/15/GOJA-DBUS-DESIGN after fixing PDF render issues.

### Related Files

- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/design-doc/01-goja-d-bus-module-intern-design-and-implementation-guide.md — Rendered in uploaded bundle
- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/reference/01-investigation-diary.md — Includes upload failure and fix notes


## 2026-06-15

Phase 1: normalized goja-dbus repository identity and command path (commit 1d42a9151ffba42e35614d297131e09955c8c3ba).

### Related Files

- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/Makefile — Updated project paths and binary names
- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/README.md — Replaced template README
- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/cmd/goja-dbus-demo/main.go — Renamed demo command
- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/go.mod — Renamed module path


## 2026-06-15

Phase 2: added dbus native module skeleton and typed helper tests (commit 7927235e43d6c9246dd9fa6ef6193433b5e497db).

### Related Files

- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/pkg/dbuscore/types.go — Pure Go typed D-Bus value helpers
- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/pkg/dbusgoja/typed_values.go — JavaScript-facing typed helper exports
- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/pkg/modules/dbus/module.go — NativeModule and TypeScript declaration implementation
- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/pkg/modules/dbus/module_test.go — Runtime require integration tests


## 2026-06-15

Phase 3: added D-Bus connection lifecycle, policy checks, scalar codec, JavaScript builders, and opt-in GetId integration test (commit 7ffee5514db236c03e5bf034e0fe2e8697589fd7).

### Related Files

- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/pkg/dbuscore/bus.go — Connection lifecycle and method call execution
- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/pkg/dbuscore/codec.go — Scalar marshaling and unmarshaling
- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/pkg/dbusgoja/builders.go — JavaScript bus and method builders
- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/pkg/modules/dbus/module_test.go — Promise denial and opt-in integration tests


## 2026-06-15

Phase 4: added D-Bus signal match/subscription support and EventEmitter-based JavaScript listen API (commit a2b3d5c8d07f547046cf409c69d8c0f9ebfdcd36).

### Related Files

- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/pkg/dbuscore/signals.go — Signal match and subscription lifecycle
- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/pkg/dbusgoja/signals.go — EventEmitter-based signal delivery
- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/pkg/modules/dbus/module.go — Signal TypeScript declaration updates


## 2026-06-15

Phase 5: added service export checkpoint and deferred echo-service implementation until runtime cleanup and compound codecs are stronger (commit c5acc524bf2cb3f2d699a574d0ed62000cb07038).

### Related Files

- /home/manuel/workspaces/2026-06-15/goja-dbus/goja-dbus/ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/design-doc/02-service-export-checkpoint.md — Checkpoint decision and future pseudocode

