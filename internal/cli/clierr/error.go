package clierr

import (
	log "github.com/sirupsen/logrus"
)

func HandleDefault(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
