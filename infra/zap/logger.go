package zap

import (
	"go.uber.org/zap"
	"log"
)

// Logger is a global logger
var Logger *zap.Logger

// NewLogger returns a zap logger
func NewLogger(opts ...zap.Option) (*zap.Logger, error) {
	zapCfg := zap.Config{
		Level:            zap.NewAtomicLevel(),
		Encoding:         "json",
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}

	return zapCfg.Build(opts...)
}

func init() {
	logger, err := NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	Logger = logger
}
