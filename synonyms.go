package retextaigo

import (
	"github.com/karalef/retextaigo/api"
)

// Synonyms generates synonyms for source.
func (c *Client) Synonyms(source string, lang ...string) (*Task[api.Synonyms], error) {
	l, err := c.lang(source, api.TaskSynonyms, lang...)
	if err != nil {
		return nil, err
	}
	return queue[api.Synonyms](c, api.TaskSynonyms, source, map[string]any{
		"lang": l,
	})
}
