package taskrrepo

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"go.gllm.dev/trackr/domain/task"
)

type repo struct {
	db *bbolt.DB
}

func New() (*repo, error) {
	db, err := bbolt.Open("my.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	return &repo{db: db}, db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		return err
	})
}

func (r *repo) Create(t *task.Task) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		return b.Put([]byte(t.Name), buf)
	})
}

func (r *repo) Get(name string, t *task.Task) error {
	return r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		v := b.Get([]byte(name))
		if v == nil {
			return task.ErrTaskNotExists
		}
		return json.Unmarshal(v, &t)
	})
}
func (r *repo) Update(t *task.Task) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		return b.Put([]byte(t.Name), buf)
	})
}
func (r *repo) List(ts *[]task.Task) error {
	return r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		if err := b.ForEach(func(k, v []byte) error {
			var t task.Task
			if err := json.Unmarshal(v, &t); err != nil {
				return err
			}
			*ts = append(*ts, t)
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
}
