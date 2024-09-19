package handlers

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/kxrxh/queue-master/api/pb"
)

type TaskQueueHandler struct {
	pb.UnimplementedTaskQueueServer
}

func NewTaskQueueHandler() *TaskQueueHandler {
	return &TaskQueueHandler{}
}

func (h *TaskQueueHandler) SubmitTask(ctx context.Context, req *pb.SubmitTaskRequest) (*pb.SubmitTaskResponse, error) {
	return &pb.SubmitTaskResponse{}, nil
}

func (h *TaskQueueHandler) GetTaskStatus(ctx context.Context, req *pb.GetTaskStatusRequest) (*pb.GetTaskStatusResponse, error) {
	return &pb.GetTaskStatusResponse{}, nil
}

func (h *TaskQueueHandler) StreamTasksResults(grpc.BidiStreamingServer[pb.StreamTasksResultsRequest, pb.StreamTaskResultResponse]) error {
	return nil
}
