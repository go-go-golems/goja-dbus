module github.com/go-go-golems/goja-dbus

go 1.26.1

toolchain go1.26.4

require (
	github.com/dop251/goja v0.0.0-20251103141225-af2ceb9156d7
	github.com/go-go-golems/go-go-goja v0.0.0
	github.com/go-go-golems/logcopter v0.1.0
	github.com/godbus/dbus/v5 v5.2.2
)

require (
	github.com/dlclark/regexp2 v1.11.5 // indirect
	github.com/dop251/base64dec v0.0.0-20231022112746-c6c9f9a96217 // indirect
	github.com/dop251/goja_nodejs v0.0.0-20250409162600-f7acab6894b0 // indirect
	github.com/go-sourcemap/sourcemap v2.1.4+incompatible // indirect
	github.com/google/pprof v0.0.0-20241029153458-d1b30febd7db // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v1.14.32 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/zerolog v1.35.1 // indirect
	golang.org/x/mod v0.36.0 // indirect
	golang.org/x/net v0.55.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/sys v0.45.0 // indirect
	golang.org/x/text v0.37.0 // indirect
	golang.org/x/tools v0.45.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

tool github.com/go-go-golems/logcopter/cmd/logcopter-gen

replace github.com/go-go-golems/go-go-goja => ../go-go-goja
