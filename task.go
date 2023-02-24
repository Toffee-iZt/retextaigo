package retextaigo

import (
	"context"
	"time"

	"github.com/karalef/retextaigo/api"
)

func queue[T any](ctx context.Context, c *Client, task api.TaskType, source string, add map[string]any) (*Task[T], error) {
	resp, err := c.api.QueueTask(ctx, task, source, add)
	if err != nil {
		return nil, err
	}
	if resp.Status != api.StatusOK {
		return nil, Error{resp.Status}
	}
	return &Task[T]{
		c:    c,
		id:   resp.Data.TaskID,
		lang: resp.Data.SourceLang,
	}, nil
}

// Task represents a task in the queue.
type Task[T any] struct {
	c    *Client
	id   string
	lang string
}

func (t *Task[T]) ID() string {
	return t.id
}

func (t *Task[T]) Lang() string {
	return t.lang
}

func (t *Task[T]) Check(ctx context.Context) (*api.Checked[T], error) {
	resp, err := api.QueueCheck[T](ctx, t.c.api, t.id)
	if err != nil {
		return nil, err
	}
	if resp.Status != api.StatusOK {
		return nil, Error{resp.Status}
	}
	return resp.Data, nil
}

func (t *Task[T]) Wait(interval ...time.Duration) (*Result[T], error) {
	return t.WaitContext(context.Background(), interval...)
}

func (t *Task[T]) WaitContext(ctx context.Context, interval ...time.Duration) (*Result[T], error) {
	inter := time.Second * 1
	if len(interval) > 0 {
		inter = interval[0]
	}
	for {
		resp, err := t.Check(ctx)
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		if err != nil {
			return nil, err
		}
		if resp.Ready {
			return &Result[T]{Successful: resp.Successful, Result: resp.Result}, nil
		}
		time.Sleep(inter)
	}
}

// Result is a result of a task.
type Result[T any] struct {
	Successful bool
	Result     T
}
