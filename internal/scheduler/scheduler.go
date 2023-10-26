package scheduler

import (
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

var (
	statusWait     = "Wait"
	statusRun      = "Run"
	statusFinished = "Finished"
)

// NewScheduler is a constructor for a Scheduler type
func NewScheduler(maxTask int32) *Scheduler {

	return &Scheduler{
		qMaxCount: maxTask,
	}
}

// Scheduler is a custom scheduler for managing working pool of tasks
type Scheduler struct {
	qMaxCount int32   // total data in queue
	runQ      []*Task // Queue of tasks
	qCount    int32

	lock sync.Mutex
}

// InsertTask ...
func (s *Scheduler) InsertTask(t *Task) {

	t.Create = time.Now()
	t.Status = statusWait

	s.lock.Lock()
	s.runQ = append(s.runQ, t)
	s.lock.Unlock()

}

func (s *Scheduler) GetSortQueue() []*Task {

	tmp := s.runQ

	sort.Slice(tmp, func(i, j int) bool {
		if s.runQ[i].Status == statusRun {
			return s.runQ[i].Status > s.runQ[j].Status
		}
		if s.runQ[j].Status == statusRun {
			return s.runQ[i].Status < s.runQ[j].Status
		}
		if s.runQ[i].Status == statusFinished {
			return s.runQ[i].Status < s.runQ[j].Status
		}
		if s.runQ[j].Status == statusFinished {
			return s.runQ[i].Status > s.runQ[j].Status
		}
		return s.runQ[i].Status == s.runQ[j].Status
	})
	return tmp
}

func (s *Scheduler) Run() {
	c := make(chan struct{})

	for {
		s.lock.Lock()
		for i, v := range s.runQ {
			if v.Status == statusFinished && time.Since(v.Finish) >= v.TTL*time.Second {

				s.runQ = s.runQ[1:]

			}
			if s.qCount < s.qMaxCount && v.Status == statusWait {

				s.runQ[i].Status = statusRun

				atomic.AddInt32(&s.qCount, 1)

				go s.runQ[i].Do(c)

			}

		}
		s.lock.Unlock()

		select {
		case <-c:
			atomic.AddInt32(&s.qCount, -1)
			continue
		default:
		}

	}
}
