// Package yolpweather provides a client for the YOLP weather API,
// defined as https://developer.yahoo.co.jp/webapi/map/openlocalplatform/v1/weather.html
//
package yolpweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const Endpoint = "https://map.yahooapis.jp/weather/V1/place"

// New returns a new Client.
func New(clientID string) *Client {
	return &Client{ClientID: clientID}
}

// Client provides access to YOLP weather API.
type Client struct {
	ClientID string       // YOLP API Client ID (required)
	Client   *http.Client // Default to http.DefaultClient
	Endpoint string       // Default to Endpoint
}

// Get sends the request and returns the response.
func (c *Client) Get(req *Request) (*Response, error) {
	endpoint := c.Endpoint
	if endpoint == "" {
		endpoint = Endpoint
	}

	q := req.Values().Encode()
	hreq, err := http.NewRequest("GET", endpoint+"?"+q, nil)
	if err != nil {
		return nil, fmt.Errorf("error while creating a HTTP request: %s", err)
	}
	hreq.Header.Set("user-agent", "Yahoo AppID: "+c.ClientID)

	client := c.Client
	if client == nil {
		client = http.DefaultClient
	}
	hresp, err := client.Do(hreq)
	if err != nil {
		return nil, fmt.Errorf("error while sending a HTTP request: %s", err)
	}
	defer hresp.Body.Close()
	if hresp.StatusCode != 200 {
		return nil, fmt.Errorf("server returned status code %d", hresp.StatusCode)
	}

	var resp Response
	expires, err := http.ParseTime(hresp.Header.Get("expires"))
	if err == nil {
		resp.Expires = expires
	}
	d := json.NewDecoder(hresp.Body)
	if err := d.Decode(&resp.Payload); err != nil {
		return nil, fmt.Errorf("error while decoding JSON: %s", err)
	}
	return &resp, nil
}
