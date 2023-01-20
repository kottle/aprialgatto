package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aprialgatto/internal/core"
	server "github.com/aprialgatto/internal/detection"
	"github.com/aprialgatto/internal/motors"
	"github.com/aprialgatto/internal/sensors"
	"github.com/sirupsen/logrus"
)

func init() {
	core.GetCore().Init()
	core.GetCore().GetEventBus().Subscribe(core.DETECTED_OBJ, detectedObject)
}

var gate *motors.Gate

func main() {
	logrus.Error(">>>>>>>>>>>>>>>> STARTED <<<<<<<<<<<<<<<<<<<<")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	service := server.NewService()
	service.Start()
	gate = motors.NewGate("outA", "outB")
	gate.Close()
	proximity := sensors.NewProximityColor("in2")
	proximity.Init(2)
	ctx, cancel := context.WithCancel(context.Background())
	go proximity.Run(ctx)

	<-sigs
	cancel()
}

func detectedObject(object string) {
	logrus.Infof("Detected objet: %s", object)
	if !gate.IsOpened() && (object == "person" || object == "cat") {
		gate.Open()
		go func() {
			ticker := time.NewTicker(5 * time.Minute)
			<-ticker.C
			gate.Close()
		}()
	}
	if gate.IsOpened() && object == "dog" {
		gate.Close()
	}
}
