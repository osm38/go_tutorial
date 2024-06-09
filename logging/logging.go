package logging

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
)

var instance *MyLogger
var once sync.Once

type MyLogger struct {
	*slog.Logger
}

func (myLogger *MyLogger) Eerror(err error) {
	myLogger.Logger.Error(fmt.Sprintf("%+v", err))
}

func GetLogger() *MyLogger {
	if instance == nil {
		once.Do(func() {
			_logger := slog.New(createHandler())
			instance = &MyLogger{Logger: _logger}
		})
	}
	// once.Do(func() {
	// 	logger = slog.New(createHandler())
	// })
	// return logger
	return instance
}

func createHandler() slog.Handler {
	dir := "./log/"
	if _, err := os.Stat(dir); errors.Is(err, fs.ErrNotExist) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}

	filePath := filepath.Join(dir, "app.log")

	logfile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logfile)

	handler := slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo})
	return handler
}
