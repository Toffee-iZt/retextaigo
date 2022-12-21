package api

const tokenizeEndpoint = "tokenize"

// Tokenized contains tokenize result.
type Tokenized struct {
	Source []struct {
		Delimiter bool   `json:"delimiter"`
		Data      string `json:"data"`
	} `json:"source"`
	SourceLang string `json:"source_lang"`
}

// Tokenize text.
func (a *API) Tokenize(source string, requestFrom string) (*Response[Tokenized], error) {
	return post[Tokenized](a, tokenizeEndpoint, map[string]any{
		"request_from": requestFrom,
		"source":       source,
	})
}
