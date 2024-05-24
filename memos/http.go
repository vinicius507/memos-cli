package memos

import (
	"net"
	"net/http"
	"time"
)

func newHttpClient() *http.Client {
	dialer := &net.Dialer{Timeout: time.Second}
	transport := &http.Transport{DialContext: dialer.DialContext}

	return &http.Client{
		Timeout:   5 * time.Second,
		Transport: transport,
	}
}
