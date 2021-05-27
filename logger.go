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
	var correlationId string
	if ctx != nil {
		correlationId = ctx.Value(CorrelationIdContextKey).(string)
	}
	log.Println(fmt.Sprintf("[CorrelationID=%s]", correlationId), msg)
}

func (l *LogWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("2006-01-02 15:04:05.000000000") + " " + string(bytes))
}
