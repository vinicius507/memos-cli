package memos

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type MemosClient struct {
	ServerAddr  string
	AccessToken string
}

func (c *MemosClient) request(method, path string, body io.Reader) (*http.Response, error) {
	httpClient := newHttpClient()
	url := c.ServerAddr + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.AccessToken)
	return httpClient.Do(req)
}

func (c *MemosClient) NewMemo(content string) (*Memo, error) {
	body := fmt.Sprintf(`{"content": %q}`, content)
	reader := strings.NewReader(body)

	res, err := c.request(http.MethodPost, "/memos", reader)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	rawMemo, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	memo := &Memo{}
	if err := json.Unmarshal(rawMemo, memo); err != nil {
		return nil, err
	}
	return memo, nil
}
