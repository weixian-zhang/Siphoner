package main

import (
	"os"
	"time"
	"github.com/brianvoe/gofakeit/v5"
	logrus "github.com/sirupsen/logrus"
)

type JsonMsg struct {
	CustomerName string		`fake:"{firstname}"`
	OrderId string			`fake:"{number:1,999}"`
	CustomerId string		`fake:"{number:1,999}"`
	CreditCardNo string		`faker:"CreditCardNumber"`
	Msg string				`fake:"{sentence:5}"`

	
}

func main() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	go func() {

		for {

			msg := JsonMsg{}
			gofakeit.Struct(&msg)
			cf := logrus.WithFields(
				logrus.Fields{
					"CustomerName": msg.CustomerName,
				})
			cf = logrus.WithFields(
				logrus.Fields{
					"CustomerId": msg.CustomerId,
				})
			cf = logrus.WithFields(
				logrus.Fields{
					"orderid": msg.OrderId,
				})
			
			cf.Info(msg)

			time.Sleep(3 * time.Second)

			logrus.Error(gofakeit.Sentence(30))

			time.Sleep(2 * time.Second)
		}
	}()

	select{}
}

