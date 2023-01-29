package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/aprialgatto/internal/core"
	"github.com/aprialgatto/internal/detection"
	"github.com/aprialgatto/internal/fsm"
	"github.com/aprialgatto/internal/motors"
	"github.com/aprialgatto/internal/sensors"
	"github.com/sirupsen/logrus"
)

var ifsm fsm.FiniteStateMachine

func init() {
	core.GetCore().Init()
	core.GetCore().GetEventBus().Subscribe(core.DETECTED_OBJ, detectedObject)
	core.GetCore().GetEventBus().Subscribe(core.OPEN_CAMERA, openCamera)
	core.GetCore().GetEventBus().Subscribe(core.CLOSE_CAMERA, closeCamera)
	core.GetCore().GetEventBus().Subscribe(core.CLOSE_GATE, closeGate)

	ifsm = fsm.FiniteStateMachine{
		Initial: "wait",
		StateMap: fsm.StateMap{
			"wait": fsm.Transition{
				"proximity_obj": fsm.Action{
					Destination: "open_camera",
				},
			},
			"open_camera": fsm.Transition{
				"detected": fsm.Action{
					Destination: "open_gate",
				},
				"undetected": fsm.Action{
					Destination: "wait",
				},
			},
			"open_gate": fsm.Transition{
				"cat_gone": fsm.Action{
					Destination: "wait",
				},
			},
		},
		CallbackMap: fsm.CallbackMap{
			"wait": func(e fsm.Event, s fsm.State) {
				logrus.Infof("STATE[%s->%s] with event %s ", s, "wait", e)

				gate.Close()

				if colorproximity_cancel != nil {
					colorproximity_cancel()
				}

				var ctx context.Context
				ctx, colorproximity_cancel = context.WithCancel(context.Background())
				go colorproximity.Run(ctx)
			},
			"open_camera": func(e fsm.Event, s fsm.State) {
				logrus.Infof("STATE[%s->%s] with event %s ", s, "open_camera", e)

				if colorproximity_cancel != nil {
					colorproximity_cancel()
				}
				colorproximity_cancel = nil

				//flipper.Open()
				core.GetCore().SendMessage("detectObject", "detect")
			},
			"open_gate": func(e fsm.Event, s fsm.State) {
				gate.Open()

				if irledproximity_cancel != nil {
					irledproximity_cancel()
				}
				irledproximity_cancel = nil

				var ctx context.Context
				ctx, irledproximity_cancel = context.WithCancel(context.Background())
				go irledproximity.Run(ctx)
			},
		},
	}

}

var gate *motors.Gate
var flipper *motors.Flipper
var colorproximity *sensors.ProximityColor
var colorproximity_cancel context.CancelFunc

var irledproximity *sensors.Proximity
var irledproximity_cancel context.CancelFunc

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
	colorproximity = sensors.NewProximityColor("in2")
	irledproximity = sensors.NewProximity("in1") //controllo che il gatto stia mangiando prima di chiudere

	service.Start()
	gate.Close()

	colorproximity.Init(2)
	irledproximity.Init(8)

	//go irledproximity.Run(ctx)

	ifsm.Start()
	<-sigs
}

func openCamera() {
	logrus.Infof("open camera")
	defer logrus.Debugf("open camera finished")

	if irledproximity_cancel != nil {
		irledproximity_cancel()
	}

	//flipper.Open()
	core.GetCore().SendMessage("detectObject", "detect")

	// core.GetCore().GetEventBus().Publish(core.OBJECT_NEAR)
}

func detectedObject(object string) {
	logrus.Infof("Detected objet: %s", object)
	if !gate.IsOpened() && (object == "person" || object == "cat") {
		core.GetCore().SendMessage("detectObject", "detect_ok")

		gate.Open()

		closeCamera()

		//t := time.NewTimer(30 * time.Second)
		//<-t.C

		go irledproximity.Run(context.Background())
	}
	if gate.IsOpened() && object == "dog" {
		core.GetCore().SendMessage("detectObject", "detect_ok")
		gate.Close()
		closeCamera()
	}

}

func closeCamera() {
	logrus.Infof("close camera")
	//flipper.Close()
}

func closeGate() {
	logrus.Infof("close gate")
	gate.Close()
}
