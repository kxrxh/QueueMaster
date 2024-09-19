package main

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/kxrxh/queue-master/api/pb"
	"github.com/kxrxh/queue-master/internal/handlers"
	"github.com/kxrxh/queue-master/internal/metrics"
	"github.com/kxrxh/queue-master/internal/taskserver"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	metrics.SetupMetrics()

	server, err := taskserver.NewServer(8080)
	if err != nil {
		logger.Fatal(err.Error())
	}

	taskHandler := handlers.NewTaskQueueHandler()

	server.RegisterService(func(s *grpc.Server) {
		pb.RegisterTaskQueueServer(s, taskHandler)
	})

	defer server.Stop()

	if err := server.Start(); err != nil {
		logger.Fatal(err.Error())
	}
}
