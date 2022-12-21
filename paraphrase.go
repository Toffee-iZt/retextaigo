package retextaigo

import "github.com/karalef/retextaigo/api"

// Paraphrase generates paraphrased text for source.
func (c *Client) Paraphrase(source string, lang string) (*Task[[]string], error) {
	var add map[string]any
	if lang != "" {
		add = map[string]any{"lang": lang}
	}
	return queue[[]string](c, api.TaskParaphrase, source, add)
}
