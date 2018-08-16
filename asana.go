package asana

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const userAgent = "Asana Tiny Go v0.1"

type Client struct {
	personalAccessToken string
}

func New(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("no personal access token provided")
	}

	return &Client{token}, nil
}

func (c *Client) CreateTask(tr *TaskRequest) (*Task, error) {
	// prepare post body
	type Request struct {
		Data *TaskRequest `json:"data"`
	}

	body, err := json.Marshal(Request{tr})
	if err != nil {
		return nil, err
	}

	// prepare HTTP request
	req, err := http.NewRequest("POST", "https://app.asana.com/api/1.0/tasks", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}

	// unmarshal response into Task
	var data struct {
		T Task `json:"data"`
	}
	err = json.Unmarshal(res, &data)
	return &data.T, err
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.personalAccessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, errors.New(string(body))
	}

	return body, nil
}
