package api

import "net/url"

const queueEndpoint = "queue_task"

const queueCheckEndpoint = "queue_check"

// TaskType is type of task.
type TaskType string

// task types.
const (
	TaskSynonyms   TaskType = "synonyms"
	TaskParaphrase TaskType = "paraphrase"
	TaskSummarize  TaskType = "summarize"
	TaskExtension  TaskType = "expand"
)

// Queued type.
type Queued struct {
	TaskID     string `json:"taskId"`
	SourceLang string `json:"source_lang"`
}

// QueueTask adds source to queue.
func (a *API) QueueTask(task TaskType, source string, add map[string]any) (*Response[Queued], error) {
	return post[Queued](a, queueEndpoint, map[string]any{
		"task":        task,
		"source_text": source,
		"additional":  add,
	})
}

// Checked type.
type Checked[T any] struct {
	Ready      bool `json:"ready"`
	Successful bool `json:"successful"`
	Result     *T   `json:"result"`
}

// QueueCheck checks task for completion.
func QueueCheck[T any](a *API, taskID string) (*Response[Checked[T]], error) {
	return get[Checked[T]](a, queueCheckEndpoint, url.Values{
		"taskId": {taskID},
	})
}

// QueueCheck checks task for completion.
func (a *API) QueueCheck(taskID string) (*Response[Checked[any]], error) {
	return QueueCheck[any](a, taskID)
}
