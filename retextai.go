package retextaigo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

const api = "https://api.retext.ai/api/v1/"

// StatusOK is normal status.
const StatusOK = "ok"

// NewClient makes new client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &Client{
		http: httpClient,
	}
}

// Client for retext.ai.
type Client struct {
	http *http.Client
}

type responseType struct {
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
}

func (c *Client) do(response any, req *http.Request) (string, error) {
	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var r responseType
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", err
	}
	return r.Status, json.Unmarshal(r.Data, response)
}

func (c *Client) get(response any, meth string, q url.Values) (string, error) {
	req, _ := http.NewRequest(http.MethodGet, api+meth+"?"+q.Encode(), nil)
	return c.do(response, req)
}

func (c *Client) post(response any, meth string, data map[string]any) (string, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	req, _ := http.NewRequest(http.MethodPost, api+meth, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	return c.do(response, req)
}

func (c *Client) options(meth string) (int, error) {
	req, _ := http.NewRequest(http.MethodOptions, api+meth, nil)
	resp, err := c.http.Do(req)
	return resp.StatusCode, err
}

// IsAvailable returns true if all endpoints are available.
func (c *Client) IsAvailable() (bool, error) {
	endpoints := []string{"tokenize", "queue", "queue_paraphrase", "queue_check"}
	for _, e := range endpoints {
		status, err := c.options(e)
		if err != nil || status != http.StatusOK {
			return false, err
		}
	}
	return true, nil
}
