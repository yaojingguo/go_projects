package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
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
	sugar.Info("failed to fetch URL: ", url)
	sugar.Info("one ", "two ", "three")
	sugar.Infof("%s %s %s", "one", "two", "three")
	fmt.Println()
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
	fmt.Println()
}

func main() {
	sugar()
	performant()
}
