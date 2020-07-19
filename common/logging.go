package common

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	Log.SetFormatter(&logrus.JSONFormatter{})

	Log.SetOutput(os.Stdout)

	Log.SetLevel(logrus.WarnLevel)
}
