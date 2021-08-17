package logger

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    "os"
    "skeleton/configs"
    "time"
)

func New() {
    if !configs.Config.Logger.Mode {
        return
    }
    ec := zap.NewProductionEncoderConfig()
    ec.EncodeTime = zapcore.ISO8601TimeEncoder
    ec.EncodeLevel = zapcore.CapitalLevelEncoder
    ec.EncodeName = zapcore.FullNameEncoder

    now := time.Now()

    hook := lumberjack.Logger{
        Filename: fmt.Sprintf("%s/skeleton-%04d%02d%02d.log",
            configs.Config.Logger.SavePath,
            now.Year(),
            now.Month(),
            now.Day(),
        ),
        MaxSize:    configs.Config.Logger.MaxSize,
        MaxAge:     configs.Config.Logger.MaxAge,
        MaxBackups: configs.Config.Logger.MaxBackups,
        Compress:   configs.Config.Logger.Compress,
    }

    var writes = []zapcore.WriteSyncer{zapcore.AddSync(&hook)}

    if gin.IsDebugging() {
        writes = append(writes, zapcore.AddSync(os.Stdout))
    }

    zap.ReplaceGlobals(
        zap.New(
            zapcore.NewCore(
                zapcore.NewJSONEncoder(ec),
                zapcore.NewMultiWriteSyncer(writes...),
                zap.NewAtomicLevelAt(zap.DebugLevel),
            ),
            zap.AddCaller(),
            zap.Development(),
        ),
    )
}
