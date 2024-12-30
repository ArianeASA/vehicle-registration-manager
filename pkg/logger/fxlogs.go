package logger

import (
	"go.uber.org/fx/fxevent"
)

type FxLogger struct {
	logger *Logger
}

func NewFxLogger() *FxLogger {
	log := NewLoggerWithTrace("fx-system").WithSkip(4)
	return &FxLogger{logger: log}
}

func (l *FxLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.logger.Infof("OnStart hook executing: %s", e.FunctionName)
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.logger.Errorf("OnStart hook failed: %s", e.Err, e.FunctionName)
		} else {
			l.logger.Infof("OnStart hook executed %s, took: %v", e.FunctionName, e.Runtime)
		}
	case *fxevent.OnStopExecuting:
		l.logger.Infof("OnStop hook executing: %s", e.FunctionName)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.logger.Errorf("OnStop hook failed: %s", e.Err, e.FunctionName)
		} else {
			l.logger.Infof("OnStop hook executed: %s, took: %v", e.FunctionName, e.Runtime)
		}
	case *fxevent.Supplied:
		if e.Err != nil {
			l.logger.Errorf("Error supplying", e.Err)
		} else {
			l.logger.Infof("Supplied: %s", e.TypeName)
		}
	case *fxevent.Provided:
		if e.Err != nil {
			l.logger.Errorf("Error providing", e.Err)
		} else {
			l.logger.Infof("Provided: %s", e.ConstructorName)
		}
	case *fxevent.Invoking:
		l.logger.Infof("Invoking: %s", e.FunctionName)
	case *fxevent.Invoked:
		if e.Err != nil {
			l.logger.Errorf("Error invoking: %s", e.Err, e.FunctionName)
		} else {
			l.logger.Infof("Invoked: %s", e.FunctionName)
		}
	case *fxevent.Stopping:
		l.logger.Infof("Stopping: %s", e.Signal.String())
	case *fxevent.Stopped:
		if e.Err != nil {
			l.logger.Errorf("Error stopping", e.Err)
		} else {
			l.logger.Infof("Stopped")
		}
	default:
		l.logger.Infof("Event: %T", event)
	}
}
