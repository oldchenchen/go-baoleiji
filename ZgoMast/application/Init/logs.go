package Init

import (
	"ZgoMast/application/Config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(cfg *Config.LogConfig) (err error) {
	// 日志初始化
	// 定制日志的格式
	writeSyncer := getLogWriter(cfg.Filename, cfg.Maxsize, cfg.MaxBackups, cfg.Maxage)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	if err = l.UnmarshalText([]byte(cfg.Level)); err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)
	Logger = zap.New(core, zap.AddCaller())
	// logger, _ := zap.NewProduction()  // 用于项目生产阶段，格式: json【适合用于集成到第三方日志分析系统中的】
	// logger, _ := zap.NewDevelopment()    // 用于项目开发阶段，格式: 普通文本格式【适合在终端查看】
	// 替换了zap的全局日志配置
	zap.ReplaceGlobals(Logger)
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
