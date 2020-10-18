package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

type Logger struct {
}

func NewLogger(format log.Formatter, logFile string) *log.Logger {
	var file, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}

	mw := io.MultiWriter(os.Stdout, file)

	logger := log.New()
	logger.SetOutput(mw)
	logger.SetFormatter(format)

	return logger
}
