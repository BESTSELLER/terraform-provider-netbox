package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/BESTSELLER/terraform-provider-netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Client holds the client info for netbox
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

// AvailablePrefixBody returns the body needed to get available prefix
func AvailablePrefixBody(d *schema.ResourceData) models.AvailablePrefixes {
	return models.AvailablePrefixes{
		PrefixLength: d.Get("prefix_length").(int),
		Site:         d.Get("site").(int),
		Tenant:       d.Get("tenant").(int),
		Status:       d.Get("status").(string),
		Role:         d.Get("role").(int),
		Description:  d.Get("description").(string),
	}
}

// GetPrefix will return a prefix
func (client *Client) GetPrefix(newCidr string) (*models.ResponsePrefix, error) {
	if newCidr == "" {
		return nil, fmt.Errorf("[ERROR] 'cidr_notation' is empty")
	}
	prefixPath := models.PathAvailablePrefixes + "?prefix=" + newCidr

	resp, err := client.SendRequest("GET", prefixPath, nil, 200)
	if err != nil {
		return nil, err
	}

	var jsonData models.ResponsePrefix
	err = json.Unmarshal([]byte(resp), &jsonData)
	if err != nil {
		return nil, err
	}

	return &jsonData, nil
}

// GetSite will return a site by name
func (client *Client) GetSite(siteName string) (*models.ResponseSites, error) {
	sitesPath := "/dcim/sites/?name=" + siteName

	resp, err := client.SendRequest("GET", sitesPath, nil, 200)
	if err != nil {
		return nil, err
	}

	var jsonData models.ResponseSites
	err = json.Unmarshal([]byte(resp), &jsonData)
	if err != nil {
		return nil, err
	}

	return &jsonData, nil
}

// DeletePrefix will delete a given prefix
func (client *Client) DeletePrefix(d *schema.ResourceData) error {
	_, err := client.SendRequest("DELETE", d.Id(), nil, 204)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePrefix will patch a prefix
func (client *Client) UpdatePrefix(d *schema.ResourceData) error {
	body := AvailablePrefixBody(d)
	_, err := client.SendRequest("PATCH", d.Id(), body, 200)
	if err != nil {
		return err
	}
	return nil
}

// CreatePrefix will create the prefix
func (client *Client) CreatePrefix(body *models.AvailablePrefixes, parentID int) (*models.ReponseAvailablePrefixes, error) {
	path := fmt.Sprintf("%s%d/available-prefixes/", models.PathAvailablePrefixes, parentID)

	resp, err := client.SendRequest("POST", path, body, 201)
	if err != nil {
		return nil, err
	}
	var jsonData models.ReponseAvailablePrefixes
	err = json.Unmarshal([]byte(resp), &jsonData)
	if err != nil {
		return nil, err
	}
	return &jsonData, nil
}

// GetAvailablePrefix will return all available prefixes
func (client *Client) GetAvailablePrefix(id string) (*models.GetAvailablePrefixResponse, error) {
	resp, err := client.SendRequest("GET", id, nil, 200)
	if err != nil {
		return nil, err
	}
	var jsonData models.ReponseAvailablePrefixes
	err = json.Unmarshal([]byte(resp), &jsonData)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`(?m)(?:[0-9]{1,3}\.){3}[0-9]{1,3}/`)
	prefixLength, _ := strconv.Atoi(re.ReplaceAllString(jsonData.Prefix, ""))

	resp2, err := client.SendRequest("GET", models.PathAvailablePrefixes+"?q="+jsonData.Prefix, nil, 200)
	if err != nil {
		return nil, err
	}
	var jsonData2 models.ResponeListOfPrefixes
	err = json.Unmarshal([]byte(resp2), &jsonData2)
	if err != nil {
		return nil, err
	}

	returnValue := &models.GetAvailablePrefixResponse{
		PrefixLength: prefixLength,
		Description:  jsonData.Description,
		// ID:             models.PathAvailablePrefixes + strconv.Itoa(jsonData.ID) + "/",
		ID:             fmt.Sprintf("%s%d/", models.PathAvailablePrefixes, jsonData.ID),
		ParentPrefixID: jsonData2.Results[0].ID,
		Prefix:         jsonData.Prefix,
		PrefixID:       jsonData.ID,
	}
	return returnValue, nil
}

func (client *Client) SendRequest(method string, path string, payload interface{}, statusCode int) (value string, err error) {
	baseUrl, err := url.Parse(client.endpoint)
	if err != nil {
		return "", err
	}

	baseUrl = baseUrl.JoinPath("api")

	url := fmt.Sprintf("%s/%s", baseUrl.String(), path)

	httpClient := &http.Client{}

	var requestBody io.Reader
	if payload != nil {
		b := new(bytes.Buffer)
		err = json.NewEncoder(b).Encode(payload)
		if err != nil {
			return "", err
		}
		requestBody = b
	}

	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "token "+client.apiToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
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

func (client *Client) GetDeviceType(id int) (*models.ResponseDeviceTypes, error) {
	path := fmt.Sprintf("%s%d/", models.PathDeviceTypes, id)
	resp, err := client.SendRequest("GET", path, nil, 200)
	if err != nil {
		return nil, err
	}

	var jsonData models.ResponseDeviceTypes
	err = json.Unmarshal([]byte(resp), &jsonData)
	if err != nil {
		return nil, err
	}

	return &jsonData, nil
}
