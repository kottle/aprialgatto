package core

import (
	"os"
	"sync"

	"github.com/asaskevich/EventBus"
	evbus "github.com/asaskevich/EventBus"
	log "github.com/sirupsen/logrus"
)

var OBJECT_NEAR string = "obj_near"

var lock = &sync.Mutex{}

type core struct {
	bus evbus.Bus
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
	// open a file
	f, err := os.OpenFile("/home/robot/aprialgatto.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(f)

	// Only log the warning severity or above.
	log.SetLevel(log.TraceLevel)

	c.bus = EventBus.New()
}

func (c *core) GetEventBus() evbus.Bus {
	return c.bus
}
