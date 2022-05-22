package retextaigo

// Tokenized contains tokenize result.
type Tokenized struct {
	Source []struct {
		Delimiter bool   `json:"delimiter"`
		Data      string `json:"data"`
	} `json:"source"`
	SourceLang string `json:"source_lang"`
}

// Tokenize text.
func (c *Client) Tokenize(source string, lang string) (string, *Tokenized, error) {
	var data = map[string]any{
		"source": source,
		"lang":   lang,
	}
	var response Tokenized
	status, err := c.post(&response, "tokenize", data)
	if err != nil {
		return "", nil, err
	}
	return status, &response, nil
}
