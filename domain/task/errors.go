package task

import "fmt"

var (
	ErrTaskExists      = fmt.Errorf("task already exists")
	ErrTaskNotExists   = fmt.Errorf("task not exists")
	ErrNoTimeSlots     = fmt.Errorf("task has not time slots")
	ErrSlotFinished    = fmt.Errorf("time slot alerady finished")
	ErrSlotNotFinished = fmt.Errorf("time slot not finished")
	ErrTaskFinished    = fmt.Errorf("task is finished")
)
