package api

import "context"

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
func (a *API) Tokenize(ctx context.Context, source string) (*Response[Tokenized], error) {
	return post[Tokenized](ctx, a, tokenizeEndpoint, map[string]any{
		"source": source,
	})
}
