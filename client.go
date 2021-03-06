package razorpay

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// BaseURL stores the API base URL
const BaseURL = "https://api.razorpay.com/v1"

// Client struct
type Client struct {
	Key    string
	Secret string

	httpClient *http.Client
}

// Resource interface is to be used for generic decoding of object
type Resource interface {
	New() Resource
	Endpoint() string
}

// NewClient returns a pointer to the razorpay client
func NewClient(key, secret string) *Client {
	var c = &Client{
		Key:    key,
		Secret: secret,
	}
	c.httpClient = getHTTPClient(10)
	return c
}

func getHTTPClient(timeout int) *http.Client {
	var httpClient = &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	return httpClient
}

func sendResp(resp *http.Response, err error, rs Resource) (Resource, error) {
	var newResource = rs.New()
	if err != nil {
		return newResource, err
	}
	body, readErr := readBody(resp)
	if readErr != nil {
		return newResource, readErr
	}
	parseError := json.Unmarshal(body, newResource)
	return newResource, parseError
}

func readBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

func (c *Client) getURL(path string) string {
	return BaseURL + path
}

func (c *Client) makeRequest(method, path string, body *bytes.Buffer, headers map[string]string) (*http.Response, error) {
	req, _ := http.NewRequest(method, c.getURL(path), body)
	req.SetBasicAuth(c.Key, c.Secret)
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	resp, err := c.httpClient.Do(req)
	return resp, err
}

// Get method makes a GET request to the resource
func (c *Client) Get(path string) (*http.Response, error) {
	var body = ""
	return c.makeRequest("GET", path, bytes.NewBuffer([]byte(body)), nil)
}

// Post method makes a POST and sends data in json format
func (c *Client) Post(path string, body []byte) (*http.Response, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	return c.makeRequest("POST", path, bytes.NewBuffer(body), headers)
}

// Delete method makes a DELETE request to the resource
func (c *Client) Delete(path string, body []byte) (*http.Response, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	return c.makeRequest("DELETE", path, bytes.NewBuffer(body), headers)
}

// Patch method makes a PATCH and sends data in json format
func (c *Client) Patch(path string, body []byte) (*http.Response, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	return c.makeRequest("PATCH", path, bytes.NewBuffer(body), headers)
}
