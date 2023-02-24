package retextaigo

import (
	"context"

	"github.com/karalef/retextaigo/api"
)

// Paraphrase generates paraphrased text for source.
func (c *Client) Paraphrase(ctx context.Context, source string, lang ...string) (*Task[api.Paraphrased], error) {
	l, err := c.lang(ctx, source, lang...)
	if err != nil {
		return nil, err
	}
	return queue[api.Paraphrased](ctx, c, api.TaskParaphrase, source, map[string]any{
		"lang": l,
	})
}

// AwaitParaphrase generates paraphrased text for source.
func (c *Client) AwaitParaphrase(ctx context.Context, source string, lang ...string) (*Result[api.Paraphrased], error) {
	t, err := c.Paraphrase(ctx, source, lang...)
	if err != nil {
		return nil, err
	}
	return t.WaitContext(ctx)
}
