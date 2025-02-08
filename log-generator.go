package main

import (
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

type Loglevel int

const (
	Error Loglevel = iota
	Debug
	Info
	Warn
	Security
	Severe
)

type logitem struct {
	loglevel Loglevel
	message  string
}

var services = []string{
	"web-ui", "bank", "mysqld", "nfs",
}
var events = []string{
	"startup", "login", "payment", "sync",
}

var messages = []logitem{
	logitem{Debug, "User login successfull"},
	logitem{Debug, "Payment processed succesfully"},
	logitem{Severe, "Database connection error"},
	logitem{Security, "Unauthorized access attempt detected"},
	logitem{Debug, "Session expired"},
	logitem{Info, "Cache invalidated"},
	logitem{Warn, "Network slowdown detected"},
	logitem{Error, "service restarting due to memory pressure"},
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.WithFields(logrus.Fields{
		"service": "log-generator",
		"event":   "startup",
	}).Info("Log generator is starting")
	for {
		time.Sleep(time.Second / (time.Duration)(rand.Intn(10)+1))
		message := selectMessage()
		service := selectService()
		event := selectEvent()

		callable := resolveLogMethod(log, event, service, message.loglevel)
		callable(message.message)
	}

}

func resolveLogMethod(log *logrus.Logger, event string, service string, ll Loglevel) func(args ...interface{}) {
	switch ll {
	case Error, Security, Severe:
		return log.WithFields(logrus.Fields{
			"service": service,
			"event":   event,
		}).Error
	case Info:
		return log.WithFields(logrus.Fields{
			"service": service,
			"event":   event,
		}).Info
	case Debug:
		return log.WithFields(logrus.Fields{
			"service": service,
			"event":   event,
		}).Debug
	case Warn:
		return log.WithFields(logrus.Fields{
			"service": service,
			"event":   event,
		}).Warn
	default:
		return log.WithFields(logrus.Fields{
			"service": service,
			"event":   event,
		}).Info
	}
}

func getRandomIndex(length int) int {
	randomIndex := rand.Intn(length)
	return randomIndex
}

func selectEvent() string {
	pick := events[getRandomIndex(len(events))]
	return pick
}

func selectService() string {
	pick := services[getRandomIndex(len(services))]
	return pick
}

func selectMessage() logitem {

	pick := messages[getRandomIndex(len(messages))]
	return pick
}
