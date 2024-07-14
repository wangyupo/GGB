package core

import (
	"fmt"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

// Zap 自定义zap日志
func Zap() (logger *zap.Logger) {
	// 1-创建日志存放目录
	ok, _ := utils.PathExists(global.GGB_CONFIG.Zap.Director)
	if !ok {
		fmt.Printf("初始化日志存放目录 %v\n", global.GGB_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GGB_CONFIG.Zap.Director, os.ModePerm)
	}

	// 2-获取所有日志等级
	levels := getZapLevels()
	length := len(levels)

	// 3-创建不同等级的日志核心
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := newZapCore(levels[i])
		cores = append(cores, core)
	}

	// 4-将多个日志核心合并成一个，最终生成自定义的zap logger
	logger = zap.New(zapcore.NewTee(cores...))
	if global.GGB_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller()) // AddCaller() 添加句柄，显示文件名和行号
	}

	return logger
}

// GetZapLevels 读取yaml配置，返回从 zap.level 到 FatalLevel 组成的日志级别切片
func getZapLevels() []zapcore.Level {
	// 1-新建一个空切片，长度为0，容量为7
	levels := make([]zapcore.Level, 0, 7)

	// 2-读取 yaml 配置中的日志等级，解析日志级别（将字符串形式的日志级别解析为相应的 zapcore.Level 类型）
	level, err := zapcore.ParseLevel(global.GGB_CONFIG.Zap.Level)
	if err != nil {
		level = zapcore.DebugLevel
	}

	// 3-生成日志级别切片，从解析到的 level 开始，到 zapcore.FatalLevel 结束
	for ; level < zapcore.FatalLevel; level++ {
		levels = append(levels, level)
	}

	return levels
}

// 新建zap日志核心
func newZapCore(level zapcore.Level) zapcore.Core {
	core := zapcore.NewCore(getEncoder(), getWriteSync(level), getLevelEnabler(level))
	return core
}

// zap日志编码器（负责将日志编译成指定格式）
func getEncoder() zapcore.Encoder {
	// 1-复制一份 zap 的预设编码器 zap.NewProductionEncoderConfig()，以供自定义
	config := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "msg",
		StacktraceKey: global.GGB_CONFIG.Zap.StacktraceKey, // 堆栈跟踪信息的key
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(global.GGB_CONFIG.Zap.Prefix + t.Local().Format(time.DateTime))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 2-根据 yaml 配置返回 json 或 控制台 格式的编码
	if global.GGB_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(config) // 输出json格式：{"level":"error","time":"2024-06-24 09:56:09","msg":"222"}
	}
	return zapcore.NewConsoleEncoder(config) // 输出控制台格式：2024-06-23 15:55:59	error	222
}

// zap日志写入器（定义日志的信息应该写到哪里）
func getWriteSync(level zapcore.Level) zapcore.WriteSyncer {
	// 1-根据 yaml 配置新建日志存放目录
	logDirPath := fmt.Sprintf("./%s", global.GGB_CONFIG.Zap.Director)
	_ = os.MkdirAll(logDirPath, os.ModePerm)

	// 2-按当前日期组织日志存放目录（当天日志存放在当然日期文件夹下，如：./2006-01-02/error.log）
	date := time.Now().Local().Format(time.DateOnly)
	fileName := filepath.Join(logDirPath, date, fmt.Sprintf("./%s.log", level.String()))

	// 3-使用 lumberjack 分割日志
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,                         //日志文件存放目录
		MaxSize:    global.GGB_CONFIG.Zap.MaxSize,    //文件大小限制,单位M
		MaxBackups: global.GGB_CONFIG.Zap.MaxBackups, //最大保留日志文件数量
		MaxAge:     global.GGB_CONFIG.Zap.MaxAge,     //日志文件保留天数
		Compress:   global.GGB_CONFIG.Zap.Compress,   //是否压缩处理
	}

	// 4-读取 yaml 配置，判断是否将日志输出在控制台上
	if global.GGB_CONFIG.Zap.LogInConsole {
		// zapcore.NewMultiWriteSyncer 是 zap 日志库中的一个工具函数，用于创建一个多重写入器（multi-writer）。
		// 这个多重写入器可以将日志数据同时写入多个 WriteSyncer，实现日志的多目标输出功能。常见的用例是将日志同时写
		// 入到控制台（标准输出）和文件。
		multiSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberjackLogger))
		return multiSyncer
	}

	return zapcore.AddSync(lumberjackLogger)
}

// zap日志级别使能器（保证对应级别的日志写入对应名称的日志文件中，如：error 日志写入 error.log 文件中）
func getLevelEnabler(level zapcore.Level) zapcore.LevelEnabler {
	LevelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == level
	})
	return LevelEnabler
}
