package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	endpoint   string
	apiToken   string
	httpClient *http.Client
}

// NewClient creates common settings
func NewClient(endpoint string, apitoken string) *Client {

	return &Client{
		endpoint:   endpoint,
		apiToken:   apitoken,
		httpClient: &http.Client{},
	}
}

func (c *Client) SendRequest(method string, path string, payload interface{}, statusCode int) (value string, err error) {
	url := c.endpoint + path
	client := &http.Client{}

	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "token:"+c.apiToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()

	strbody := string(body)

	if statusCode != 0 {
		if resp.StatusCode != statusCode {

			return "", fmt.Errorf("[ERROR] unexpected status code got: %v expected: %v \n %v", resp.StatusCode, statusCode, strbody)
		}
	}

	return strbody, nil
}
