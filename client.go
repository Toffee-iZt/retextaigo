package retextaigo

import (
	"context"
	"net/http"

	"github.com/karalef/retextaigo/api"
)

// New creates new client.
func New(http *http.Client) *Client {
	return &Client{
		api: api.New(http),
	}
}

// Client provides high level retext.ai API.
type Client struct {
	api *api.API
}

// IsAvailable checks if retext.ai is available.
func (c *Client) IsAvailable(ctx ...context.Context) (bool, error) {
	return c.api.IsAvailable(ctx...)
}

// Tokenize text.
func (c *Client) Tokenize(ctx context.Context, source string) (*api.Tokenized, error) {
	r, err := c.api.Tokenize(ctx, source)
	if err != nil {
		return nil, err
	}
	if r.Status != api.StatusOK {
		return nil, Error{Status: r.Status}
	}
	return r.Data, nil
}

func (c *Client) lang(ctx context.Context, source string, specified ...string) (string, error) {
	if len(specified) > 0 {
		return specified[0], nil
	}
	t, err := c.Tokenize(ctx, source)
	if err != nil {
		return "", err
	}
	return t.SourceLang, nil
}

// Error is error returned by retext.ai.
type Error struct {
	Status string
}

func (e Error) Error() string {
	return e.Status
}
