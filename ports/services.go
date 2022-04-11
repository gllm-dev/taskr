package ports

import "go.gllm.dev/trackr/domain/task"

type TaskrService interface {
	AddTask(name string, tags ...string) error
	FinishTask(name string) error
	ResumeTask(name string) error
	PauseTask(name string) error
	GetTask(name string) (task.Task, error)
	ListTasks() ([]task.Task, error)
}
