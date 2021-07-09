package main

import (
	"go.uber.org/zap"
	"time"
)

func sugar() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()
	url := "www.google.com"
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second)
	sugar.Info("faile to fetch URL: %s", url)
}

func performant() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		//zap.
		zap.String("url", "www.google.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func main() {
	sugar()
	performant()
}
