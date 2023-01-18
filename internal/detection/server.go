package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"

	"github.com/aprialgatto/internal/core"
	"github.com/aprialgatto/internal/detection/api"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
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
		/*tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}*/

		grpcServer := grpc.NewServer(
		//grpc.Creds(tlsCredentials),
		)
		api.RegisterDetectionServiceServer(grpcServer, s)

		grpcServer.Serve(lis)
	}()
}

func (s *Service) DetectedObject(context.Context, *api.DetectReq) (*api.DetectRes, error) {

	return nil, status.Errorf(codes.Unimplemented, "method DetectedObject not implemented")
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

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("/home/robot/cert/ca-cert.pem", "/home/robot/cert/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}
