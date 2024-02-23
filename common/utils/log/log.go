package log

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	rotatelogs "github.com/Mu-Exchange/Mu-End/common/utils/log/file-rotatelog"
)

var defaultFormatter = &Formatter{
	FieldsOrder:     []string{"request_id", "http_code", "total_time", "ip", "method", "uri", "err_code", "err_msg"},
	TimestampFormat: time.RFC3339,
	CallerFirst:     true,
	NoFieldsBracket: true,
	CustomCallerFormatter: func(f *runtime.Frame) string {
		_, filename := filepath.Split(f.File)
		return fmt.Sprintf("%18s:%-4d", filename, f.Line)
	},
}

func New(level string, filePath string, fileName string, maxSize int64, maxAge time.Duration, rotationTime time.Duration) *logrus.Logger {
	log := logrus.New()
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		panic(err)
	}

	log.SetFormatter(defaultFormatter)
	log.SetReportCaller(true)
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.ErrorLevel
	}
	log.SetLevel(lvl)
	log.SetOutput(os.Stdout)

	log.AddHook(newRotateHook(filePath, fileName, maxSize, maxAge, rotationTime, defaultFormatter))
	return log
}

func newRotateHook(logPath string, logFileName string, maxSize int64, maxAge time.Duration, rotationTime time.Duration, formatter logrus.Formatter) *lfshook.LfsHook {
	baseLogName := path.Join(logPath, logFileName)
	if maxSize < 1024 {
		maxSize = 10 * 1024 * 1024
	}
	writer, err := rotatelogs.New(
		baseLogName+"_%Y%m%d%H%M%S.log",
		rotatelogs.WithLinkName(baseLogName),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithMaxSize(maxSize),
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithRedirectStdErr(true),
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
		return nil
	}

	return lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, formatter)
}
