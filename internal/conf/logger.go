package conf

import (
	"github.com/alimy/tryst/cfg"
	"github.com/getsentry/sentry-go"
	sentrylogrus "github.com/getsentry/sentry-go/logrus"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"time"
)

func setupLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(loggerSetting.logLevel())

	cfg.On(cfg.Actions{
		"LoggerFile": func() {
			out := newFileLogger()
			logrus.SetOutput(out)
		},
	})
}

func newFileLogger() io.Writer {
	return &lumberjack.Logger{
		Filename:  loggerFileSetting.SavePath + "/" + loggerFileSetting.FileName + loggerFileSetting.FileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}
}

func setupSentryLogrus(opts sentry.ClientOptions) {
	// 只获取Error以上的级别
	sentryLevels := []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
	sentryHook, err := sentrylogrus.New(sentryLevels, opts)
	if err != nil {
		panic(err)
	}
	logrus.AddHook(sentryHook)
	// 在系统退出前进行刷新
	logrus.RegisterExitHandler(func() { sentryHook.Flush(5 * time.Second) })
}
