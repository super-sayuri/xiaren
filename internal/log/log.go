package log

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync"
	"time"

	"xiaren/internal/config"
	"xiaren/internal/constant"
	"xiaren/internal/saving"
)

var _logger *slog.Logger

func clearLogFile() error {
	gData := saving.GetGData()
	err := gData.DeleteObject(saving.LOG_OBJ)
	if err != nil {
		return err
	}
	return nil
}

func createLogFile() (string, error) {
	logFile := fmt.Sprintf("v%s_%s.log", config.GetConfig().Version, time.Now().Format("20060102150405"))
	gData := saving.GetGData()
	err := gData.SaveObjectProp(saving.LOG_OBJ, logFile, []byte(""))
	if err != nil {
		return "", err
	}
	return gData.ObjectPropPath(saving.LOG_OBJ, logFile), nil
}

func InitLogger(env string) error {
	return sync.OnceValue(func() error {
		var logLevel slog.Level
		var writer io.Writer
		// remove previous log
		err := clearLogFile()
		if err != nil {
			return err
		}
		// add new one
		logfilePath, err := createLogFile()
		if err != nil {
			return err
		}
		logFile, err := os.OpenFile(logfilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		if env == "release" {
			logLevel = slog.LevelWarn
			writer = logFile
		} else {
			logLevel = slog.LevelDebug
			writer = io.MultiWriter(logFile, os.Stdout)
		}

		handler := slog.NewTextHandler(
			writer, &slog.HandlerOptions{
				Level: logLevel,
			})
		_logger = slog.New(handler)
		slog.SetDefault(_logger)
		slog.Debug("Logger initialized", "env", env, "level", logLevel)
		return nil
	})()
}

func GetLog(ctx context.Context) slog.Logger {
	logCtx := ctx.Value(constant.CTX_LOGGER)
	logger := _logger.With("ctx", logCtx)
	return *logger
}
