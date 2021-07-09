package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	defer logger.Sync()

	url := "www.google.com"

	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
