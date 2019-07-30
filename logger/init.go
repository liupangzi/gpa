package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
	Log.SetOutput(os.Stderr)
}
