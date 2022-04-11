package memkvs

import (
	"go.gllm.dev/trackr/domain/errors"
	"go.gllm.dev/trackr/domain/task"
)

type memkvs struct {
	kvs map[string]task.Task
}

func NewMemKVS() *memkvs {
	return &memkvs{
		kvs: make(map[string]task.Task),
	}
}

func (m *memkvs) Create(t *task.Task) error {
	if _, exists := m.kvs[t.Name]; exists {
		return errors.ErrTaskExists
	}
	m.kvs[t.Name] = *t
	return nil
}

func (m *memkvs) Get(name string) (*task.Task, error) {
	if t, exists := m.kvs[name]; exists {
		return &t, nil
	}
	return nil, errors.ErrTaskNotExists
}
func (m *memkvs) Update(t *task.Task) error {
	if _, exists := m.kvs[t.Name]; exists {
		m.kvs[t.Name] = *t
		return nil
	}
	return errors.ErrTaskNotExists
}
func (m *memkvs) List() ([]task.Task, error) {
	var tasks []task.Task
	for _, t := range m.kvs {
		tasks = append(tasks, t)
	}
	return tasks, nil
}
