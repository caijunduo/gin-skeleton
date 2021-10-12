package pkg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
	"os"
)

var (
	Group errgroup.Group
)

func Logger() {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	ec.EncodeLevel = zapcore.CapitalColorLevelEncoder
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
