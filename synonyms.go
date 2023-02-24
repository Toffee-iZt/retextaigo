package retextaigo

import (
	"context"

	"github.com/karalef/retextaigo/api"
)

// Synonyms generates synonyms for source.
func (c *Client) Synonyms(ctx context.Context, source string, lang ...string) (*Task[api.Synonyms], error) {
	l, err := c.lang(ctx, source, lang...)
	if err != nil {
		return nil, err
	}
	return queue[api.Synonyms](ctx, c, api.TaskSynonyms, source, map[string]any{
		"lang": l,
	})
}

// AwaitSynonyms generates synonyms for source.
func (c *Client) AwaitSynonyms(ctx context.Context, source string, lang ...string) (*Result[api.Synonyms], error) {
	t, err := c.Synonyms(ctx, source, lang...)
	if err != nil {
		return nil, err
	}
	return t.WaitContext(ctx)
}
