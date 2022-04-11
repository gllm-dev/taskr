package service

import (
	"go.gllm.dev/trackr/domain/errors"
	"go.gllm.dev/trackr/domain/task"
	"go.gllm.dev/trackr/ports"
	"time"
)

type Service struct {
	repo ports.Repo
}

func NewService(repo ports.Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (srv *Service) AddTask(name string, tags ...string) error {
	t := task.NewTask(name, tags...)
	return srv.repo.Create(t)
}

func (srv *Service) FinishTask(name string) error {
	t, err := srv.GetTask(name)
	if err != nil {
		return err
	}

	slotsLen := len(t.TimeSlots)
	if slotsLen == 0 {
		return errors.ErrNoTimeSlots
	}

	if t.TimeSlots[slotsLen-1].End == nil {
		now := time.Now()
		t.TimeSlots[slotsLen-1].End = &now
	}
	t.Finished = true
	return srv.repo.Update(&t)
}

func (srv *Service) ResumeTask(name string) error {
	t, err := srv.GetTask(name)
	if err != nil {
		return err
	}

	if t.Finished {
		return errors.ErrTaskFinished
	}
	slotsLen := len(t.TimeSlots)
	if slotsLen == 0 {
		return errors.ErrNoTimeSlots
	}

	if t.TimeSlots[slotsLen-1].End == nil {
		return errors.ErrSlotNotFinished
	}

	t.TimeSlots = append(t.TimeSlots, *task.NewTimeSlot())
	return srv.repo.Update(&t)
}

func (srv *Service) PauseTask(name string) error {
	t, err := srv.GetTask(name)
	if err != nil {
		return err
	}
	slotsLen := len(t.TimeSlots)
	if slotsLen == 0 {
		return errors.ErrNoTimeSlots
	}

	if t.TimeSlots[slotsLen-1].End != nil {
		return errors.ErrSlotFinished
	}

	now := time.Now()
	t.TimeSlots[slotsLen-1].End = &now

	return srv.repo.Update(&t)
}

func (srv *Service) GetTask(name string) (task.Task, error) {
	var t task.Task
	err := srv.repo.Get(name, &t)
	return t, err
}

func (srv *Service) ListTasks() ([]task.Task, error) {
	var ts []task.Task
	err := srv.repo.List(&ts)
	return ts, err
}
