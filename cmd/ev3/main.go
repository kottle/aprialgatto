package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aprialgatto/internal/core"
	"github.com/aprialgatto/internal/detection"
	"github.com/aprialgatto/internal/motors"
	"github.com/aprialgatto/internal/sensors"
	"github.com/sirupsen/logrus"
)

func init() {
	core.GetCore().Init()
	core.GetCore().GetEventBus().Subscribe(core.DETECTED_OBJ, detectedObject)
	core.GetCore().GetEventBus().Subscribe(core.OPEN_CAMERA, openCamera)
	core.GetCore().GetEventBus().Subscribe(core.CLOSE_CAMERA, closeCamera)
}

var gate *motors.Gate
var flipper *motors.Flipper

func main() {
	logrus.Error(">>>>>>>>>>>>>>>> STARTED <<<<<<<<<<<<<<<<<<<<")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	flipper = motors.NewFlipper("outC")
	defer flipper.Reset()
	//	flipper.Open()
	//	time.Sleep(4 * time.Second)
	//	flipper.Close()
	gate = motors.NewGate("outA", "outB")
	defer gate.Reset()

	service := detection.NewService()
	proximity := sensors.NewProximityColor("in2")

	service.Start()
	gate.Close()

	proximity.Init(2)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go proximity.Run(ctx)
	<-sigs
}
func openCamera() {
	logrus.Infof("open camera")
	defer logrus.Debugf("open camera finished")

	//flipper.Open()
	core.GetCore().SendMessage("detectObject", "detect")

	// core.GetCore().GetEventBus().Publish(core.OBJECT_NEAR)
}

func closeCamera() {
	logrus.Infof("close camera")
	//flipper.Close()
}

func detectedObject(object string) {
	logrus.Infof("Detected objet: %s", object)
	if !gate.IsOpened() && (object == "person" || object == "cat") {
		core.GetCore().SendMessage("detectObject", "detect_ok")

		gate.Open()
		closeCamera()
		go func() {
			ticker := time.NewTicker(5 * time.Minute)
			<-ticker.C
			gate.Close()
		}()
	}
	if gate.IsOpened() && object == "dog" {
		core.GetCore().SendMessage("detectObject", "detect_ok")
		gate.Close()
		closeCamera()
	}

}
