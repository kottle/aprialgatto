package main

import (
	"github.com/aprialgatto/internal/core"
	"github.com/aprialgatto/internal/motors"
	"github.com/aprialgatto/internal/sensors"
	"github.com/ev3go/ev3"
)

func init() {
	core.GetCore().Init()
}

func main() {

	ev3.LCD.Init(true)
	defer ev3.LCD.Close()

	testGate()
	testColors()
	/*
		proximity := internal.NewProximity("in1")
		proximity.Init(20)

		ctx, cancel := context.WithCancel(context.Background())
		done := make(chan bool, 1)
		ticker := time.NewTicker(2 * time.Minute)

		go proximity.Run(ctx, done)

		for {
			log.Warnf("Started!\n")
			select {
			case <-ticker.C:
				cancel()
				return
			case <-done:
				log.Debugf("Found!\n")
				gate.Open()
				time.Sleep(5 * time.Second)
				gate.Close()
				time.Sleep(1 * time.Second)
			}
		}*/

}

func testGate() {
	gate := motors.NewGate("outA", "outB")
	gate.Close()
	/*gate.Open()
	time.Sleep(5 * time.Second)
	gate.Close()*/
}

func testColors() {
	color := sensors.NewColor("in1")
	color.Init(20, 29, 0)
}
