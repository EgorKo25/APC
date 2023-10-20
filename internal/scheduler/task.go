package scheduler

import (
	"time"

	"github.com/EgorKo25/APC/internal/apc"
)

func newTask(ap *apc.AP, i float64) *task {
	t := &task{
		I:      time.Duration(i) * time.Second,
		Create: time.Now(),
	}

	t.D = ap.D
	t.N1 = ap.N1

	return t
}

type task struct {
	apc.AP
	I         time.Duration `json:"interval,omitempty"` // interval between iter
	Iter      int           `json:"iter"`
	iterCount int           // number of iter

	Status string        `json:"status"`
	TTL    time.Duration `json:"ttl"`

	Create time.Time `json:"create"`
	Start  time.Time `json:"start"`
	Finish time.Time `json:"finish"`
}

func (t *task) Do() bool {
	t.Start = time.Now()
	for i := 0; i < t.iterCount; i++ {
		t.Count()
		t.Iter = i + 1
		time.Sleep(t.I)
	}
	t.Status = statusFinished
	return true
}
