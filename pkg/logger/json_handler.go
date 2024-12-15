package logger

import (
	"encoding/json"
	"fmt"
	"io"
)

type Handler interface {
	Handle(*LogRecord) error
}

type JSONHandler struct {
	Writer io.Writer
}

func newJSONHandler(writer io.Writer) *JSONHandler {
	return &JSONHandler{Writer: writer}
}

func (h *JSONHandler) Handle(record *LogRecord) error {
	jsonMsg, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("failed to marshal log record: %w", err)
	}
	_, err = h.Writer.Write([]byte(fmt.Sprintf("%s\n", jsonMsg)))
	return fmt.Errorf("failed to write log message: %w", err)
}

func (l *Logger) SetHandler(handler Handler) {
	l.handler = handler
}
