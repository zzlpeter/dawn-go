package log

import (
	"errors"
	"fmt"
	"github.com/zzlpeter/dawn-go/libs/tomlc"
	"github.com/zzlpeter/dawn-go/libs/utils"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path"
	"runtime"
	"sync"
)

var loggerN *zap.Logger
var once sync.Once
var INFO = "INFO"
var DEBUG = "debug"
var WARNING = "warning"
var ERROR = "error"
var FATAL = "fatal"

func getLogger() {
	once.Do(func() {
		makeLogger()
	})
}

func makeLogger() {
	logFile := tomlc.Config{}.BasicConf()["log"].(string)
	if _, err := os.Stat(logFile); err != nil {
		log.Fatalf("LogPath: %v is not existed", logFile)
	}
	logFile = path.Join(logFile, "root.log")
	hook := lumberjack.Logger{
		Filename:   logFile, 		// 日志文件路径
		MaxSize:    128,            // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 7,              // 日志文件最多保存多少个备份
		MaxAge:     30,             // 文件最多保存多少天
		Compress:   false,          // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)),  // 打印到文件
		atomicLevel,                                          // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	hostName, _ := utils.HostManager{}.LocalHostName()
	hostIp, _ := utils.HostManager{}.LocalIp()
	filed := zap.Fields(zap.String("host_name", hostName), zap.String("host_ip", hostIp))
	// 构造日志
	loggerN = zap.New(core, caller, development, filed)
}

func init() {
	getLogger()
}

// 记录日志
// 具体使用方法参考 dawn-go/middlewares/logger_middleware.go:28
func LogRecord(c *gin.Context, level, msg string, kvs ...interface{}) error {
	if len(kvs) % 2 == 1 {
		return errors.New("kvs should be coupled")
	}
	err, zapFields := makeZapFields(c, kvs...)
	if err != nil {
		return err
	}
	switch level {
	case DEBUG:
		loggerN.Debug(msg, zapFields...)
	case INFO:
		loggerN.Info(msg, zapFields...)
	case WARNING:
		loggerN.Warn(msg, zapFields...)
	case ERROR:
		loggerN.Error(msg, zapFields...)
	case FATAL:
		loggerN.Fatal(msg, zapFields...)
	default:
		loggerN.Info(msg, zapFields...)
	}
	return nil
}

// 组装zap需要的日志参数
func makeZapFields(c *gin.Context, kvs ...interface{}) (error, []zap.Field) {
	var arr []zap.Field
	// 校验kvs是否合法
	l := len(kvs)
	var idx = 0
	for {
		if idx >= l {
			break
		}
		k, ok := kvs[idx].(string)
		if !ok {
			return errors.New(fmt.Sprintf("Type of k: %v is not string", kvs[idx])), arr
		}
		arr = append(arr, zap.Any(k, kvs[idx+1]))
		idx+=2
	}
	// 设置trace-ID
	if c != nil {
		traceId := c.Writer.Header().Get("X-Request-Id")
		ip := c.ClientIP()
		arr = append(arr, zap.String("trace_id", traceId), zap.String("access_ip", ip))
	}
	// 获取调用栈信息
	_, file, line, _ := runtime.Caller(2)
	arr = append(arr, zap.String("file", file), zap.Int("line", line))

	return nil, arr
}