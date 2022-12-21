package retextaigo

import (
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
func (c *Client) IsAvailable() (bool, error) {
	return c.api.IsAvailable()
}

// Tokenize text.
func (c *Client) Tokenize(source string, requestFrom string) (*api.Response[api.Tokenized], error) {
	return c.api.Tokenize(source, requestFrom)
}

// Error is error returned by retext.ai.
type Error struct {
	Status string
}

func (e Error) Error() string {
	return e.Status
}
