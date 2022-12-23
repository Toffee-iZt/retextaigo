package retextaigo

import "github.com/karalef/retextaigo/api"

// Summarize generates summary text for source.
func (c *Client) Summarize(source string, maxLength int) (*Task[api.Summarized], error) {
	if maxLength < 1 {
		maxLength = 40
	}
	return queue[api.Summarized](c, api.TaskSummarize, source, map[string]any{
		"max_length": maxLength,
	})
}
