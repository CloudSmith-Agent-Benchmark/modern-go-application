package command

import (
	"context"
	"os"
	"strconv"
	"time"

	todov1 "github.com/sagikazarmark/todobackend-go-kit/api/todo/v1"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type listOptions struct {
	client todov1.TodoListServiceClient
}

// NewListCommand creates a new cobra.Command for listing todo items.
func NewListCommand(context Context) *cobra.Command {
	options := listOptions{}

	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Short:   "List todo items",
		RunE: func(cmd *cobra.Command, args []string) error {
			options.client = context.GetTodoClient()

			cmd.SilenceErrors = true
			cmd.SilenceUsage = true

			return runList(options)
		},
	}

	return cmd
}

func runList(options listOptions) error {
	req := &todov1.ListItemsRequest{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := options.client.ListItems(ctx, req)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Completed", "Order"})

	for _, item := range resp.GetItems() {
		table.Append([]string{
			item.GetId(),
			item.GetTitle(),
			strconv.FormatBool(item.GetCompleted()),
			strconv.FormatInt(item.GetOrder(), 10),
		})
	}

	table.Render()

	return nil
}
