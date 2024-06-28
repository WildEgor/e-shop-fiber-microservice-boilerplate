package configs

import "github.com/google/wire"

// Set contains project configs
var Set = wire.NewSet(
	NewConfigurator,
	NewAppConfig,
	NewLoggerConfig,
	NewProfilerConfig,
)
