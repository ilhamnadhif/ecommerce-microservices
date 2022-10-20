package helper

import "github.com/sirupsen/logrus"

func LogFatalIfError(err error) {
	if err != nil {
		logrus.Fatal(err.Error())
	}
}
