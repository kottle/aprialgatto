package motors

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"github.com/ev3go/ev3dev"
)

type Gate struct {
	motors []*ev3dev.TachoMotor
	isOpen bool
	logger logrus.Entry
}

func NewGate(out1 string, out2 string) *Gate {
	g := &Gate{}
	g.motors = []*ev3dev.TachoMotor{}
	g.logger = *logrus.WithFields(log.Fields{
		"gate": "lego-ev3-l-motor",
	})
	motor1, err := ev3dev.TachoMotorFor(fmt.Sprintf("ev3-ports:%s", out1), "lego-ev3-l-motor")
	if err != nil {
		g.logger.Fatalf("failed to find medium motor on %s: %v", out1, err)
	}
	g.motors = append(g.motors, motor1)

	motor2, err := ev3dev.TachoMotorFor(fmt.Sprintf("ev3-ports:%s", out2), "lego-ev3-l-motor")
	if err != nil {
		g.logger.Fatalf("failed to find medium motor on %s: %v", out2, err)
	}
	g.motors = append(g.motors, motor2)
	g.init()
	g.isOpen = false
	return g
}

func (g *Gate) init() {
	g.exec(func(m *ev3dev.TachoMotor) {
		m.Command("reset").SetStopAction("hold")
	})
}

func (g *Gate) Stop() {
	g.logger.Debugf("Stop\n")
	g.exec(func(m *ev3dev.TachoMotor) {
		m.SetStopAction("hold").Command("stop")
	})
}

func (g *Gate) Open() {
	if g.isOpen {
		return
	}
	g.isOpen = true
	g.logger.Debugf("Open\n")
	wg := sync.WaitGroup{}
	g.exec(func(m *ev3dev.TachoMotor) {
		g.logger.Debugf("Open: %s\n", m.String())
		m.SetSpeedSetpoint(100).Command("run-forever")
		go wStop(g.logger, m, &wg)
	})
	wg.Wait()
}

func (g *Gate) Close() {
	if !g.isOpen {
		return
	}
	log.Debugf("Close\n")
	wg := sync.WaitGroup{}
	g.exec(func(m *ev3dev.TachoMotor) {
		m.SetSpeedSetpoint(-100).Command("run-forever")
		go wStop(g.logger, m, &wg)
	})
	wg.Wait()
	g.isOpen = false
	g.Reset()
}

func (g *Gate) exec(callback func(m *ev3dev.TachoMotor)) {
	for _, m := range g.motors {
		callback(m)
	}
}

func (g *Gate) IsOpened() bool {
	return g.isOpen
}
func (g *Gate) Reset() {
	g.logger.Debugf("Reset")
	g.exec(func(m *ev3dev.TachoMotor) {
		m.Command("reset")
	})
}
