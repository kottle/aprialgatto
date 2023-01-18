package server

import (
	"context"
	"fmt"
	"net"

	"github.com/aprialgatto/internal/detection/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

func (s *Service) DetectedObject(context.Context, *api.DetectReq) (*api.DetectRes, error) {

	return nil, status.Errorf(codes.Unimplemented, "method DetectedObject not implemented")
}
