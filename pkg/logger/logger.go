package logger

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    "os"
    "time"
)

func New() (err error) {
    ec := zap.NewProductionEncoderConfig()
    ec.EncodeTime = zapcore.ISO8601TimeEncoder
    ec.EncodeLevel = zapcore.CapitalLevelEncoder
    ec.EncodeName = zapcore.FullNameEncoder

    now := time.Now()

    hook := lumberjack.Logger{
        Filename: fmt.Sprintf("%s/skeleton-%04d%02d%02d.log",
            viper.GetString("logger.savePath"),
            now.Year(),
            now.Month(),
            now.Day(),
        ),
        MaxSize:    viper.GetInt("logger.maxSize"),
        MaxAge:     viper.GetInt("logger.maxAge"),
        MaxBackups: viper.GetInt("logger.maxBackups"),
        Compress:   viper.GetBool("logger.compress"),
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

    return
}
