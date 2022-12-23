package retextaigo

import "github.com/karalef/retextaigo/api"

// Paraphrase generates paraphrased text for source.
func (c *Client) Paraphrase(source string, lang ...string) (*Task[api.Paraphrased], error) {
	l, err := c.lang(source, api.TaskParaphrase, lang...)
	if err != nil {
		return nil, err
	}
	return queue[api.Paraphrased](c, api.TaskParaphrase, source, map[string]any{
		"lang": l,
	})
}
