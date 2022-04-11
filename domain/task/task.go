package task

import (
	"fmt"
	"time"
)

type Task struct {
	Name      string
	Tags      []string
	TimeSlots []TimeSlot
	Finished  bool
}

type TimeSlot struct {
	Start time.Time
	End   *time.Time
}

func NewTimeSlot() *TimeSlot {
	return &TimeSlot{
		Start: time.Now(),
		End:   nil,
	}
}

func NewTask(name string, tags ...string) *Task {
	return &Task{
		Name:      name,
		Tags:      tags,
		TimeSlots: []TimeSlot{*NewTimeSlot()},
		Finished:  false,
	}
}

func (t *Task) Fmt() string {
	paused := false
	var duration time.Duration
	for _, tt := range t.TimeSlots {
		if tt.End != nil {
			duration += tt.End.Sub(tt.Start)
		} else {
			duration += time.Now().Sub(tt.Start)
			paused = true
		}
	}
	status := "paused"
	if !paused {
		status = "not " + status
	}
	if t.Finished {
		status = "finished"
	}
	return fmt.Sprintf("Task %s: is %s with %s tracked time", t.Name, status, duration.String())
}