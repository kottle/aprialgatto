package main

import (
	"context"
	"time"

	"github.com/aprialgatto/internal/core"
	server "github.com/aprialgatto/internal/detection"
	"github.com/aprialgatto/internal/motors"
	"github.com/aprialgatto/internal/sensors"
	"github.com/sirupsen/logrus"

	"github.com/ev3go/ev3"
)

func init() {
	core.GetCore().Init()
	core.GetCore().GetEventBus().Subscribe(core.DETECTED_OBJ, detectedObject)
}

var gate *motors.Gate

func main() {

	ev3.LCD.Init(true)
	defer ev3.LCD.Close()
	service := server.NewService()
	service.Start()
	gate = motors.NewGate("outA", "outB")
	proximity := sensors.NewProximityColor("in2")
	proximity.Init(2)

	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(5 * time.Minute)

	go proximity.Run(ctx)

	<-ticker.C
	cancel()
}

func detectedObject(object string) {
	logrus.Infof("Detected objet: %s", object)
	if !gate.IsOpened() && object == "person" {
		gate.Open()
	}
	if gate.IsOpened() && object == "dog" {
		gate.Close()
	}
}
