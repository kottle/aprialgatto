package sensors

import (
	"context"
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/ev3go/ev3dev"
)

type Proximity struct {
	sensor    *ev3dev.Sensor
	threshold int
}

func (p *Proximity) Init(t int) {
	p.threshold = t
}

func (p *Proximity) Run(ctx context.Context, found chan bool) {
	ticker := time.NewTicker(1 * time.Second)
	foundIt := false
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			v, err := p.sensor.Value(0)
			if err != nil {
				log.Errorf("error reading value: %v\n", err)
				continue
			}
			value, err := strconv.Atoi(v)
			if err != nil {
				log.Errorf("error reading value: %v\n", err)
				continue
			}
			log.Warnf("proximity value: %d pct\n", value)
			prevFound := foundIt
			if value <= p.threshold {
				foundIt = true
			} else {
				foundIt = false
			}
			if prevFound != foundIt && foundIt {
				log.Tracef("send\n")
				found <- true
				log.Tracef("sent\n")
			}
		}
	}
}

func NewProximity(port string) *Proximity {
	p := &Proximity{}
	var err error
	p.sensor, err = ev3dev.SensorFor(fmt.Sprintf("ev3-ports:%s", port), "lego-ev3-ir")
	if err != nil {
		log.Fatalf("failed to find sensor on %s: %v", port, err)
	}
	return p
}
