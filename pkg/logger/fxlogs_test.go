package logger

import (
	"errors"
	"go.uber.org/fx/fxevent"
	"os"
	"testing"
	"time"
)

type FxLoggerTestSuite struct {
	logger *FxLogger
}

func (suite *FxLoggerTestSuite) SetupTest() {
	suite.logger = NewFxLogger()
}

func TestFxLogger_LogEvent(t *testing.T) {
	suite := &FxLoggerTestSuite{}
	suite.SetupTest()

	tests := []struct {
		name  string
		event fxevent.Event
	}{
		{
			name:  "OnStartExecuting",
			event: &fxevent.OnStartExecuting{FunctionName: "TestFunction"},
		},
		{
			name:  "OnStartExecutedSuccess",
			event: &fxevent.OnStartExecuted{FunctionName: "TestFunction", Runtime: time.Second},
		},
		{
			name:  "OnStartExecutedError",
			event: &fxevent.OnStartExecuted{FunctionName: "TestFunction", Err: errors.New("start error")},
		},
		{
			name:  "OnStopExecuting",
			event: &fxevent.OnStopExecuting{FunctionName: "TestFunction"},
		},
		{
			name:  "OnStopExecutedSuccess",
			event: &fxevent.OnStopExecuted{FunctionName: "TestFunction", Runtime: time.Second},
		},
		{
			name:  "OnStopExecutedError",
			event: &fxevent.OnStopExecuted{FunctionName: "TestFunction", Err: errors.New("stop error")},
		},
		{
			name:  "SuppliedSuccess",
			event: &fxevent.Supplied{TypeName: "TestType"},
		},
		{
			name:  "SuppliedError",
			event: &fxevent.Supplied{Err: errors.New("supply error")},
		},
		{
			name:  "ProvidedSuccess",
			event: &fxevent.Provided{ConstructorName: "TestConstructor"},
		},
		{
			name:  "ProvidedError",
			event: &fxevent.Provided{Err: errors.New("provide error")},
		},
		{
			name:  "Invoking",
			event: &fxevent.Invoking{FunctionName: "TestFunction"},
		},
		{
			name:  "InvokedSuccess",
			event: &fxevent.Invoked{FunctionName: "TestFunction"},
		},
		{
			name:  "InvokedError",
			event: &fxevent.Invoked{FunctionName: "TestFunction", Err: errors.New("invoke error")},
		},
		{
			name:  "Stopping",
			event: &fxevent.Stopping{Signal: os.Interrupt},
		},
		{
			name:  "StoppedSuccess",
			event: &fxevent.Stopped{},
		},
		{
			name:  "StoppedError",
			event: &fxevent.Stopped{Err: errors.New("stop error")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suite.logger.LogEvent(tt.event)
		})
	}
}
