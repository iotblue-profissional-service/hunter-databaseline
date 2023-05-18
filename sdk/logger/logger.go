package logger

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var file *os.File

func init() {
	logrus.SetFormatter(&log.JSONFormatter{})

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModePerm)
	}

	var err error
	file, err = os.OpenFile(getTodaysFileName(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	timer, err := strconv.ParseUint(logFilesTimeInDays, 10, 32)
	if err != nil {
		panic(err)
	}
	gocron.Every(timer).Days().Do(delteOldFile)
	gocron.Every(1).Day().At("12:00").Do(openDayFileTask)

	go gocron.Start()

}

// LogMessage ...
// possible levels (info - panic - fatal - warning - error - debug - trace)
func LogMessage(level string, msg ...interface{}) {

	mw := io.MultiWriter(file)
	logrus.SetOutput(mw)
	if enableLogs == "true" {
		switch level {
		case "info":
			log.Info(fmt.Sprintln(msg...))
			break
		case "panic":
			log.Panic(fmt.Sprintln(msg...))
			break
		case "fatal":
			log.Fatal(fmt.Sprintln(msg...))
			break
		case "warning":
			log.Warn(fmt.Sprintln(msg...))
			break
		case "error":
			log.Error(fmt.Sprintln(msg...))
			break
		case "debug":
			log.Debug(fmt.Sprintln(msg...))
			break
		case "trace":
			log.Trace(fmt.Sprintln(msg...))
			break
		default:
			log.Info(fmt.Sprintln(msg...))
		}
	}
}

func delteOldFile() {
	timer, err := strconv.ParseInt(logFilesTimeInDays, 10, 64)
	if err != nil {
		log.Println(err)
	}
	daysAgoDate := time.Now().Add(time.Duration(timer) * time.Hour)

	if _, err := os.Stat(daysAgoDate.Format("01-02-2006")); err == nil {
		e := os.Remove(daysAgoDate.Format("01-02-2006"))
		if e != nil {
			panic(e)
		}
	}
}

func openDayFileTask() {
	file.Close()
	var err error
	file, err = os.OpenFile(getTodaysFileName(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

}

func getTodaysFileName() string {
	dt := time.Now()
	return fmt.Sprintf("%v/%v", "logs", dt.Format("01-02-2006"))
}
