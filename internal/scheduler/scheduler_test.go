package scheduler

import (
	"testing"

	"github.com/EgorKo25/APC/internal/apc"
)

func TestScheduler_InsertTask(t *testing.T) {
	tests := []struct {
		name string
		args []*Task
	}{
		{
			name: "Simple Test",
			args: []*Task{
				{
					TTL: 3,
					AP: apc.AP{
						N1: 1,
						D:  2,
					},
					I: 10,
				},
			},
		},
		{
			name: "Some Tasks",
			args: []*Task{
				{

					TTL: 3,
					AP: apc.AP{
						N1: 1,
						D:  2,
					},
					I: 10,
				},
				{
					TTL: 3,
					AP: apc.AP{
						N1: 1,
						D:  2,
					},
					I: 10,
				},
				{

					TTL: 3,
					AP: apc.AP{
						N1: 1,
						D:  2,
					},
					I: 10,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s := &Scheduler{
				qMaxCount: 5,
			}

			for _, a := range tt.args {
				s.InsertTask(a)
			}

			if len(s.runQ) != len(tt.args) {
				t.Errorf("want length queue: %d, real: %d", len(s.runQ), len(tt.args))
			}

		})
	}
}
