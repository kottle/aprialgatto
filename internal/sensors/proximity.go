package sensors

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aprialgatto/internal/core"
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
	ticker := time.NewTicker(1 * time.Second)
	foundIt := false
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			v, err := p.sensor.Value(0)
			if err != nil {
				p.logger.Errorf("error reading value: %v\n", err)
				continue
			}
			value, err := strconv.Atoi(v)
			if err != nil {
				p.logger.Errorf("error reading value: %v\n", err)
				continue
			}
			p.logger.Infof("proximity value: %d pct\n", value)
			prevFound := foundIt
			if p.threshold <= value {
				foundIt = true
			} else {
				foundIt = false
			}
			if prevFound != foundIt && foundIt {
				core.GetCore().GetEventBus().Publish(core.OPEN_CAMERA)
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
