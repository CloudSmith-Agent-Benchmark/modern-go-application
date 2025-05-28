package commonadapter

import (
	"context"
	"testing"

	"logur.dev/logur"
	"logur.dev/logur/conformance"
	"logur.dev/logur/logtesting"
)

func TestLogger(t *testing.T) {
	t.Parallel()
	t.Run("WithFields", func(t *testing.T) {
		t.Parallel()
		testLoggerWithFields(t)
	})
	t.Run("WithContext", func(t *testing.T) {
		t.Parallel()
		testLoggerWithContext(t)
	})

	suite := conformance.TestSuite{
		LoggerFactory: func(level logur.Level) (logur.Logger, conformance.TestLogger) {
			testLogger := &logur.TestLoggerFacade{}

			return NewLogger(testLogger), testLogger
		},
	}
	t.Run("Conformance", func(t *testing.T) {
		t.Parallel()
		suite.Run(t)
	})
}

func TestContextAwareLogger(t *testing.T) {
	t.Parallel()
	t.Run("WithContext", func(t *testing.T) {
		t.Parallel()
		testContextAwareLoggerWithContext(t)
	})

	suite := conformance.TestSuite{
		LoggerFactory: func(level logur.Level) (logur.Logger, conformance.TestLogger) {
			testLogger := &logur.TestLoggerFacade{}

			return NewContextAwareLogger(
				testLogger,
				func(ctx context.Context) map[string]interface{} {
					return nil
				},
			), testLogger
		},
	}
	t.Run("Conformance", func(t *testing.T) {
		t.Parallel()
		suite.Run(t)
	})
}

func testLoggerWithFields(t *testing.T) {
	testLogger := &logur.TestLoggerFacade{}

	fields := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	logger := NewLogger(testLogger).WithFields(fields)

	logger.Debug("message", nil)

	event := logur.LogEvent{
		Level:  logur.Debug,
		Line:   "message",
		Fields: fields,
	}

	logtesting.AssertLogEventsEqual(t, event, *(testLogger.LastEvent()))
}

func testLoggerWithContext(t *testing.T) {
	testLogger := &logur.TestLoggerFacade{}

	logger := NewLogger(testLogger).WithContext(context.Background())

	logger.Debug("message", nil)

	event := logur.LogEvent{
		Level: logur.Debug,
		Line:  "message",
	}

	logtesting.AssertLogEventsEqual(t, event, *(testLogger.LastEvent()))
}

func testContextAwareLoggerWithContext(t *testing.T) {
	testLogger := &logur.TestLoggerFacade{}

	logger := NewContextAwareLogger(
		testLogger,
		func(_ context.Context) map[string]interface{} {
			return map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			}
		},
	).WithContext(context.Background())

	logger.Debug("message", nil)

	event := logur.LogEvent{
		Level: logur.Debug,
		Line:  "message",
		Fields: map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	logtesting.AssertLogEventsEqual(t, event, *(testLogger.LastEvent()))
}
