package motors

import (
	"sync"
	"time"

	"github.com/ev3go/ev3dev"
	log "github.com/sirupsen/logrus"
)

func wStop(motor *ev3dev.TachoMotor, g *sync.WaitGroup) {
	g.Add(1)
	log.Debugf("routine: %s\n", motor.String())
	pos := -10
	for {
		p, _ := motor.Position()
		log.Infof("%s pos: %d %d\n", motor.String(), pos, p)
		if pos == p {
			log.Infof("%s stop!\n", motor.String())
			motor.Command("stop")
			g.Done()
			return
		}
		time.Sleep(time.Second * 1)
		pos = p
	}
}
