package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	baseURL string
	apiKey  string
	c       *http.Client
}

func NewClient(url, apiKey string) (*Client, error) {
	return &Client{
		baseURL: url,
		apiKey:  apiKey,
		c: &http.Client{
			Timeout: 10 * time.Second,
		},
	}, nil
}

func (c *Client) Updates(ks []Kind, ts time.Time) ([]Item, error) {
	ksStr := make([]string, 0, len(ks))
	for _, k := range ks {
		ksStr = append(ksStr, string(k))
	}
	u := fmt.Sprintf("%s/sync?ks=", c.baseURL, strings.Join(ksStr, ","))
	if !ts.IsZero() {
		u = fmt.Sprintf("%s&ts=", url.QueryEscape(ts.Format(time.RFC3339)))
	}
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}

	res, err := c.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not get response: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned status %d", res.StatusCode)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %v", err)
	}

	var items []Item
	if err := json.Unmarshal(body, &items); err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %v", err)
	}

	return items, nil
}
