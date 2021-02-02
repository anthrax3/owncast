package admin

import (
	"net/http"

	"github.com/owncast/owncast/core/data"
	log "github.com/sirupsen/logrus"
)

func ResetYPRegistration(w http.ResponseWriter, r *http.Request) {
	log.Traceln("Resetting YP registration key")
	if err := data.SetDirectoryRegistrationKey(""); err != nil {
		log.Errorln(err)
	}
}
