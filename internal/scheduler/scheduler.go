package scheduler

import (
	"encoding/json"
	"log"
	"os"
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
func NewScheduler(storageInterval, maxTask int32, storeFilePath string) (*Scheduler, error) {

	file, err := os.OpenFile(storeFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return nil, err
	}

	return &Scheduler{
		qMaxCount:       maxTask,
		file:            file,
		storageInterval: time.Duration(storageInterval) * time.Second,
	}, nil
}

// Scheduler is a custom scheduler for managing working pool of tasks
type Scheduler struct {
	qMaxCount int32   // total data in queue
	runQ      []*Task // Queue of tasks
	qCount    int32

	lock sync.Mutex

	// for storaging of tasks
	storageInterval time.Duration
	file            *os.File
}

// InsertTask append task to queue
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

		return s.runQ[i].Status < s.runQ[j].Status
	})
	return tmp
}

// WriteAll write data to file
func (s *Scheduler) WriteAll() (err error) {
	var data []byte

	tmp := s.GetSortQueue()

	if data, err = json.Marshal(tmp); err != nil {
		return err
	}

	if err = s.file.Truncate(0); err != nil {
		return err
	}

	if _, err = s.file.Write(data); err != nil {
		return err
	}

	if _, err = s.file.Write([]byte("\n")); err != nil {
		return err
	}

	return nil
}

func (s *Scheduler) Run() {
	c := make(chan struct{})
	tickerSave := time.NewTicker(s.storageInterval)
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
		case <-tickerSave.C:
			if err := s.WriteAll(); err != nil {
				log.Printf("error: \"%s\"", err)
				return
			}

		case <-c:
			atomic.AddInt32(&s.qCount, -1)
		default:

		}
	}
}
