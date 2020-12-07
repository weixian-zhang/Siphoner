package main

import(
	"os"
	logrus "github.com/sirupsen/logrus"
)

type Logger interface { 
	Init()
	Info(message string)
	Err(err error)
}

type StdLog struct {}

func (*StdLog) Init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

func (*StdLog) Info(message string) {
	if len(message) > 0 {
		logrus.Info(message)
	}
}

func (*StdLog) Err(err error) {
	if err != nil {
		logrus.Error(err)
	}
	

	//log signin in user field when Authn module is up
	// logrus.WithFields(logrus.Fields{
    //     "user": "admin",
    // }).Info("Some interesting info")
}

