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
func (c *Client) Synonyms(source string, lang string) (*Task[Synonyms], error) {
	var add map[string]any
	if lang != "" {
		add = map[string]any{"lang": lang}
	}
	return queue[Synonyms](c, api.TaskSynonyms, source, add)
}
