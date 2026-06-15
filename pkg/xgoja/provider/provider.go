package provider

import (
	"embed"
	"fmt"

	"github.com/dop251/goja_nodejs/require"
	"github.com/go-go-golems/go-go-goja/modules"
	"github.com/go-go-golems/go-go-goja/pkg/tsgen/spec"
	"github.com/go-go-golems/go-go-goja/pkg/xgoja/providerapi"
	_ "github.com/go-go-golems/goja-dbus/pkg/modules/dbus"
)

const PackageID = "dbus"

//go:embed docs/help/*.md
var helpFS embed.FS

//go:embed verbs/*.js
var verbsFS embed.FS

func Register(registry *providerapi.ProviderRegistry) error {
	mod := modules.GetModule("dbus")
	if mod == nil {
		return fmt.Errorf("dbus module is not registered")
	}
	return registry.Package(PackageID,
		nativeModuleEntry(mod),
		providerapi.HelpSource{
			Name:        "docs",
			Description: "goja-dbus getting started, user guide, and API reference help pages.",
			FS:          helpFS,
			Root:        "docs/help",
		},
		providerapi.VerbSource{
			Name:        "examples",
			Description: "Example jsverbs demonstrating goja-dbus helpers and policy behavior.",
			FS:          verbsFS,
			Root:        "verbs",
		},
	)
}

func nativeModuleEntry(mod modules.NativeModule) providerapi.Module {
	return providerapi.Module{
		Name:        mod.Name(),
		DefaultAs:   mod.Name(),
		Description: mod.Doc(),
		TypeScript:  nativeModuleTypeScript(mod),
		NewModuleFactory: func(providerapi.ModuleSetupContext) (require.ModuleLoader, error) {
			return mod.Loader, nil
		},
	}
}

func nativeModuleTypeScript(mod modules.NativeModule) *spec.Module {
	declarer, ok := mod.(modules.TypeScriptDeclarer)
	if !ok {
		return nil
	}
	return declarer.TypeScriptModule()
}
