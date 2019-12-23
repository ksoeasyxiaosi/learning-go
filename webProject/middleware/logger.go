package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
	"webProject/utils"
)

type LogConfig struct {
	DebugLevel logrus.Level
}

func (l LogConfig) GetFileName() string {
	return "runtime/log/test.log"
}

var Config LogConfig

func LoggerToFile() gin.HandlerFunc {

	fileName := Config.GetFileName()

	// 打开日志文件的句柄
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	logger := logrus.New()

	// 设置输出
	logger.Out = src

	// 日志级别
	logger.SetLevel(logrus.DebugLevel)

	logWriter, err := rotatelogs.New(
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	// 日志格式化
	lfHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增Hook
	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 执行操作
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 消耗的时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUrl := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求ip
		clientIp := c.ClientIP()

		// 其它信息
		msg := strings.Join(utils.Log().Msg, "|")

		// 日志格式

		logger.Infof("| %3d | %13v | %15s | %s | %s | %s",
			statusCode,
			latencyTime,
			clientIp,
			reqMethod,
			reqUrl,
			msg,
		)
	}
}
