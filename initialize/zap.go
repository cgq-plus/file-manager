package initialize

import (
	"file-manager/global"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"runtime/debug"
	"strings"
	"time"
)

func Zap() (logger *zap.Logger) {
	hook := lumberjack.Logger{
		Filename:   path.Join(global.CONFIG.Zap.Path, time.Now().Format("2006-01-02")+".log"), // 日志文件路径
		MaxSize:    10,                                                                        // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 50,                                                                        // 日志文件最多保存多少个备份
		MaxAge:     10,                                                                        // 文件最多保存多少天
		Compress:   true,                                                                      // 是否压缩
		LocalTime:  true,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)
	// 开启文件及行号
	development := zap.Development()
	// 构造日志
	logger = zap.New(core, development)
	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	logger.Info("log 初始化成功")
	return logger
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestPath := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)

		logger.Info(requestPath,
			zap.Int("status", c.Writer.Status()),
			zap.String("log_type", "runtime"),
			zap.String("method", c.Request.Method),
			zap.String("path", requestPath),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.String("log_type", "runtime"),
						zap.String("request", string(httpRequest)),
						zap.Any("error", err),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				if stack {
					debug.PrintStack()
					logger.Error("[Recovery from panic]",
						zap.String("log_type", "runtime"),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
						zap.Any("error", errorToString(err)),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.String("request", string(httpRequest)),
						zap.Any("error", errorToString(err)),
					)
				}
				c.JSON(http.StatusInternalServerError, gin.H{
					"data": nil,
					"msg":  "系统异常",
				})
			}
		}()
		c.Next()
	}
}

func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
