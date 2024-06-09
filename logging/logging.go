package logging

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path"
	"sample/util"
	"strings"
)

var (
	instances map[string]*MyLogger = map[string]*MyLogger{}
)

const (
	DEFAULT_LOGGER_NAME = "default"
	DEFAULT_LOG_FILE    = "./log/app.log"
	DEFAULT_LOG_LEVEL   = "INFO"
)

func init() {
	createLogger()
}

func createLogger() {
	conf := map[string]util.LogConfig{}
	err := util.LoadConfig("log-config.yml", conf)

	if err != nil {
		createDefaultLogger()
		logger := instances[DEFAULT_LOGGER_NAME]
		logger.Eerror(err)
		return
	}

	if _, ok := conf[DEFAULT_LOGGER_NAME]; !ok {
		createDefaultLogger()
	}

	for k, v := range conf {
		handler := createHandler(v)
		instances[k] = &MyLogger{Logger: slog.New(handler)}
	}
}

func createDefaultLogger() {
	handler := createHandler(util.LogConfig{LogFile: DEFAULT_LOG_FILE, LogLevel: DEFAULT_LOG_LEVEL})
	instances[DEFAULT_LOGGER_NAME] = &MyLogger{Logger: slog.New(handler)}
}

type MyLogger struct {
	*slog.Logger
}

func (myLogger *MyLogger) Eerror(err error) {
	myLogger.Logger.Error(fmt.Sprintf("%+v", err))
}

func (myLogger *MyLogger) Infof(format string, obj any) {
	myLogger.Logger.Info(fmt.Sprintf(format, obj))
}

func (myLogger *MyLogger) Debugf(format string, obj any) {
	myLogger.Logger.Debug(fmt.Sprintf(format, obj))
}

func GetLogger(name string) *MyLogger {
	logger, ok := instances[name]
	if !ok {
		logger = instances[DEFAULT_LOGGER_NAME]
		logger.Warn("Logger Name[" + name + "] is not found")
	}

	return logger
}

func GetDefaultLogger() *MyLogger {
	return GetLogger(DEFAULT_LOGGER_NAME)
}

func createHandler(conf util.LogConfig) slog.Handler {
	filePath := conf.LogFile
	dir := path.Dir(filePath)
	if _, err := os.Stat(dir); errors.Is(err, fs.ErrNotExist) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}

	logfile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logfile)

	handler := slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{AddSource: true, Level: logLevel(conf.LogLevel)})
	return handler
}

func logLevel(level string) slog.Level {

	switch strings.ToUpper(level) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		panic("Invalid Log Level")
	}
}
