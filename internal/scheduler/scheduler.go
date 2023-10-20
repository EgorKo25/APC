package scheduler

import "sync"

// NewScheduler is a constructor for a Scheduler type
func NewScheduler(maxTask uint) *Scheduler {
	return &Scheduler{
		qcount: maxTask,
	}
}

// Scheduler is a custom scheduler for managing working pool of tasks
type Scheduler struct {
	qcount uint     // total data in queue
	RunQ   []func() // Queue of tasks
	lock   sync.Mutex
}
