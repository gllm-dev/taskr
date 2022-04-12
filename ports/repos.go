package ports

import "go.gllm.dev/taskr/domain/task"

type TaskrRepository interface {
	Create(t *task.Task) error
	Get(name string, t *task.Task) error
	Update(t *task.Task) error
	List(ts *[]task.Task) error
}
