package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EgorKo25/APC/internal/apc"
	"github.com/EgorKo25/APC/internal/scheduler"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetTasksList(t *testing.T) {
	type want struct {
		code        int
		response    []scheduler.Task
		contentType string
	}
	tests := []struct {
		name string
		want want
		s    *scheduler.Scheduler
	}{
		{
			name: "Empty queue test",
			want: want{
				code:        200,
				response:    []scheduler.Task{},
				contentType: "application/json",
			},
			s: scheduler.NewScheduler(6),
		},
		{
			name: "Single task test",
			want: want{
				code: 200,
				response: []scheduler.Task{
					{
						AP:     apc.AP{N1: 1, D: 1},
						I:      10,
						Iter:   100,
						TTL:    20,
						Status: "Wait",
					},
				},
				contentType: "application/json",
			},
			s: scheduler.NewScheduler(6),
		},
		{
			name: "3 tasks test",
			want: want{
				code: 200,
				response: []scheduler.Task{
					{
						AP:     apc.AP{N1: 1, D: 1},
						I:      10,
						Iter:   100,
						TTL:    20,
						Status: "Wait",
					},
					{
						AP:     apc.AP{N1: 1, D: 1},
						I:      10,
						Iter:   100,
						TTL:    20,
						Status: "Wait",
					},
					{
						AP:     apc.AP{N1: 1, D: 1},
						I:      10,
						Iter:   100,
						TTL:    20,
						Status: "Wait",
					},
				},
				contentType: "application/json",
			},
			s: scheduler.NewScheduler(6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var task []scheduler.Task

			var h = NewHandler(tt.s)

			for _, v := range tt.want.response {
				h.s.InsertTask(&v)
			}

			r := httptest.NewRequest(http.MethodGet, "/get/", nil)

			w := httptest.NewRecorder()
			h.GetTasksList(w, r)

			req := w.Result()

			body, err := io.ReadAll(req.Body)
			if err != nil {
				t.Fatalf("error: \"%s\"", err)
			}

			if len(body) > 0 {
				err = json.Unmarshal(body, &task)
				if err != nil {
					t.Fatalf("error: \"%s\"\nbody: \"%s\"", err, body)
				}
			}
			// check status
			assert.Equal(t, tt.want.code, req.StatusCode)
			// check Content-Type
			assert.Equal(t, tt.want.contentType, req.Header.Get("Content-Type"))
			// check answer
			for i, v := range tt.want.response {
				assert.Equal(t, v.TTL, task[i].TTL)
				assert.Equal(t, v.AP, task[i].AP)
				assert.Equal(t, v.I, task[i].I)
				assert.Equal(t, v.Iter, task[i].Iter)
				assert.Equal(t, v.Status, "Wait")
			}

		})
	}
}

func TestHandler_SetTaskToQueue(t *testing.T) {
	type want struct {
		code        int
		response    []scheduler.Task
		contentType string
	}
	tests := []struct {
		name string
		want want
		s    *scheduler.Scheduler
	}{
		{
			name: "Empty queue test",
			want: want{
				code:        200,
				response:    []scheduler.Task{},
				contentType: "text/plain; charset=utf-8",
			},
			s: scheduler.NewScheduler(6),
		},
		{
			name: "1 task from queue test",
			want: want{
				code: 200,
				response: []scheduler.Task{
					{
						AP:     apc.AP{N1: 1, D: 1},
						I:      10,
						Iter:   100,
						TTL:    20,
						Status: "Wait",
					},
				},
				contentType: "text/plain; charset=utf-8",
			},
			s: scheduler.NewScheduler(6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
