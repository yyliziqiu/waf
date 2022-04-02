package tsq

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var (
	serviceConfigMap map[string]ServiceConfig

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

func Initialize(apiTokenName string, apiToken string, configs ...ServiceConfig) {
	ApiTokenName, ApiToken = apiTokenName, apiToken

	serviceConfigMap = make(map[string]ServiceConfig, len(configs))
	for _, config := range configs {
		serviceConfigMap[config.Name] = config
	}
}

func ToUrl(serviceName string) string {
	return serviceConfigMap[serviceName].ToUrl()
}

func JoinUrl(serviceName string, postfix string) string {
	return serviceConfigMap[serviceName].JoinUrl(postfix)
}

func ServiceGet(serviceName string, postfix string, result interface{}) error {
	return Get(JoinUrl(serviceName, postfix), result)
}

func Get(url string, result interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	return Do(req, result)
}

func ServicePost(serviceName string, postfix string, contentType string, body io.Reader, result interface{}) error {
	return Post(JoinUrl(serviceName, postfix), contentType, body, result)
}

func Post(url string, contentType string, body io.Reader, result interface{}) error {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", contentType)

	return Do(req, result)
}

func ServicePut(serviceName string, postfix string, contentType string, body io.Reader, result interface{}) error {
	return Put(JoinUrl(serviceName, postfix), contentType, body, result)
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
