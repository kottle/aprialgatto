package main

import (
	"context"
	"time"

	"github.com/aprialgatto/internal/core"
	"github.com/aprialgatto/internal/motors"
	"github.com/aprialgatto/internal/sensors"
	"github.com/ev3go/ev3"
)

func init() {
	core.GetCore().Init()
	core.GetCore().GetEventBus().Subscribe(core.OBJECT_NEAR, onObjectNear)
}

var gate *motors.Gate

func main() {

	ev3.LCD.Init(true)
	defer ev3.LCD.Close()

	gate = motors.NewGate("outA", "outB")
	proximity := sensors.NewProximityColor("in2")
	proximity.Init(2)

	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(5 * time.Minute)

	go proximity.Run(ctx)

	<-ticker.C
	cancel()
}

func onObjectNear() {
	// gate.Open()
}
