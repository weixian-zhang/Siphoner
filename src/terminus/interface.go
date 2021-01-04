package terminus

import (
	"time"
)

type Terminus interface {
	New() (Terminus, error)
	AppendLog(log LogEntry) error
}

type LogEntry struct {
	TimeGenerated time.Time
	Namespace string
	PodName string
	Container string
	Message string
}