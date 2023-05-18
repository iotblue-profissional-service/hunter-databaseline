package cervello

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&log.JSONFormatter{})
}
