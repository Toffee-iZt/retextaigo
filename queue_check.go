package retextaigo

import "net/url"

// Checked type.
type Checked[ResultType []string | string] struct {
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
