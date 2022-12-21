package retextaigo

import "github.com/karalef/retextaigo/api"

// Summarize generates summary text for source.
func (c *Client) Summarize(source string, maxLength int) (*Task[string], error) {
	var add map[string]any
	if maxLength > 0 {
		add = map[string]any{"max_length": maxLength}
	}
	return queue[string](c, api.TaskSummarize, source, add)
}
