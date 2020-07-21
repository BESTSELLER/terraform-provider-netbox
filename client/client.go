package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

type Client struct {
	endpoint string
	apiToken string
}

// NewClient creates common settings
func NewClient(endpoint string, apitoken string) *Client {

	return &Client{
		endpoint: endpoint,
		apiToken: apitoken,
	}
}

const pathAvailablePrefixes = "/ipam/prefixes/"

// GetAvailablePrefix will return all available prefixes
func (client *Client) GetAvailablePrefix(id string) (*GetAvailablePrefixResponse, error) {

	resp, err := client.sendRequest("GET", id, nil, 200)
	if err != nil {
		return nil, err
	}
	var jsonData ReponseAvailablePrefixes
	json.Unmarshal([]byte(resp), &jsonData)

	re := regexp.MustCompile(`(?m)(?:[0-9]{1,3}\.){3}[0-9]{1,3}/`)
	prefixLength, _ := strconv.Atoi(re.ReplaceAllString(jsonData.Prefix, ""))

	resp2, err := client.sendRequest("GET", pathAvailablePrefixes+"?q="+jsonData.Prefix, nil, 200)
	if err != nil {
		return nil, err
	}
	var jsonData2 ResponeListOfPrefixes
	err = json.Unmarshal([]byte(resp2), &jsonData2)
	if err != nil {
		return nil, err
	}

	returnValue := &GetAvailablePrefixResponse{
		PrefixLength:   prefixLength,
		Description:    jsonData.Description,
		ID:             pathAvailablePrefixes + strconv.Itoa(jsonData.ID) + "/",
		ParentPrefixID: jsonData2.Results[0].ID,
		Prefix:         jsonData.Prefix,
		PrefixID:       jsonData.ID,
	}
	return returnValue, nil
}

func (client *Client) sendRequest(method string, path string, payload interface{}, statusCode int) (value string, err error) {
	url := client.endpoint + path
	httpClient := &http.Client{}

	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "token "+client.apiToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
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

			return "", fmt.Errorf("[ERROR] unexpected status code got: %v expected: %v  \n %v  \n %v", resp.StatusCode, statusCode, strbody, url)
		}
	}

	return strbody, nil
}
