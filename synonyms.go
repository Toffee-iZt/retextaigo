package retextaigo

import (
	"github.com/karalef/retextaigo/api"
)

// Synonyms is synonyms response.
type Synonyms struct {
	Range   []int    `json:"range"`
	Synonym []string `json:"synonym"`
}

// Synonyms generates synonyms for source.
func (c *Client) Synonyms(source string, lang ...string) (*Task[Synonyms], error) {
	l, err := c.lang(source, api.TaskSynonyms, lang...)
	if err != nil {
		return nil, err
	}
	return queue[Synonyms](c, api.TaskSynonyms, source, map[string]any{
		"lang": l,
	})
}
