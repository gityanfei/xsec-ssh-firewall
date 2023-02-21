package util

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"xsec-ssh-firewall/settings"
)

var lg *zap.Logger

// InitLogger 初始化Logger
func InitLogger() error {
	//zap.S().Debug(agentConfig.Log.LogConfig.Filename, agentConfig.Log.LogConfig.MaxSize, agentConfig.Log.LogConfig.MaxBackups, agentConfig.Log.LogConfig.MaxAge)
	writeSyncer := getLogWriter(settings.SettingConfig.LogConfig.Filename, settings.SettingConfig.LogConfig.MaxSize, settings.SettingConfig.LogConfig.MaxBackups, settings.SettingConfig.LogConfig.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(settings.SettingConfig.LogConfig.Level))
	if err != nil {
		return err
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	lg = zap.New(core, zap.AddCaller())

	zap.ReplaceGlobals(lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	return err
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	//return zapcore.NewConsoleEncoder(encoderConfig)
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	ws := io.MultiWriter(lumberJackLogger, os.Stdout)
	return zapcore.AddSync(ws)
}
