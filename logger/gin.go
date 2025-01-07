package logger

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ginLogger      *logrus.Logger
	ginLogFile     *lumberjack.Logger
	lastGinLogTime time.Time
	skipPaths      []string
)

const (
	errorStatusCode   = 500
	warningStatusCode = 400
)

func ginLoggerInit() {
	ginFileName := getLogFileName("Gin")
	logLevel = getLogLevel()
	skipPaths = append(skipPaths, "/health")

	ginLogFile = &lumberjack.Logger{
		Filename:   ginFileName,
		MaxSize:    cf.GetLog().MaxSize,
		MaxBackups: cf.GetLog().MaxBackup, // 최대 3개의 백업 파일 유지
		MaxAge:     cf.GetLog().MaxAge,    // 30일 동안 로그 파일 유지
		Compress:   cf.GetLog().Compress,  // 백업 파일 압축
	}

	ginLogger = logrus.New()
	ginLogger.SetOutput(ginLogFile)

	ginLogger.SetLevel(logLevel)
	ginLogger.SetFormatter(&CustomFormatter{})

	lastGinLogTime = time.Now()
}

// CustomFormatter는 logrus의 Formatter를 커스터마이즈한 구조체입니다.
type CustomFormatter struct{}

// Format은 로그를 원하는 형식으로 포맷합니다.
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04:05.000")
	level := entry.Level.String()
	method := entry.Data["method"].(string)
	status := strconv.Itoa(entry.Data["status"].(int))
	path := entry.Data["path"]
	duration := entry.Data["duration"]
	ip := entry.Data["ip"]

	logLine := fmt.Sprintf("[%s][%s][%s][%s][%s][%s][%s]\n", timestamp, level, method, status, duration, ip, path)
	return []byte(logLine), nil
}

// GinLogrus gin의 로깅 미들웨어로 logrus로 로그를 출력하도록 설정합니다.
func GinLogrus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if shouldSkipPath(c.Request.URL.Path) {
			c.Next()
			return
		}

		checkLogFile("Gin", &lastGinLogTime, ginLogFile, ginLogger)
		startTime := time.Now()
		// gin에서 로그 레벨을 가져와 logrus 레벨로 변환합니다.
		level := logrus.InfoLevel
		if c.Writer.Status() >= errorStatusCode {
			level = logrus.ErrorLevel
		} else if c.Writer.Status() >= warningStatusCode {
			level = logrus.WarnLevel
		}
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		// gin의 로그 메시지를 logrus로 전달합니다.
		ginLogger.WithFields(logrus.Fields{
			"time":     time.Now().Format("2006-01-02 15:04:05.000"),
			"level":    level,
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"status":   c.Writer.Status(),
			"duration": latencyTime,
			"ip":       c.ClientIP(),
		}).Log(level, "")
	}
}

func shouldSkipPath(path string) bool {
	for _, skipPath := range skipPaths {
		if path == skipPath {
			return true
		}
	}
	return false
}

func PanicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if panicMsg := recover(); panicMsg != nil {
				PanicMessageHandling(c.Request.Context(), fmt.Sprintf("%v", panicMsg))
			}
		}()
		c.Next()
	}
}
