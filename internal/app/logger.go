package app

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")

	return arr[len(arr)-1]
}

func InitLogger() error {
	logrus.SetReportCaller(true)

	logLvl, err := logrus.ParseLevel(CFG.Logger.LogLevel)
	if err != nil {
		return err
	}

	formater := &logrus.TextFormatter{
		TimestampFormat:        "02-02-2006 15:44:45",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		ForceColors:            true,
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}

	logrus.SetLevel(logLvl)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(formater)

	return nil
}
