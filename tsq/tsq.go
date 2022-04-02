package tsq

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var (
	ApiTokenName string
	ApiToken     string

	MaxIdleConns        = 200
	MaxConnsPerHost     = 20
	MaxIdleConnsPerHost = 10
	IdleConnTimeout     = 600 * time.Second

	TLSHandshakeTimeout   = 10 * time.Second
	ExpectContinueTimeout = 5 * time.Second
	ResponseHeaderTimeout = 5 * time.Second

	Timeout = 20 * time.Second
)

var client = http.Client{
	Transport: &http.Transport{
		Proxy:       http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{Timeout: 5 * time.Second, KeepAlive: 60 * time.Second}).DialContext,

		MaxIdleConns:        MaxIdleConns,
		MaxConnsPerHost:     MaxConnsPerHost,
		MaxIdleConnsPerHost: MaxIdleConnsPerHost,
		IdleConnTimeout:     IdleConnTimeout,

		TLSHandshakeTimeout:   TLSHandshakeTimeout,
		ExpectContinueTimeout: ExpectContinueTimeout,
		ResponseHeaderTimeout: ResponseHeaderTimeout,

		DisableKeepAlives: false,
		ForceAttemptHTTP2: false,
	},

	Timeout: Timeout,
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e errorResponse) ToString() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func Get(url string, result interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	return Do(req, result)
}

func Post(url string, contentType string, body io.Reader, result interface{}) error {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", contentType)

	return Do(req, result)
}

func Put(url string, contentType string, body io.Reader, result interface{}) error {
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", contentType)

	return Do(req, result)
}

func Do(req *http.Request, result interface{}) error {
	req.Header.Set(ApiTokenName, ApiToken)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		errorResponse := errorResponse{}
		err = json.Unmarshal(buf, &errorResponse)
		if err != nil {
			return err
		}
		return errors.New(errorResponse.ToString())
	} else {
		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		if result == nil {
			return nil
		}
		err = json.Unmarshal(buf, result)
		if err != nil {
			return err
		}
		return nil
	}
}
