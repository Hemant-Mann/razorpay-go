package razorpay

import (
	"bytes"
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

func (c *Client) Get(path string) (*http.Response, error) {
	return c.makeRequest("GET", path, nil, nil)
}

func (c *Client) Post(path string, body []byte) (*http.Response, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	return c.makeRequest("POST", path, bytes.NewBuffer(body), headers)
}

func (c *Client) Delete(path string, body []byte) (*http.Response, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	return c.makeRequest("DELETE", path, bytes.NewBuffer(body), headers)
}
