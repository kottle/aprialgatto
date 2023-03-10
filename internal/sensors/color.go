package sensors

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aprialgatto/internal/core"
	"github.com/aprialgatto/internal/utils"
	"github.com/ev3go/ev3dev"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type ProximityColor struct {
	sensor    *ev3dev.Sensor
	threshold int
	logger    logrus.Entry
}

func (p *ProximityColor) Init(thr int) {
	p.threshold = thr
	p.sensor.SetMode("COL-REFLECT")
}

func NewProximityColor(port string) *ProximityColor {
	p := &ProximityColor{}
	var err error
	p.sensor, err = ev3dev.SensorFor(fmt.Sprintf("ev3-ports:%s", port), "lego-ev3-color")
	p.logger = *logrus.WithFields(log.Fields{
		"sensor": "lego-ev3-color" + port,
	})
	utils.CheckErr(err, fmt.Sprintf("failed to find sensor on %s", port))
	return p
}

func (p *ProximityColor) Run(ctx context.Context) {
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
			p.logger.Tracef("proximity value: %d pct\n", value)
			prevFound := foundIt
			if p.threshold <= value {
				foundIt = true
			} else {
				foundIt = false
			}
			if prevFound != foundIt && foundIt {
				core.GetCore().GetEventBus().Publish(core.OPEN_CAMERA)
				return
			}
		}
	}
}
