---
Title: "Getting started with goja-dbus"
Slug: "getting-started"
Short: "Build and run the generated goja-dbus xgoja binary, then try the bundled examples."
Topics:
- goja
- dbus
- xgoja
Commands:
- eval
- run
- verbs
Flags: []
IsTopLevel: true
IsTemplate: false
ShowPerDefault: true
SectionType: Tutorial
---

This tutorial explains how to use the generated goja-dbus xgoja binary. The binary bundles the `require("dbus")` native module, Glazed help pages, and JavaScript verb examples so a user can explore the module without wiring a Go host application by hand.

The generated binary is intended as both a smoke-test host and a documentation carrier. It exposes ordinary xgoja commands such as `eval`, `run`, and `verbs`, while the help system exposes this page, the user guide, and the API reference.

## Build the binary

Build from this workspace with the checked-in xgoja spec:

```bash
make xgoja-build
```

The target runs the sibling `go-go-goja` checkout's `xgoja` command and writes the generated binary back into this repository. If you already have an `xgoja` binary installed, you can also run `xgoja build -f xgoja.yaml` from the repository root.

The default output is:

```bash
dist/goja-dbus-xgoja
```

## Inspect bundled help

Use the help command to confirm the pages are embedded:

```bash
./dist/goja-dbus-xgoja help
./dist/goja-dbus-xgoja help getting-started
./dist/goja-dbus-xgoja help user-guide
./dist/goja-dbus-xgoja help api-reference
```

## Run bundled example verbs

The binary includes example jsverbs from the provider source named `examples`. List them first:

```bash
./dist/goja-dbus-xgoja verbs list
```

Then run the safe examples:

```bash
./dist/goja-dbus-xgoja verbs examples typed-values
./dist/goja-dbus-xgoja verbs examples denied-system-bus
./dist/goja-dbus-xgoja verbs examples get-id-script
```

The `typed-values` verb demonstrates D-Bus typed helper objects. The `denied-system-bus` verb intentionally exercises default policy denial and does not require a running D-Bus daemon. The `get-id-script` verb prints a script you can run on a machine with a session bus.

## Try direct evaluation

Use `eval` when you want to experiment with the module directly:

```bash
./dist/goja-dbus-xgoja eval 'const dbus = require("dbus"); JSON.stringify(dbus.u32(42))'
```

For a real session-bus call, use the script printed by `get-id-script` or write a file and run it with the generated `run` command.

## Troubleshooting

| Problem | Cause | Solution |
| --- | --- | --- |
| `require("dbus")` fails | The generated binary was not built from this repository's `xgoja.yaml` | Rebuild with `GOWORK=off go run ../go-go-goja/cmd/xgoja build -f xgoja.yaml` |
| `system().connect()` rejects | The default policy denies system-bus access | Use session bus examples first; host policy support should be reviewed before enabling system bus access |
| A real D-Bus call fails | No session bus is available or the service is missing | Run on a desktop/session environment and start with `org.freedesktop.DBus.GetId` |
| Help pages are missing | The provider help source was not selected in `xgoja.yaml` | Check the `sources` entry with `kind: help` and provider source `docs` |

## See Also

- `xgoja help user-guide`
- `xgoja help api-reference`
- `xgoja help xgoja-v2-reference`
