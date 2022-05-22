package retextaigo

import (
	"net/url"
)

// Queued type.
type Queued struct {
	TaskID string `json:"taskId"`
}

// QueueSummarization adds source to summarization queue.
func (c *Client) QueueSummarization(source string, maxLength int, lang string) (string, *Queued, error) {
	var data = map[string]any{
		"source":     source,
		"max_length": maxLength,
		"lang":       lang,
	}
	var response Queued
	status, err := c.post(&response, "queue", data)
	if err != nil {
		return "", nil, err
	}
	return status, &response, nil
}

// QueueParaphrase adds source to paraphrase queue.
func (c *Client) QueueParaphrase(source string, lang string) (string, *Queued, error) {
	var data = map[string]any{
		"source": source,
		"lang":   lang,
	}
	var response Queued
	status, err := c.post(&response, "queue_paraphrase", data)
	if err != nil {
		return "", nil, err
	}
	return status, &response, nil
}

// Checked type.
type Checked[ResultType string | []string] struct {
	Ready      bool       `json:"ready"`
	Successful bool       `json:"successful"`
	Result     ResultType `json:"result"`
}

// QueueCheckParaphrase checks paraphrase task for completion.
func (c *Client) QueueCheckParaphrase(taskID string) (string, *Checked[[]string], error) {
	var q = url.Values{
		"taskId": {taskID},
	}
	var response Checked[[]string]
	status, err := c.get(&response, "queue_check", q)
	if err != nil {
		return "", nil, err
	}
	return status, &response, nil
}

// QueueCheckSummarization checks summarization task for completion.
func (c *Client) QueueCheckSummarization(taskID string) (string, *Checked[string], error) {
	var q = url.Values{
		"taskId": {taskID},
	}
	var response Checked[string]
	status, err := c.get(&response, "queue_check", q)
	if err != nil {
		return "", nil, err
	}
	return status, &response, nil
}
