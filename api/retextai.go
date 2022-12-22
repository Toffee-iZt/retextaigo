package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const api = "https://api.retext.ai/api/v1/"

// StatusOK is normal status.
const StatusOK = "ok"

// New makes new api instance.
func New(httpClient *http.Client) *API {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &API{
		http: httpClient,
	}
}

// API provides retext.ai methods.
type API struct {
	http *http.Client
}

// Response represents response from retext.ai.
type Response[T any] struct {
	Status string `json:"status"`
	Data   *T     `json:"data"`
}

func post[T any](a *API, meth string, data map[string]any) (*Response[T], error) {
	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := a.http.Post(api+meth, "application/json", bytes.NewReader(d))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("http: " + resp.Status)
	}

	var response Response[T]
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func get[T any](a *API, meth string, query url.Values) (*Response[T], error) {
	resp, err := a.http.Get(api + meth + "?" + query.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("http: " + resp.Status)
	}

	var response Response[T]
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *API) options(meth string) (int, error) {
	req, _ := http.NewRequest(http.MethodOptions, api+meth, nil)
	resp, err := a.http.Do(req)
	return resp.StatusCode, err
}

// IsAvailable returns true if all endpoints are available.
func (a *API) IsAvailable() (bool, error) {
	endpoints := [...]string{tokenizeEndpoint, queueEndpoint, queueCheckEndpoint}
	for _, e := range endpoints {
		status, err := a.options(e)
		if err != nil || status != http.StatusOK {
			return false, err
		}
	}
	return true, nil
}
