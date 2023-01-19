package server

import (
	"context"
	"fmt"
	"net"

	"github.com/aprialgatto/internal/core"
	"github.com/aprialgatto/internal/detection/api"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var cssContext api.DetectionService_OnDetectObjectServer

type Service struct {
	api.UnimplementedDetectionServiceServer

	running   bool
	ipaddress string
}

func NewService() *Service {
	log.Infof("Start NewService")
	defer log.Infof("End NewService")

	srv := &Service{}
	srv.running = false
	srv.ipaddress = fmt.Sprintf("%s:%d", "0.0.0.0", 5555)

	core.GetCore().GetEventBus().Subscribe(core.OBJECT_NEAR, onObjectNear)

	return srv
}

func (s *Service) Start() {
	s.running = true
	go func() {
		lis, err := net.Listen("tcp", s.ipaddress)
		if err != nil {
			log.Errorf("error on grp listener: %v", err)
			return
		}
		grpcServer := grpc.NewServer()
		api.RegisterDetectionServiceServer(grpcServer, s)

		grpcServer.Serve(lis)
	}()
}

func (s *Service) DetectedObject(ctx context.Context, req *api.DetectReq) (*api.DetectRes, error) {
	core.GetCore().GetEventBus().Publish(core.DETECTED_OBJ, req.GetObject())
	return &api.DetectRes{}, nil
}

func (s *Service) OnDetectObject(res *api.OnDetectRes, css api.DetectionService_OnDetectObjectServer) error {
	logrus.Infof(">>>>>>>>>>>>>>>OnDetectObject<<<<<<<<<<<<<<<<<<<")

	cssContext = css
	for {
		select {
		case <-css.Context().Done():
			logrus.Infof("on detect object closes")
			return nil
		}
	}
}

func onObjectNear() {
	// gate.Open()
	logrus.Infof("onObjectNear")
	if cssContext != nil {
		cssContext.Send(&api.OnDetectReq{})
	}
}
