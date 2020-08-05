package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoot(t *testing.T) {
	testserver := httptest.NewServer(http.HandlerFunc(IndexHandler))
	defer testserver.Close()

	res, err := http.Get(testserver.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error("a response code is not 200")
	}

	if string(body) != "success" {
		t.Error("a response is not success")
	}
}
