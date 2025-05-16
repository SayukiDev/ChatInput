package tasks

import (
	"errors"
	"github.com/puzpuzpuz/xsync/v4"
)

type Tasks struct {
	runners *xsync.Map[string, *Runner]
}

func New() *Tasks {
	return &Tasks{
		runners: xsync.NewMap[string, *Runner](),
	}
}

func (t *Tasks) Add(name string, handle Handle) error {
	r := NewRunner(handle)
	_, e := t.runners.LoadAndStore(name, r)
	if e {
		r.Close()
		return errors.New("the task is already added")
	}
	r.Start()
	return nil
}

func (t *Tasks) Remove(name string) error {
	r, e := t.runners.LoadAndDelete(name)
	if !e {
		return errors.New("the task not exist")
	}
	r.Close()
	return nil
}
