package taskclient

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/kxrxh/queue-master/api/pb"
)

type Client struct {
	pb.TaskQueueClient
}

// NewClient creates a new Client from a gRPC connection.
//
// The Client allows you to interact with the task queue service.
func NewClient(conn *grpc.ClientConn) *Client {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	defer logger.Sugar().Infof("Connected to task queue service on %s", conn.Target())

	return &Client{pb.NewTaskQueueClient(conn)}
}

// SubmitTask submits a task to the task queue.
//
// The task is identified by the combination of taskId and taskType. The taskPayload
// is the payload of the task that will be passed to the task handler.
//
// Returns a SubmitTaskResponse with the task ID and status of the task.
func (c *Client) SubmitTask(taskId int, taskType string, taskPayload string) (*pb.SubmitTaskResponse, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Sugar().Infof("Submitting task %d of type %s with payload %s", taskId, taskType, taskPayload)
	taskRequst := pb.SubmitTaskRequest{TaskId: int32(taskId), TaskType: taskType, TaskPayload: taskPayload}
	return c.TaskQueueClient.SubmitTask(context.Background(), &taskRequst)
}

// GetTaskStatus returns the status of a task.
//
// Returns a GetTaskStatusResponse with the task ID and status of the task.
func (c *Client) GetTaskStatus(taskId int) (*pb.GetTaskStatusResponse, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Sugar().Infof("Getting status of task %d", taskId)
	taskRequst := pb.GetTaskStatusRequest{TaskId: int32(taskId)}
	return c.TaskQueueClient.GetTaskStatus(context.Background(), &taskRequst)
}

// StreamTasksResults streams the results of a task.
//
// Returns a StreamTaskResultResponse with the task ID and result of the task.
// func (c *Client) StreamTasksResults(taskId int) (*pb.StreamTaskResultResponse, error) {
// 	logger, _ := zap.NewProduction()
// 	defer logger.Sync()

// 	logger.Sugar().Infof("Streaming results of task %d", taskId)
// 	taskRequst := pb.StreamTasksResultsRequest{TaskId: int32(taskId)}
// 	return c.TaskQueueClient.StreamTasksResults(context.Background())
// }
