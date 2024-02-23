package async

import (
	"github.com/sirupsen/logrus"
)

func SafeAsyncExecute(logger *logrus.Logger, executor func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Warnf("Async execute error: %v", err)
			}
		}()
		executor()
	}()
}
