package sensors

import (
	"fmt"

	"github.com/ev3go/ev3dev"
	log "github.com/sirupsen/logrus"
)

type Color struct {
	sensor    *ev3dev.Sensor
	threshold int
}

func (p *Color) Init(r int, g int, b int) {
	//p.threshold = t
}

func NewColor(port string) *Color {
	p := &Color{}
	var err error
	p.sensor, err = ev3dev.SensorFor(fmt.Sprintf("ev3-ports:%s", port), "lego-ev3-color")
	if err != nil {
		log.Fatalf("failed to find sensor on %s: %v", port, err)
	}
	return p
}
