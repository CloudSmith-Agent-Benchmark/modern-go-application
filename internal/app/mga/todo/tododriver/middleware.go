package tododriver

import (
	"context"

	todo2 "github.com/sagikazarmark/modern-go-application/internal/app/mga/todo"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	"github.com/sagikazarmark/todobackend-go-kit/todo"
)

//nolint: gochecknoglobals,lll
// CreatedTodoItemCountView counts the number of todo items created.
var CreatedTodoItemCountView = &view.View{
	Name:        "todo/created_count",
	Description: "Count of todo items created",
	Measure:     todo2.CreatedTodoItemCount,
	Aggregation: view.Count(),
	TagKeys:     []tag.Key{},
}

//nolint: gochecknoglobals
// CompleteTodoItemCountView counts the number of todo items marked as complete.
var CompleteTodoItemCountView = &view.View{
	Name:        "todo/completed_count",
	Description: "Count of todo items marked as complete",
	Measure:     todo2.CompletedTodoItemCount,
	Aggregation: view.Count(),
	TagKeys:     []tag.Key{},
}

// InstrumentationMiddleware is a service middleware that instruments certain methods.
type InstrumentationMiddleware struct {
	next todo.Service
}

// NewInstrumentationMiddleware returns a new instrumentation middleware.
func NewInstrumentationMiddleware() InstrumentationMiddleware {
	return InstrumentationMiddleware{}
}

// Attach attaches the middleware to a service.
func (mw InstrumentationMiddleware) Attach(next todo.Service) todo.Service {
	mw.next = next

	return mw
}

func (mw InstrumentationMiddleware) AddItem(ctx context.Context, newItem todo.NewItem) (todo.Item, error) {
	id, err := mw.next.AddItem(ctx, newItem)
	if err == nil {
		stats.Record(ctx, todo2.CreatedTodoItemCount.M(1))
	}

	return id, err
}

func (mw InstrumentationMiddleware) ListItems(ctx context.Context) ([]todo.Item, error) {
	return mw.next.ListItems(ctx)
}

func (mw InstrumentationMiddleware) DeleteItems(ctx context.Context) error {
	return mw.next.DeleteItems(ctx)
}

func (mw InstrumentationMiddleware) GetItem(ctx context.Context, id string) (todo.Item, error) {
	return mw.next.GetItem(ctx, id)
}

func (mw InstrumentationMiddleware) UpdateItem(ctx context.Context, id string, itemUpdate todo.ItemUpdate) (todo.Item, error) { //nolint: lll
	return mw.next.UpdateItem(ctx, id, itemUpdate)
}

func (mw InstrumentationMiddleware) DeleteItem(ctx context.Context, id string) error {
	return mw.next.DeleteItem(ctx, id)
}
