package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dkeng/pkg/http/header"
)

const (
	// 默认超时
	defaultTimeout = time.Second * 3
)

// HTTPClient Http客户端
type HTTPClient struct {
	Header *header.HTTPHeader
	client *http.Client
	// Timeout time.Duration
}

// New 创建HttpClient
func New() *HTTPClient {
	return &HTTPClient{
		Header: new(header.HTTPHeader),
		// Timeout: time.Second * 3,
		client: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

// NewHTTPClient 创建HttpClient
func NewHTTPClient(header *header.HTTPHeader) *HTTPClient {
	return &HTTPClient{
		Header: header,
		client: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

// SetHeader 设置 http request header
func (h *HTTPClient) SetHeader(header *header.HTTPHeader) {
	h.Header = header
}

// SetTimeout 设置超时
func (h *HTTPClient) SetTimeout(timeout time.Duration) {
	h.client.Timeout = timeout
}

// SetTransport 设置Transport
func (h *HTTPClient) SetTransport(transport http.RoundTripper) {
	h.client.Transport = transport
}

// Get send get request.
func (h *HTTPClient) Get(url string, values url.Values) (body []byte, err error) {
	if values != nil {
		url += "?" + values.Encode()
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	for k, v := range *h.Header {
		req.Header.Set(k, v)
	}

	response, err := h.client.Do(req)
	if err != nil {
		return
	}
	return parseResponse(response)
}

// Post send post request.
func (h *HTTPClient) Post(url string, values url.Values) (body []byte, err error) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(values.Encode()))
	if err != nil {
		return
	}
	for k, v := range *h.Header {
		req.Header.Set(k, v)
	}

	response, err := h.client.Do(req)
	if err != nil {
		return
	}
	return parseResponse(response)
}

// PostJSON send post request.
func (h *HTTPClient) PostJSON(url string, jsonObject interface{}) (body []byte, err error) {
	json, err := json.Marshal(jsonObject)
	if err != nil {
		return
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(json))
	if err != nil {
		return
	}
	for k, v := range *h.Header {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")

	response, err := h.client.Do(req)
	if err != nil {
		return
	}
	return parseResponse(response)
}

// parseResponse 解析响应
func parseResponse(response *http.Response) (body []byte, err error) {
	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return
	}
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	return
}
