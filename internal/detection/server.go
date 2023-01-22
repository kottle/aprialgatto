package server

import (
	"context"
	"fmt"
	"net"

	"github.com/aprialgatto/internal/core"
	"github.com/aprialgatto/internal/detection/api"
	"github.com/aprialgatto/internal/utils/events"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func grpcAuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	reg, res := events.Pub(core.VerifyAuthToken, core.NewVerifyAuthToken(token))
	if reg == 1 {
		result := <-res
		if result.Error != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", result.Error)
		}
	}

	return ctx, nil
}

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
		var grpcServer *grpc.Server
		if core.ENABLE_AUTH {
			grpcServer = grpc.NewServer(
				grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(grpcAuthFunc)),
				grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(grpcAuthFunc)),
			)
		} else {
			grpcServer = grpc.NewServer()
		}
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
	core.GetCore().SendMessage("detectObject", "detect")
	/*	// gate.Open()
		logrus.Infof("onObjectNear")
		if cssContext != nil {
			cssContext.Send(&api.OnDetectReq{})
		}
	*/
}
