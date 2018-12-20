package logger

import (
	"os"

	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

// NewLogger create a logger
func NewLogger() *logrus.Logger {
	log.Formatter = new(logrus.JSONFormatter)
	/*
		log.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
		log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	*/
	log.Level = logrus.TraceLevel
	log.Out = os.Stdout

	return log
}
