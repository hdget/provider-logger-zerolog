package zerolog

import (
	"github.com/hdget/common/types"
	"github.com/hdget/provider-logger-zerolog/pkg"
	"go.uber.org/fx"
)

var Capability = &types.Capability{
	Category: types.ProviderCategoryLogger,
	Name:     types.ProviderNameLoggerZerolog,
	Module: fx.Module(
		string(types.ProviderNameLoggerZerolog),
		fx.Provide(pkg.New),
	),
}
