package logging

import (
	"context"
	"fmt"
	"log"
	"time"
)

type LogWriter struct {
}

var Logger = &LogWriter{}

func init() {
	log.SetFlags(0)
	log.SetOutput(Logger)
}

func (l *LogWriter) Log(ctx context.Context, msg ...interface{}) {
	log.Println(fmt.Sprintf("[CorrelationID=%v]", ctx.Value(CorrelationIdContextKey)), msg)
}

func (l *LogWriter) Write(bytes []byte) (int, error) {
	return fmt.Println(time.Now().Format(time.RFC3339Nano) + " " + string(bytes))
}
