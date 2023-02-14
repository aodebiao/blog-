package log

/*
系统环境变量:
	 需要在 main 包的 init 函数中使用 os.Setenv() 设置环境变量
   GIN_LOG_PATH: 设置文件路径及其名称,例如: ./log/server.log 或者 /var/log/demo_server.log
   GIN_DEBUG: 设置是否开启开发模式，开启后会输出堆栈跟踪信息

依赖:
	gopkg.in/natefinch/lumberjack.v2
	go.uber.org/zap
*/

import (
	"aodebiao/conf"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	once    sync.Once
)

var (
	Debug func(msg string, fields ...zap.Field)
	Info  func(msg string, fields ...zap.Field)
	Warn  func(msg string, fields ...zap.Field)
	Error func(msg string, fields ...zap.Field)
	Fatal func(msg string, fields ...zap.Field)
)

func InitLogger() {
	once.Do(func() {
		logger = logConfig()
	})

	Debug = logger.Debug
	Info = logger.Info
	Warn = logger.Warn
	Error = logger.Error
	Fatal = logger.Fatal
}

// LogConfig 日志配置
func logConfig() *zap.Logger {

	filePath :=conf.Setting.Log.Path
	//level := conf.Setting.Log.Level
	if filePath == "" {
		filePath = "./logs/server.log"
	}

	loggerFileWriter := lumberjack.Logger{
		Filename:   filePath, // 日志文件路径
		MaxSize:    10,       // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 5,        // 日志文件最多保存多少个备份
		MaxAge:     7,        // 文件最多保存多少天
		Compress:   true,     // 是否压缩
	}

	// 日志文件输出配置
	fileEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,                        // 全大写日志等级标识
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"), // 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 终端输出配置
	stdEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	fileEncoder := zapcore.NewJSONEncoder(fileEncoderConfig)  // 设置日志文件内容编码格式：json格式
	stdEncoder := zapcore.NewConsoleEncoder(stdEncoderConfig) // 终端格式

	// 创建写入的目标 writer
	fileWriter := zapcore.NewMultiWriteSyncer(zapcore.AddSync(&loggerFileWriter))
	stdWriter := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))

	// 设置同时写入文件和终端
	//logWriter := zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook), zapcore.AddSync(os.Stdout))

	// 定义日志级别
	debugLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel
	})

	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.InfoLevel
	})

	// 组合所有配置，分别设置写入文件的参数和显示到终端的参数
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileWriter, infoLevel),
		zapcore.NewCore(stdEncoder, stdWriter, debugLevel),
	)

	if os.Getenv("GIN_DEBUG") == "true" {
		// 开启开发模式，堆栈跟踪
		caller := zap.AddCaller()
		// // 开启文件及行号
		development := zap.Development()
		// 构造日志对象
		return zap.New(core, caller, development)
	} else {
		return zap.New(core)
	}
}