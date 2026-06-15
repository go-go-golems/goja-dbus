# goja-dbus

`goja-dbus` is a planned native module for [go-go-goja](https://github.com/go-go-golems/go-go-goja) that exposes D-Bus client and service capabilities to JavaScript through `require("dbus")`.

The implementation goal is strict and boring on the Go side:

- JavaScript describes bus intent, destinations, object paths, interfaces, signatures, handlers, and lifecycle.
- Go owns D-Bus execution, marshaling, policy enforcement, cleanup, and Goja runtime scheduling.
- All JavaScript callbacks, `goja.Value` creation, and Promise settlement must happen on the go-go-goja runtime owner.

See the docmgr ticket at `ttmp/2026/06/15/GOJA-DBUS-DESIGN--goja-d-bus-module-intern-design-guide/` for the detailed intern-facing design and implementation guide.

## Development

```bash
GOWORK=off go test ./...
GOWORK=off go generate ./...
```

The demo command currently exists only as a placeholder:

```bash
go run ./cmd/goja-dbus-demo
```
