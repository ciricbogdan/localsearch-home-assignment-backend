package client

import "net/http"

// Client defines a wrapper around http.Client
type Client struct {
	http.Client
	URL string
}

// New retuns a new instance of the Client
func New(url string) (*Client, error) {
	c := http.Client{}

	return &Client{Client: c, URL: url}, nil
}

// Get wraps around http.Client method to call on clients url + path params if needed
func (c *Client) Get(pathParts ...string) (*http.Response, error) {

	url := c.URL

	for _, pathPart := range pathParts {
		url += "/" + pathPart
	}

	return c.Client.Get(url)
}
