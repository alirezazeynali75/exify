package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HttpClient struct {
	baseUrl        *url.URL
	defaultHeaders map[string]string
	basicAuth      *BasicAuthentication
}

type BasicAuthentication struct {
	Username string
	Password string
}

func NewClient(baseUrl string, defaultHeaders map[string]string, basicAuth *BasicAuthentication) *HttpClient {
	parsedUrl, err := url.Parse(baseUrl)
	if err != nil {
		panic(err)
	}

	return &HttpClient{
		baseUrl:        parsedUrl,
		defaultHeaders: defaultHeaders,
		basicAuth:      basicAuth,
	}
}

func (hc *HttpClient) getFullUrl(uri string) string {
	return hc.baseUrl.JoinPath(uri).String()
}

func (hc *HttpClient) Post(ctx context.Context, uri string, body string, header map[string]string) (string, error) {
	fullUrl := hc.getFullUrl(uri)
	payload := strings.NewReader(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullUrl, payload)
	if err != nil {
		return "", err
	}

	if hc.basicAuth != nil {
		req.SetBasicAuth(hc.basicAuth.Username, hc.basicAuth.Password)
	}

	for key, value := range hc.defaultHeaders {
		req.Header.Add(key, value)
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode >= 400 || res.StatusCode < 200 {
		return "", fmt.Errorf("request from server failed with status %s and response body is %s", res.Status, string(resBody))
	}
	return string(resBody), nil
}

func (hc *HttpClient) Get(ctx context.Context, uri string, header map[string]string) (string, error) {
	fullUrl := hc.getFullUrl(uri)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullUrl, nil)
	if err != nil {
		return "", err
	}

	for key, value := range hc.defaultHeaders {
		req.Header.Add(key, value)
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}
