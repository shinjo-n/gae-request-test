package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
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

func TestJson(t *testing.T) {
	testserver := httptest.NewServer(http.HandlerFunc(jsonHandlerFunc))
	defer testserver.Close()

	json := bytes.NewBuffer([]byte(`{"token":"test","device":"test2"}`))
	res, err := http.Post(testserver.URL, "application/json", json)
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
