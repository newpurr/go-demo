package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/thinkeridea/go-extend/exbytes"
)

type Adapter struct {
	pool sync.Pool
}

func New() *Adapter {
	return &Adapter{
		pool: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 4096))
			},
		},
	}
}

type Request struct {
	StatusCode int `json:"status_code"`
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Meta       interface{} `json:"meta"`
}

func (api *Adapter) Request(r *Request) (*Response, error) {
	var err error
	buffer := api.pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer func() {
		if buffer != nil {
			api.pool.Put(buffer)
			buffer = nil
		}
	}()

	e := jsoniter.NewEncoder(buffer)
	err = e.Encode(r)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"request": r,
		}).Errorf("jsoniter.Marshal failure: %v", err)
		return nil, fmt.Errorf("jsoniter.Marshal failure: %v", err)
	}

	data := buffer.Bytes()
	spew.Dump(buffer.Bytes())
	req, err := http.NewRequest("GET", "http://192.168.45.22:7007/v1/receive/check", buffer)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"data": exbytes.ToString(data),
		}).Errorf("http.NewRequest failed: %v", err)
		return nil, fmt.Errorf("http.NewRequest failed: %v", err)
	}

	req.Header.Set("User-Agent", "xxx")

	httpResponse, err := http.DefaultClient.Do(req)
	if httpResponse != nil {
		defer func() {
			_, _ = io.Copy(ioutil.Discard, httpResponse.Body)
			_ = httpResponse.Body.Close()
		}()
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"url": "http://xxx.com",
		}).Errorf("query service failed %v", err)
		return nil, fmt.Errorf("query service failed %v", err)
	}

	if httpResponse.StatusCode != 200 {
		logrus.WithFields(logrus.Fields{
			"url":         "http://xxx.com",
			"status":      httpResponse.Status,
			"status_code": httpResponse.StatusCode,
		}).Errorf("invalid http status code")
		return nil, fmt.Errorf("invalid http status code")
	}

	buffer.Reset()
	_, err = io.Copy(buffer, httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("adapter io.copy failure error:%v", err)
	}

	respData := buffer.Bytes()
	logrus.WithFields(logrus.Fields{
		"response_json": exbytes.ToString(respData),
	}).Debug("response json")

	res := &Response{}
	err = jsoniter.Unmarshal(respData, res)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"data": exbytes.ToString(respData),
			"url":  "http://xxx.com",
		}).Errorf("adapter jsoniter.Unmarshal failed, error:%v", err)
		return nil, fmt.Errorf("adapter jsoniter.Unmarshal failed, error:%v", err)
	}

	api.pool.Put(buffer)
	buffer = nil

	// ...
	return res, nil
}
