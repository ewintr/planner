package service

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Todoist struct {
	apiKey    string
	baseURL   string
	client    *http.Client
	syncToken string
	done      chan bool
}

func NewTodoist(apiKey, baseURL string) *Todoist {
	td := &Todoist{
		apiKey:  apiKey,
		baseURL: baseURL,
		done:    make(chan bool),
		client:  http.DefaultClient,
	}

	return td
}

func (td *Todoist) Run() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-td.done:
			return
		case <-ticker.C:
			fmt.Println("hoi")
			if err := td.Sync(); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (td *Todoist) Sync() error {
	if td.syncToken == "" {
		return td.FullSync()
	}

	return nil
}

func (td *Todoist) FullSync() error {
	res := td.do(http.MethodGet)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

type tdResp struct {
	Status int
	Body   []byte
	Error  error
}

func (td *Todoist) do(method string) tdResp {
	u, err := url.Parse(fmt.Sprintf("%s/sync/v9/sync", td.baseURL))
	if err != nil {
		return tdResp{
			Error: err,
		}
	}
	formData := url.Values{
		"syncToken":      {"*"},
		"resource_types": {`["projects"]`},
	}
	u.RawQuery = formData.Encode()
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return tdResp{
			Error: err,
		}
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", td.apiKey))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := td.client.Do(req)
	if err != nil {
		return tdResp{
			Error: err,
		}
	}

	if res.StatusCode != http.StatusOK {
		return tdResp{
			Error: fmt.Errorf("status code: %d", res.StatusCode),
		}
	}

	var body []byte
	if body, err = io.ReadAll(res.Body); err != nil {
		return tdResp{
			Error: err,
		}
	}
	return tdResp{
		Body: body,
	}

}
