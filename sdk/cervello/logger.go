package cervello

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// internal log for cervello package  ...
// possible levels (info - panic - fatal - warning - error - debug - trace)
func internalLog(level string, msg ...interface{}) {
	if cervelloSdkLogs == "true" {
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
