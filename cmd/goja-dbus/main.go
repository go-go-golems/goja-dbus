package main

import (
	"fmt"
	"os"

	"github.com/go-go-golems/go-go-goja/pkg/xgoja/app"
	"github.com/go-go-golems/go-go-goja/pkg/xgoja/providerapi"
	go_go_goja_core "github.com/go-go-golems/go-go-goja/pkg/xgoja/providers/core"
	dbus "github.com/go-go-golems/goja-dbus/pkg/xgoja/provider"
)

const embeddedRuntimePlanJSON = `{
  "schema": "xgoja/runtime/v2",
  "name": "goja-dbus",
  "app": {
    "name": "goja-dbus"
  },
  "target": {
    "kind": "xgoja",
    "output": "dist/goja-dbus"
  },
  "providers": [
    {
      "id": "go-go-goja-core"
    },
    {
      "id": "dbus"
    }
  ],
  "runtime": {
    "modules": [
      {
        "provider": "dbus",
        "name": "dbus",
        "as": "dbus"
      },
      {
        "provider": "go-go-goja-core",
        "name": "events",
        "as": "events"
      },
      {
        "provider": "go-go-goja-core",
        "name": "timer",
        "as": "timer"
      }
    ]
  },
  "sources": [
    {
      "id": "docs",
      "kind": "help",
      "provider": "dbus",
      "source": "docs"
    },
    {
      "id": "examples",
      "kind": "jsverbs",
      "provider": "dbus",
      "source": "examples"
    }
  ],
  "commands": [
    {
      "id": "eval",
      "type": "builtin.eval",
      "name": "eval"
    },
    {
      "id": "run",
      "type": "builtin.run",
      "name": "run"
    },
    {
      "id": "verbs",
      "type": "builtin.jsverbs",
      "name": "verbs",
      "sources": [
        "examples"
      ]
    }
  ],
  "artifacts": [
    {
      "id": "binary",
      "type": "binary",
      "output": "dist/goja-dbus",
      "sources": [
        "docs",
        "examples"
      ]
    },
    {
      "id": "types",
      "type": "dts",
      "output": "dist/goja-dbus.d.ts",
      "strict": true
    }
  ]
}
`

func main() {
	registry := providerapi.NewProviderRegistry()
	must(go_go_goja_core.Register(registry))
	must(dbus.Register(registry))
	root, err := app.NewRootCommand(app.Options{Providers: registry, RuntimePlanJSON: embeddedRuntimePlanJSON})
	must(err)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
