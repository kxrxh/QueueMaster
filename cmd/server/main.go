package main

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/kxrxh/queue-master/api/pb"
	"github.com/kxrxh/queue-master/internal/handlers"
	"github.com/kxrxh/queue-master/internal/metrics"
	"github.com/kxrxh/queue-master/internal/taskserver"
	"github.com/kxrxh/queue-master/pgk/utils"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	server, err := taskserver.NewServer(8080)
	utils.FailOnError(err, "Failed to create gRPC server")

	taskHandler := handlers.NewTaskQueueHandler()

	server.RegisterService(func(s *grpc.Server) {
		pb.RegisterTaskQueueServer(s, taskHandler)
	})

	// Start metrics server
	metricsServer := metrics.SetupMetrics()

	// Start gRPC server
	go func() {
		err := server.Start()
		utils.FailOnError(err, "Failed to start gRPC server")
	}()

	// Wait for SIGINT or SIGTERM
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Stop gRPC server and metrics server gracefully
	<-stop
	server.Shutdown()
	err = metricsServer.Shutdown()

	utils.FailOnError(err, "Failed to shutdown metrics server")
}
