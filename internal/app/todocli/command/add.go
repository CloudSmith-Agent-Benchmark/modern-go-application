package command

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"

	todov1 "github.com/sagikazarmark/todobackend-go-kit/api/todo/v1"
)

type createOptions struct {
	title  string
	client todov1.TodoListServiceClient
}

// NewAddCommand creates a new cobra.Command for adding a new item to the list.
func NewAddCommand(context Context) *cobra.Command {
	options := createOptions{}

	cmd := &cobra.Command{
		Use:     "add",
		Aliases: []string{"a"},
		Short:   "Add an item to the list",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			options.title = args[0]
			options.client = context.GetTodoClient()

			cmd.SilenceErrors = true
			cmd.SilenceUsage = true

			return runCreate(options)
		},
	}

	return cmd
}

func runCreate(options createOptions) error {
	req := &todov1.AddItemRequest{
		Title: options.title,
		Order: 0,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := options.client.AddItem(ctx, req)
	if err != nil {
		st := status.Convert(err)
		for _, detail := range st.Details() {
			//nolint: gocritic
			switch t := detail.(type) {
			case *errdetails.BadRequest:
				logger := log.New(os.Stderr, "", 0)
				logger.Println("Oops! Your request was rejected by the server.")
				for _, violation := range t.GetFieldViolations() {
					logger.Printf("The %q field was wrong:\n", violation.GetField())
					logger.Printf("\t%s\n", violation.GetDescription())
				}
			}
		}

		return err
	}

	logger := log.New(os.Stdout, "", 0)
	logger.Printf("Todo item %q with ID %s has been created.", options.title, resp.GetItem().GetId())

	return nil
}
