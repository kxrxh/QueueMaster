package taskclient

import (
	"context"
	"fmt"
	"io"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/kxrxh/queue-master/api/pb"
)

type Client struct {
	pb.TaskQueueClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	defer logger.Sugar().Infof("Connected to task queue service on %s", conn.Target())

	return &Client{pb.NewTaskQueueClient(conn)}
}

func (c *Client) SubmitTask(taskUuid string, taskType string, taskPayload string) (*pb.SubmitTaskResponse, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Sugar().Infof("Submitting task %d of type %s with payload %s", taskUuid, taskType, taskPayload)
	taskRequst := pb.SubmitTaskRequest{TaskUuid: taskUuid, TaskType: taskType, TaskPayload: taskPayload}
	return c.TaskQueueClient.SubmitTask(context.Background(), &taskRequst)
}

func (c *Client) GetTaskStatus(taskUuid string) (*pb.GetTaskStatusResponse, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Sugar().Infof("Getting status of task %d", taskUuid)
	taskRequst := pb.GetTaskStatusRequest{TaskUuid: taskUuid}
	return c.TaskQueueClient.GetTaskStatus(context.Background(), &taskRequst)
}

func (c *Client) StreamTasksResults(taskUuid string) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Sugar().Infof("Streaming results of task %d", taskUuid)

	stream, err := c.TaskQueueClient.StreamTasksResults(context.Background())
	if err != nil {
		return fmt.Errorf("error creating stream: %v", err)
	}

	// Send the initial request
	err = stream.Send(&pb.StreamTasksResultsRequest{TaskUuid: taskUuid})
	if err != nil {
		return fmt.Errorf("error sending initial request: %v", err)
	}

	// Start a goroutine to receive responses
	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				// End of stream
				logger.Sugar().Info("Stream closed by server")
				return
			}
			if err != nil {
				logger.Sugar().Errorf("Error receiving from stream: %v", err)
				return
			}

			// Process the response here
			logger.Sugar().Infof("Received result: %v", response)
		}
	}()

	// Close the send direction of the stream
	err = stream.CloseSend()
	if err != nil {
		return fmt.Errorf("error closing send stream: %v", err)
	}

	return nil
}
