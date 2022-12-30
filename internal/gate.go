package internal

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/ev3go/ev3dev"
)

type Gate struct {
	motors []*ev3dev.TachoMotor
}

func NewGate(out1 string, out2 string) *Gate {
	g := &Gate{}
	g.motors = []*ev3dev.TachoMotor{}
	motor1, err := ev3dev.TachoMotorFor(fmt.Sprintf("ev3-ports:%s", out1), "lego-ev3-l-motor")
	if err != nil {
		log.Fatalf("failed to find medium motor on %s: %v", out1, err)
	}
	g.motors = append(g.motors, motor1)

	motor2, err := ev3dev.TachoMotorFor(fmt.Sprintf("ev3-ports:%s", out2), "lego-ev3-l-motor")
	if err != nil {
		log.Fatalf("failed to find medium motor on %s: %v", out2, err)
	}
	g.motors = append(g.motors, motor2)
	g.init()
	return g
}

func (g *Gate) init() {
}

func (g *Gate) Open() {
	log.Debugf("Open\n")
	g.exec(func(m *ev3dev.TachoMotor) {
		m.Command("reset").SetSpeedSetpoint(300).SetPolarity(ev3dev.Normal).SetPositionSetpoint(int(-90)).Command("run-to-rel-pos")
	})
}

func (g *Gate) Close() {
	log.Debugf("Close\n")

	g.exec(func(m *ev3dev.TachoMotor) {
		m.Command("reset").SetSpeedSetpoint(300).SetPolarity(ev3dev.Normal).SetPositionSetpoint(int(90)).Command("run-to-rel-pos")
	})
}

func (g *Gate) exec(callback func(m *ev3dev.TachoMotor)) {
	for _, m := range g.motors {
		callback(m)
	}
}
