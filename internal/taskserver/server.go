package taskserver

import (
	"fmt"
	"net"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/kxrxh/queue-master/pgk/utils"
)

// Server wraps the gRPC server and its listener
type Server struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

// NewServer creates a new Server that listens on the given port.
//
// The given port is used to listen for incoming gRPC connections.
// If the port is already in use, an error is returned.
//
// The returned Server is ready to use, and can be started with the Start method.
func NewServer(port int) (*Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	utils.FailOnError(err, "Failed to listen on port "+strconv.Itoa(port))
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	return &Server{
		grpcServer: s,
		listener:   lis,
	}, nil
}

// RegisterService registers a service with the gRPC server
func (s *Server) RegisterService(registerFunc func(*grpc.Server)) {
	registerFunc(s.grpcServer)
}

// Start starts the gRPC server
func (s *Server) Start() error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Sugar().Infof("Starting gRPC server on %s", s.listener.Addr().String())
	return s.grpcServer.Serve(s.listener)
}

// Shutdown stops the gRPC server
func (s *Server) Shutdown() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Sugar().Infof("Stopping gRPC server on %s", s.listener.Addr().String())
	s.grpcServer.GracefulStop()
}
