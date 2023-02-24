package retextaigo

import (
	"context"

	"github.com/karalef/retextaigo/api"
)

// Summarize generates summary text for source.
func (c *Client) Summarize(ctx context.Context, source string, maxLength int) (*Task[api.Summarized], error) {
	if maxLength < 1 {
		maxLength = 40
	}
	return queue[api.Summarized](ctx, c, api.TaskSummarize, source, map[string]any{
		"max_length": maxLength,
	})
}

// AwaitSummarize generates summary text for source.
func (c *Client) AwaitSummarize(ctx context.Context, source string, maxLength int) (*Result[api.Summarized], error) {
	t, err := c.Summarize(ctx, source, maxLength)
	if err != nil {
		return nil, err
	}
	return t.WaitContext(ctx)
}
