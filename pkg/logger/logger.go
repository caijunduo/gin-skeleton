package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func New() {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.ISO8601TimeEncoder
	ec.EncodeLevel = zapcore.CapitalLevelEncoder
	ec.EncodeName = zapcore.FullNameEncoder

	zap.ReplaceGlobals(
		zap.New(
			zapcore.NewCore(
				zapcore.NewConsoleEncoder(ec),
				zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
				zap.NewAtomicLevelAt(zap.DebugLevel),
			),
			zap.AddCaller(),
			zap.Development(),
		),
	)
}
