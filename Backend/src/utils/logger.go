package utils

import "go.uber.org/zap"

func GetLogger(config *Config) zap.Logger {
	var cfg zap.Config
	if config.Dgn == "local" {
		cfg = zap.NewDevelopmentConfig()
		cfg.DisableStacktrace = true
	} else {
		cfg = zap.NewProductionConfig()
		cfg.DisableStacktrace = true
		cfg.Encoding = "console"
		cfg.EncoderConfig.TimeKey = "timestamp"
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	logger.Info("Logger initialized")
	return *logger
}
