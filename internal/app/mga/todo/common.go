package todo

import (
	"github.com/sagikazarmark/modern-go-application/internal/common"
	"go.opencensus.io/stats"
)

// These interfaces are aliased so that the module code is separated from the rest of the application.
// If the module is moved out of the app, copy the aliased interfaces here.

// Logger is the fundamental interface for all log operations.
type Logger = common.Logger

// ErrorHandler handles an error.
type ErrorHandler = common.ErrorHandler

//nolint: gochecknoglobals
// CreatedTodoItemCount counts the number of todo items created.
var CreatedTodoItemCount = stats.Int64("todo/created_count", "Number of todo items created", stats.UnitDimensionless)

//nolint: gochecknoglobals
// CompletedTodoItemCount counts the number of todo items marked as complete.
var CompletedTodoItemCount = stats.Int64("todo/completed_count", "Number of todo items marked as complete", stats.UnitDimensionless)
