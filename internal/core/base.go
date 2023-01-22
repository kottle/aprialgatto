package core

import (
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"github.com/asaskevich/EventBus"
	evbus "github.com/asaskevich/EventBus"
)

var ENABLE_AUTH bool = false
var OBJECT_NEAR string = "obj_near"
var DETECTED_OBJ string = "obj_detected"

var lock = &sync.Mutex{}

type core struct {
	sync.Mutex
	bus     evbus.Bus
	app     *firebase.App
	client  *firestore.Client
	mclient *messaging.Client
	aclient *auth.Client
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

	c.bus = EventBus.New()

	c.loggerInit()
	c.firebaseInit()

}

func (c *core) GetEventBus() evbus.Bus {
	return c.bus
}
