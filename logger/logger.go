package logger

import (
	"GoServer/config"
	"GoServer/config/define"
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	logrus "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	callerSkipStep = 2
)

type logFunc func(context.Context, string)

var (
	mainLogger      *logrus.Logger
	mainLogFile     *lumberjack.Logger
	logInfo         logFunc
	logWarn         logFunc
	logError        logFunc
	logFatal        logFunc
	logDebug        logFunc
	logPanic        logFunc
	lastMainLogTime time.Time
)

func InitLogger(cfg *config.Config, fileName string) (err error) {
	cf = cfg
	logsFileName = fileName
	loggerInit()
	ginLoggerInit()
	return
}

func loggerInit() {
	logFileName := getLogFileName(logsFileName)
	logLevel = getLogLevel()

	mainLogFile = &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    cf.GetLog().MaxSize,
		MaxBackups: cf.GetLog().MaxBackup, // 최대 3개의 백업 파일 유지
		MaxAge:     cf.GetLog().MaxAge,    // 30일 동안 로그 파일 유지
		Compress:   cf.GetLog().Compress,  // 백업 파일 압축
	}

	mainLogger = logrus.New()
	mainLogger.SetOutput(mainLogFile)
	mainLogger.SetFormatter(&customFormatter{})
	mainLogger.SetLevel(logLevel)

	logInfo = generateLogFunc(logrus.InfoLevel)
	logWarn = generateLogFunc(logrus.WarnLevel)
	logError = generateLogFunc(logrus.ErrorLevel)
	logFatal = generateLogFunc(logrus.FatalLevel)
	logDebug = generateLogFunc(logrus.DebugLevel)
	logPanic = generateLogFunc(logrus.PanicLevel)
	lastMainLogTime = time.Now()
}

// 로그 함수 생성
func generateLogFunc(level logrus.Level) logFunc {
	return func(ctx context.Context, msg string) {
		checkLogFile(logsFileName, &lastMainLogTime, mainLogFile, mainLogger)
		_, file, line, _ := runtime.Caller(callerSkipStep)

		// ErrorResponseLog 해당 함수 호출로 인한 logger.go 로 찍히는거 방지 코드
		result := strings.Contains(file, "logger.go")
		if result {
			_, newFile, newLine, _ := runtime.Caller(callerSkipStep + 1)
			file = newFile
			line = newLine
		}

		loggerPrint, ok := ctx.Value(define.ContextLoggerPrint).(bool)
		if ok && loggerPrint {
			log.Println(msg)
		}

		realIP := ""
		realIP, ok = ctx.Value(define.ContextUserRealIP).(string)
		if !ok {
			realIP = ""
		}

		entry := mainLogger.WithFields(logrus.Fields{
			"code_line": fmt.Sprintf("%s:%d", filepath.Base(file), line),
			"ip":        realIP,
		})

		switch level {
		case logrus.InfoLevel:
			entry.Info(msg)
		case logrus.WarnLevel:
			entry.Warn(msg)
		case logrus.ErrorLevel:
			entry.Error(msg)
		case logrus.FatalLevel:
			entry.Fatal(msg)
		case logrus.DebugLevel:
			entry.Debug(msg)
		case logrus.PanicLevel:
			entry.Panic(msg)
		case logrus.TraceLevel:
			entry.Trace(msg)
		}
	}
}

// 로그 메시지를 출력하는 함수 (Debug 레벨)
func LogDebug(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	logDebug(ctx, msg)
}

// 로그 메시지를 출력하는 함수 (Info 레벨)
func LogInfo(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	logInfo(ctx, msg)
}

// 로그 메시지를 출력하는 함수 (Warning 레벨)
func LogWarn(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	logWarn(ctx, msg)
}

// 로그 메시지를 출력하는 함수 (Error 레벨)
func LogError(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	logError(ctx, msg)
}

// 로그 메시지를 출력하는 함수 (Fatal 레벨)
func LogFatal(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	logFatal(ctx, msg)
}

// 로그 메시지를 출력하는 함수 (Panic 레벨)
func LogPanic(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	logPanic(ctx, msg)
}

// 로그 메세지 및 에러 던지는 함수
func ErrorResponseLog(c *gin.Context, err error, status int, logMessage string) {
	if err != nil {
		LogError(c.Request.Context(), "%s, err: %s", logMessage, err.Error())
		c.JSON(http.StatusForbidden, gin.H{"status": status, "msg": logMessage})
	} else {
		LogWarn(c.Request.Context(), "%s, Err: error nil", logMessage)
		c.JSON(http.StatusForbidden, gin.H{"status": status, "msg": logMessage})
	}
}

// 커스텀 로그 포맷터 정의
type customFormatter struct {
	TimestampFormat string
}

func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	level := entry.Level.String()
	codeLine := entry.Data["code_line"].(string)
	userIp := entry.Data["ip"].(string)
	msg := entry.Message

	logLine := fmt.Sprintf("[%s][%s][%s][%s][%s]\n", timestamp, level, codeLine, userIp, msg)
	return []byte(logLine), nil
}

func HandlePanicWrapper(f func()) {
	defer func() {
		if panicMsg := recover(); panicMsg != nil {
			PanicMessageHandling(context.Background(), fmt.Sprintf("%v", panicMsg))
		}
	}()
	f()
}
