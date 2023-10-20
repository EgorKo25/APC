package scheduler

import (
	"sync"
	"time"

	"github.com/EgorKo25/APC/internal/apc"
)

var (
	statusWait     = "Wait"
	statusRun      = "Run"
	statusFinished = "Finished"
)

// NewScheduler is a constructor for a Scheduler type
func NewScheduler(maxTask uint) *Scheduler {
	return &Scheduler{
		qMaxCount: maxTask,
	}
}

// Scheduler is a custom scheduler for managing working pool of tasks
type Scheduler struct {
	qMaxCount uint    // total data in queue
	RunQ      []*task // Queue of tasks
	qCount    uint
	lock      sync.Mutex
}

// InsertTask ...
func (s *Scheduler) InsertTask(ap *apc.AP, i float64) {

	t := newTask(ap, i)

	s.lock.Lock()
	s.RunQ = append(s.RunQ, t)
	s.lock.Unlock()

	t.Status = statusWait
}

func (s *Scheduler) Run() {
	for {
		for i, v := range s.RunQ {
			if v.Status == statusFinished && time.Since(v.Finish) >= v.TTL {
				s.RunQ = s.RunQ[1:]
			}
			if s.qCount < s.qMaxCount && v.Status == statusWait {

				s.RunQ[i].Status = statusRun

				s.lock.Lock()
				s.qCount++
				s.lock.Unlock()

				go func() {
					if s.RunQ[i].Do() {
						s.lock.Lock()
						s.qCount--
						s.lock.Unlock()
					}
				}()
			}
		}
	}
}
