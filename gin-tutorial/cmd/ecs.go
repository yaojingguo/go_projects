package main

import (
	"go.uber.org/zap"
	"os"
	"errors"
)
import "go.elastic.co/ecszap"

func main() {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())


	// Add fields and a logger name
	logger = logger.With(zap.String("custom", "foo"))
	logger = logger.Named("mylogger")

	// Use strongly typed Field values
	logger.Info("some logging info",
		zap.Int("count", 17),
		zap.Error(errors.New("boom")))
}
