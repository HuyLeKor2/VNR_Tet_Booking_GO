package logger

import (
	"os"

	"github.com/huyle/go-ecommerce-backend-api/pkg/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZAP struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZAP {
	logLevel := config.Log_level
	//debug -> info -> warn -> error -> fatal -> panic
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	case "panic":
		level = zapcore.PanicLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.File_log_name, // Đường dẫn file log
		MaxSize:    config.Max_size,      // megabytes
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_age,  // days
		Compress:   config.Compress, // disabled by default
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)

	// logger := zap.New(core, zap.AddCaller())
	return &LoggerZAP{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

// Hàm định dạng log (Encoder)
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// Định dạng thời gian: 2024-05-26T16:16:07.877+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"

	// Viết hoa level log: INFO, ERROR
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// Hiển thị file và dòng code gọi log: cli/main.log.go:24
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}
