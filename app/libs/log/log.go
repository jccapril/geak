package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

/*
	format : "json" or "console"
*/
type Config struct {
	Level         string `mapstructure:"level" json:"level" toml:"level"`
	Prefix 		  string `mapstructure:"prefix" json:"prefix" toml:"prefix"`
	Format        string `mapstructure:"format" json:"format" toml:"format"`
	Dir           string `mapstructure:"dir" json:"dir"  toml:"dir"`
	LinkName      string `mapstructure:"link-name" json:"link-name" toml:"link-name"`
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" toml:"show-line"`
	LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" toml:"log-in-console"`
}

var(
	logger 	*zap.Logger
	config *Config
)

func Init(cfg *Config){

	config = cfg
	if ok, _ := PathExists(cfg.Dir); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", cfg.Dir)
		_ = os.Mkdir(cfg.Dir, os.ModePerm)
	}

	encoder := getEncoder(cfg.Format == "json")

	var err error
	var ws zapcore.WriteSyncer

	var level = new(zapcore.Level)
	err = level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}

	ws, err = GetWriteSyncer(cfg) // 使用file-rotatelogs进行日志分割
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, ws, level)
	if cfg.ShowLine {
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	}else {
		logger = zap.New(core)
	}

}
func getEncoder(isJson bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = CustomTimeEncoder // 时间字符串
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder // 函数调用
	if isJson {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(config.Prefix + "2006/01/02 - 15:04:05.000"))
}


func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...) // logger.go
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg,fields...)
}

