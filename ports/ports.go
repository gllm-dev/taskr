package ports

import "go.gllm.dev/trackr/domain/task"

type Service interface {
	AddTask(name string, tags ...string) error
	FinishTask(name string) error
	ResumeTask(name string) error
	PauseTask(name string) error
	GetTask(name string) (task.Task, error)
	ListTasks() ([]task.Task, error)
}

type Repo interface {
	Create(t *task.Task) error
	Get(name string, t *task.Task) error
	Update(t *task.Task) error
	List(ts *[]task.Task) error
}
