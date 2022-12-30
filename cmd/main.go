package main

import (
	"context"
	"os"
	"time"

	"github.com/aprialgatto/internal"
	"github.com/ev3go/ev3"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)

	ev3.LCD.Init(true)
	defer ev3.LCD.Close()

	gate := internal.NewGate("outA", "outB")
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
	}

}
