package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.Out = os.Stdout

	Log.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}
	Log.SetLevel(logrus.InfoLevel)
}
