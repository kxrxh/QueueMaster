package cli

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/kxrxh/queue-master/internal/taskclient"
)

var (
	taskType    string
	taskPayload string
)

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit a task to the server",
	Args:  cobra.ExactArgs(0),
	Run:   runSubmit,
}

// runSubmit implements the logic for the "submit" subcommand.
// It creates a new UUID for the task, creates a new client, and submits the task using the client.
// It will print an error message if the submission fails, or a success message
// including the task ID if the submission succeeds.
func runSubmit(cmd *cobra.Command, args []string) {
	// Create task UUID
	uuid := uuid.NewString()

	client, err := getClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	submit(client, uuid, taskType, taskPayload)
}

// submit submits a task to the task queue server using the given client.
// It will print an error message if the submission fails, or a success message
// including the task ID if the submission succeeds.
func submit(client *taskclient.Client, uuid, taskType, taskPayload string) {
	response, err := client.SubmitTask(uuid, taskType, taskPayload)
	if err != nil {
		fmt.Printf("Failed to submit task: %v\n", err)
		return
	}

	fmt.Printf("Task submitted successfully. Task ID: %s\n", response.TaskUuid)
}

// getClient creates a new client for the task queue service at the address
// specified by the "address" viper configuration key.
//
// It returns an error if the connection to the task queue service cannot be
// established.
func getClient() (*taskclient.Client, error) {
	address := viper.GetString("address")
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to task queue service: %v", err)
	}
	return taskclient.NewClient(conn), nil
}
