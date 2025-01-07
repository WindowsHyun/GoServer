package logger

import (
	"GoServer/config"
	"GoServer/config/define"
	"context"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// logger 공용 변수
var (
	logLevel     logrus.Level
	cf           *config.Config
	logsFileName string
)

func checkLogFile(logFileName string, lastTime *time.Time, loggerFile *lumberjack.Logger, logger *logrus.Logger) {
	currentTime := time.Now()
	lastLogHour := lastTime.Truncate(time.Hour)
	currentHour := currentTime.Truncate(time.Hour)

	if !currentHour.Equal(lastLogHour) {
		*lastTime = currentTime
		fileName := getLogFileName(logFileName)
		loggerFile.Close()

		loggerFile = &lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    cf.GetLog().MaxSize,
			MaxBackups: cf.GetLog().MaxBackup, // 최대 3개의 백업 파일 유지
			MaxAge:     cf.GetLog().MaxAge,    // 30일 동안 로그 파일 유지
			Compress:   cf.GetLog().Compress,  // 백업 파일 압축
		}

		// 로그 출력을 새 파일로 설정
		if logger != nil {
			logger.SetOutput(loggerFile)
		}
	}
}

// 현재 날짜를 기준으로 로그 파일 이름 생성
func getLogFileName(logFileName string) string {
	now := time.Now()
	logsDir := cf.GetLog().Fpath
	fileName := now.Format("2006_01_02_15_"+logFileName) + ".log"
	filePath := logsDir + "/" + fileName

	if _, err := os.Stat(logsDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logsDir, os.ModePerm); err != nil {
			return filePath
		}
	}

	return filePath
}

func getLogLevel() logrus.Level {
	if cf.GetLog().Level == "debug" {
		return logrus.DebugLevel
	}
	return logrus.InfoLevel
}

func PanicMessageHandling(ctx context.Context, panicMsg string) {
	logsDir := cf.GetLog().Fpath
	if err := os.MkdirAll(logsDir, define.FolderPermission); err != nil {
		log.Fatalf("Failed to create log directory: %s", err)
	}

	logFileName := fmt.Sprintf("./%s/%s_panic.log", logsDir, time.Now().Format("2006_01_02_15"))
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, define.FilePermission)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)
	stackTrace := string(debug.Stack())
	logger.Printf("[PANIC] %s\n%s\n%s\n", panicMsg, stackTrace, define.PanicLine)
	logPanic(ctx, panicMsg)

	panic(panicMsg)
}
