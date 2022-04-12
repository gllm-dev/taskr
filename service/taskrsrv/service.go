package taskrsrv

import (
	"time"

	"go.gllm.dev/taskr/domain/task"
	"go.gllm.dev/taskr/ports"
)

type service struct {
	repo ports.TaskrRepository
}

func New(repo ports.TaskrRepository) *service {
	return &service{
		repo: repo,
	}
}

func (srv *service) AddTask(name string, tags ...string) error {
	_, err := srv.GetTask(name)
	if err == nil {
		return task.ErrTaskExists
	}

	t := task.NewTask(name, tags...)
	return srv.repo.Create(t)
}

func (srv *service) FinishTask(name string) error {
	t, err := srv.GetTask(name)
	if err != nil {
		return err
	}

	if t.Finished {
		return task.ErrTaskFinished
	}

	slotsLen := len(t.TimeSlots)
	if slotsLen == 0 {
		return task.ErrNoTimeSlots
	}

	if t.TimeSlots[slotsLen-1].End == nil {
		now := time.Now()
		t.TimeSlots[slotsLen-1].End = &now
	}
	t.Finished = true
	return srv.repo.Update(&t)
}

func (srv *service) ResumeTask(name string) error {
	t, err := srv.GetTask(name)
	if err != nil {
		return err
	}

	if t.Finished {
		return task.ErrTaskFinished
	}

	slotsLen := len(t.TimeSlots)
	if slotsLen == 0 {
		return task.ErrNoTimeSlots
	}

	if t.TimeSlots[slotsLen-1].End == nil {
		return task.ErrSlotNotFinished
	}

	t.TimeSlots = append(t.TimeSlots, *task.NewTimeSlot())
	return srv.repo.Update(&t)
}

func (srv *service) PauseTask(name string) error {
	t, err := srv.GetTask(name)
	if err != nil {
		return err
	}

	if t.Finished {
		return task.ErrTaskFinished
	}

	slotsLen := len(t.TimeSlots)
	if slotsLen == 0 {
		return task.ErrNoTimeSlots
	}

	if t.TimeSlots[slotsLen-1].End != nil {
		return task.ErrSlotFinished
	}

	now := time.Now()
	t.TimeSlots[slotsLen-1].End = &now

	return srv.repo.Update(&t)
}

func (srv *service) GetTask(name string) (task.Task, error) {
	var t task.Task
	err := srv.repo.Get(name, &t)
	return t, err
}

func (srv *service) ListTasks() ([]task.Task, error) {
	var ts []task.Task
	err := srv.repo.List(&ts)
	return ts, err
}
