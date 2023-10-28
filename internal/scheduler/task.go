package scheduler

import (
	"time"
)

type AP interface {
	Count()
}

type Task struct {
	AP
	I       time.Duration `json:"interval,omitempty"` // interval between iter
	Iter    int           `json:"iteration,omitempty"`
	IterNow int           `json:"iter_now"` // number of iter

	Status string        `json:"status,omitempty"`
	TTL    time.Duration `json:"ttl,omitempty"` // times to life before finished

	Create time.Time `json:"create,omitempty"`
	Start  time.Time `json:"start,omitempty"`
	Finish time.Time `json:"finish,omitempty"`
}

func (t *Task) Do(c chan struct{}) {

	t.Start = time.Now()
	for i := 0; i < t.Iter; i++ {
		t.Count()
		t.IterNow = i + 1
		time.Sleep(t.I * time.Second)
	}
	t.Finish = time.Now()
	t.Status = statusFinished
	c <- struct{}{}
}
