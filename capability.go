package zerolog

import (
	"github.com/hdget/common/types"
	"github.com/hdget/provider-logger-zerolog/pkg"
	"go.uber.org/fx"
)

const (
	providerName = "logger-zerolog"
)

var Capability = &types.Capability{
	Category: types.ProviderCategoryLogger,
	Name:     providerName,
	Module: fx.Module(
		providerName,
		fx.Provide(pkg.New),
	),
}
