package main

import (
	"errors"
	"log"
	"os"
)

var (
	logsFilePath = "/var/log/processd.log"
	lf           *os.File
)

func NewLogger() (*log.Logger, error) {
	var logFile *os.File
	if _, err := os.Stat(logsFilePath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f, err := os.Create(logsFilePath)
			if err != nil {
				return nil, err
			}
			defer f.Close()
			// Ok
		}
		// Internal error
		return nil, err
	}

	logFile, err := os.OpenFile(logsFilePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	lf = logFile

	return log.New(logFile, "[processd] ", log.LUTC), nil
}

func CloseLogger() {
	lf.Close()
}
