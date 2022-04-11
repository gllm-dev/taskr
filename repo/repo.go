package repo

import (
	"encoding/binary"
	"encoding/json"
	"go.etcd.io/bbolt"
	"go.gllm.dev/trackr/domain/task"
)

type repo struct {
	db *bbolt.DB
}

func NewRepo() (*repo, error) {
	db, err := bbolt.Open("my.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	return &repo{db: db}, db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		return err
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func (r *repo) Create(t *task.Task) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))

		// Marshal user data into bytes.
		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put([]byte(t.Name), buf)
	})
}

func (r *repo) Get(name string, t *task.Task) error {
	return r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		v := b.Get([]byte(name))
		return json.Unmarshal(v, &t)
	})
}
func (r *repo) Update(t *task.Task) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))

		// Marshal user data into bytes.
		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put([]byte(t.Name), buf)
	})
}
func (r *repo) List(ts *[]task.Task) error {
	return r.db.View(func(tx *bbolt.Tx) error {
		// Assume bucket exists and has keys
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
