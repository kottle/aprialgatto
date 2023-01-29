package sensors

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"github.com/ev3go/ev3dev"
)

type Proximity struct {
	sensor    *ev3dev.Sensor
	threshold int
	logger    logrus.Entry
}

func (p *Proximity) Init(t int) {
	p.threshold = t
}

func (p *Proximity) Run(ctx context.Context) {
	p.logger.Debugf("Run proximity routine")
	defer p.logger.Debugf("Ended proximity routine")

	ticker := time.NewTicker(1 * time.Second)
	foundIt := false
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			v, err := p.sensor.Value(0)
			if err != nil {
				p.logger.Errorf("error reading value: %v", err)
				continue
			}
			value, err := strconv.Atoi(v)
			if err != nil {
				p.logger.Errorf("error reading value: %v", err)
				continue
			}
			p.logger.Infof("proximity value: %d pct", value)
			prevFound := foundIt
			if value >= p.threshold {
				foundIt = true
			} else {
				foundIt = false
			}
			p.logger.Infof("%v %v %v %v", prevFound, foundIt, p.threshold, value)
			if prevFound != foundIt && foundIt {
				p.logger.Infof("trigger action")
				//core.GetCore().GetEventBus().Publish(core.CLOSE_GATE)
			}
		}
	}
}
func NewProximity(port string) *Proximity {
	p := &Proximity{}
	p.logger = *logrus.WithFields(log.Fields{
		"sensor": "lego-ev3-ir" + port,
	})
	var err error
	p.sensor, err = ev3dev.SensorFor(fmt.Sprintf("ev3-ports:%s", port), "lego-ev3-ir")
	if err != nil {
		p.logger.Fatalf("failed to find sensor on %s: %v", port, err)
	}
	return p
}
