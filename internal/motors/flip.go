package motors

import (
	"fmt"
	"time"

	"github.com/ev3go/ev3dev"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var velocity int = 100

type Flipper struct {
	motor  *ev3dev.TachoMotor
	isOpen bool
	logger logrus.Entry
}

func NewFlipper(out1 string) *Flipper {
	f := &Flipper{}
	f.logger = *logrus.WithFields(log.Fields{
		"gate": "lego-ev3-l-motor",
	})
	var err error
	f.motor, err = ev3dev.TachoMotorFor(fmt.Sprintf("ev3-ports:%s", out1), "lego-ev3-m-motor")
	if err != nil {
		log.Fatalf("failed to find medium motor on %s: %v", out1, err)
	}
	f.init()
	f.isOpen = false
	return f
}

func (f *Flipper) init() {
	f.motor.Command("reset").SetStopAction("hold")
}

func (f *Flipper) Stop() {
	f.logger.Debugf("Stop\n")
	f.motor.SetStopAction("hold").Command("stop")
}

func (f *Flipper) Open() {
	if f.isOpen {
		return
	}
	f.isOpen = true
	f.logger.Debugf("Open: %s", f.motor.String())
	defer log.Debugf("Open: %s completed", f.motor.String())
	f.motor.SetSpeedSetpoint(-velocity).SetTimeSetpoint(3 * time.Second).Command("run-timed")
	t := time.NewTicker(2 * time.Second)
	<-t.C
	f.Stop()
	f.logger.Debugf("wait: %s\n", f.motor.String())
}

func (f *Flipper) Close() {
	if !f.isOpen {
		return
	}
	f.logger.Debugf("Close\n")
	f.motor.SetSpeedSetpoint(velocity).SetTimeSetpoint(3 * time.Second).Command("run-timed")
	t := time.NewTicker(5 * time.Second)
	<-t.C
	f.Stop()
	f.isOpen = false
	f.Reset()
}

func (f *Flipper) IsOpened() bool {
	return f.isOpen
}

func (f *Flipper) Reset() {
	f.logger.Debugf("Reset")
	f.motor.Command("reset")
}
