package log

import (
	"log/slog"

	"go.uber.org/fx/fxevent"
)

// FxLogger wraps slog.Logger to implement fxevent.Logger
type FxLogger struct {
	*slog.Logger
}

func NewFxLogger(l *slog.Logger) fxevent.Logger {

	return &FxLogger{Logger: l}
}

// LogEvent implements fxevent.Logger
func (l *FxLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Debug("OnStart hook executing",
			slog.String("callee", e.FunctionName),
			slog.String("caller", e.CallerName),
		)
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Error("OnStart hook failed",
				slog.String("callee", e.FunctionName),
				slog.String("caller", e.CallerName),
				slog.String("err", e.Err.Error()),
			)
		} else {
			l.Debug("OnStart hook executed",
				slog.String("callee", e.FunctionName),
				slog.String("caller", e.CallerName),
				slog.Int64("runtime_ns", e.Runtime.Nanoseconds()),
			)
		}
	case *fxevent.OnStopExecuting:
		l.Debug("OnStop hook executing",
			slog.String("callee", e.FunctionName),
			slog.String("caller", e.CallerName),
		)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Error("OnStop hook failed",
				slog.String("callee", e.FunctionName),
				slog.String("caller", e.CallerName),
				slog.String("err", e.Err.Error()),
			)
		} else {
			l.Debug("OnStop hook executed",
				slog.String("callee", e.FunctionName),
				slog.String("caller", e.CallerName),
				slog.Int64("runtime_ns", e.Runtime.Nanoseconds()),
			)
		}
	case *fxevent.Supplied:
		l.Debug("supplied",
			slog.String("type", e.TypeName),
			slog.String("err", func() string {
				if e.Err != nil {
					return e.Err.Error()
				}
				return ""
			}()),
		)
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.Debug("provided",
				slog.String("constructor", e.ConstructorName),
				slog.String("type", rtype),
			)
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.Debug("decorated",
				slog.String("decorator", e.DecoratorName),
				slog.String("type", rtype),
			)
		}
	case *fxevent.Invoking:
		l.Debug("invoking", slog.String("function", e.FunctionName))
	case *fxevent.Started:
		if e.Err == nil {
			l.Debug("fx started")
		} else {
			l.Error("fx start failed", slog.String("err", e.Err.Error()))
		}
	case *fxevent.LoggerInitialized:
		if e.Err == nil {
			l.Debug("custom fxevent.Logger initialized",
				slog.String("constructor", e.ConstructorName),
			)
		} else {
			l.Error("fx logger initialization failed", slog.String("err", e.Err.Error()))
		}
	}
}
