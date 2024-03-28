package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"time"
)

// HttpPost PostJson 请求
func HttpPost(link string, params map[string]string, json []byte) ([]byte, error) {
	client := &http.Client{Timeout: 20 * time.Second}
	//忽略https的证书
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	p := url.Values{}
	u, _ := url.Parse(link)
	if params != nil {
		for k, v := range params {
			p.Set(k, v)
		}
	}
	u.RawQuery = p.Encode()
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d:%s", resp.StatusCode, resp.Status)
	}
	return io.ReadAll(resp.Body)
}

// HttpGet Get 请求  link：请求url
func HttpGet(link string, params map[string]string) ([]byte, error) {
	client := &http.Client{Timeout: 20 * time.Second}
	//忽略https的证书
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	p := url.Values{}
	u, _ := url.Parse(link)
	if params != nil {
		for k, v := range params {
			p.Set(k, v)
		}
	}
	u.RawQuery = p.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d:%s", resp.StatusCode, resp.Status)
	}
	return io.ReadAll(resp.Body)
}

// HttpPostJson 请求
func HttpPostJson(reqUrl string, reqParam string, signParam string) ([]byte, error) {
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer([]byte(reqParam)))
	if err != nil {
		log.Errorf("HttpPostJson http.NewRequest err=[%s]", err.Error())
		return nil, err
	}
	req.Header.Add("x-sign", signParam)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("http post json err=[%s]", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d:%s", resp.StatusCode, resp.Status)
	}
	return io.ReadAll(resp.Body)
}

// HttpComPostJson 公共post请求
func HttpComPostJson(reqUrl string, reqParam string, header map[string]string, proxy string) ([]byte, error) {
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer([]byte(reqParam)))
	if err != nil {
		log.Errorf("HttpPostJson http.NewRequest err=[%s]", err.Error())
		return nil, err
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	if proxy != "" {
		proxyUrl, err := url.Parse(proxy)
		if err != nil {
			panic(err)
		}
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("http post json err=[%s]", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d:%s", resp.StatusCode, resp.Status)
	}
	return io.ReadAll(resp.Body)
}

// HttpPubPost PostJson 请求
func HttpPubPost(link string, headers map[string]string, params map[string]string, json []byte) ([]byte, error) {
	client := &http.Client{Timeout: 20 * time.Second}
	//忽略https的证书
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	p := url.Values{}
	u, _ := url.Parse(link)
	if params != nil {
		for k, v := range params {
			p.Set(k, v)
		}
	}

	u.RawQuery = p.Encode()
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d:%s", resp.StatusCode, resp.Status)
	}
	return io.ReadAll(resp.Body)
}
