package util

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func GetExecDir() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Errorf(err.Error())
		return ""
	}
	return filepath.Dir(exePath)

}
