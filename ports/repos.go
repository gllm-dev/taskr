package ports

import "go.gllm.dev/trackr/domain/task"

type TaskrRepository interface {
	Create(t *task.Task) error
	Get(name string, t *task.Task) error
	Update(t *task.Task) error
	List(ts *[]task.Task) error
}
