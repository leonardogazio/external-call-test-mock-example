package fooserviceapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type FooClientIface interface {
	GetFoo() (out FooBar, err error)
}

var Client FooClientIface

// New ...
func New(url string) FooClientIface {
	return &fooClient{
		URL: url,
	}
}

// GetFoo ...
func (fc *fooClient) GetFoo() (out FooBar, err error) {
	res, err := http.Get(fc.URL)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	return out, json.Unmarshal(body, &out)
}
