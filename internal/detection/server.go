package server

import (
	"fmt"

	"github.com/aprialgatto/internal/detection/grpc"
	"google.golang.org/appengine/log"
)

type Service struct {
	grpc.UnimplementedDetectionServiceServer

	running   bool
	ipaddress string
}

func NewService() *Service {
	log.Infof("Start NewService")
	defer log.Infof("End NewService")

	var host string
	var port int
	var err error

	host, err = config.GetStringVal(cfg.GRPCServer)
	if err != nil {
		host = "127.0.0.1"
	}

	port, err = config.GetIntVal(cfg.GRPCPort)
	if err != nil {
		port = 50054
	}

	srv := &Service{}
	srv.running = false
	srv.ipaddress = fmt.Sprintf("%s:%d", host, port)
	srv.bChannel = &utils.BroadcastChannel{}
	srv.bChannel.Init(nil)
	events.Sub(domotics.UpdateDomo, srv.UpdateDomo)
	events.Sub(domotics.UpdateLights, srv.UpdateState)
	events.Sub(domotics.UpdateClima, srv.UpdateState)
	events.Sub(lgwebos.StatusTV, srv.UpdateState)

	srv.init()
	return srv
}
