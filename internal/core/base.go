package core

import (
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
)

var lock = &sync.Mutex{}

type core struct {
}

var coreInstance *core

func GetCore() *core {
	if coreInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if coreInstance == nil {
			coreInstance = &core{}
		}
	}

	return coreInstance
}

func (c *core) Init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)

}
