package core

import (
	"io"
	"os"
	"sync"

	"github.com/asaskevich/EventBus"
	evbus "github.com/asaskevich/EventBus"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

var OBJECT_NEAR string = "obj_near"
var DETECTED_OBJ string = "obj_detected"

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

	ljack := &lumberjack.Logger{
		Filename:   "/home/robot/logs/aprialgatto.log",
		MaxSize:    2, // megabytes
		MaxBackups: 20,
		MaxAge:     7,    //days
		Compress:   true, // disabled by default
	}
	mWriter := io.MultiWriter(os.Stderr, ljack)
	log.SetOutput(mWriter)

	// Only log the warning severity or above.
	log.SetLevel(log.TraceLevel)

	c.bus = EventBus.New()
}

func (c *core) GetEventBus() evbus.Bus {
	return c.bus
}
