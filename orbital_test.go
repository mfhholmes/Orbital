package main

import (
	"testing"
)

import (
	"net/url"
)
import (
	"net/http"
)
import (
	"net/http/httptest"
)

func TestPopulationGet(t *testing.T) {
	handler := buildHandler()
	server := httptest.NewServer(handler)
	_, err := http.Get(server.URL + "/population")
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
}
func TestPopulationSet(t *testing.T) {
	handler := buildHandler()
	server := httptest.NewServer(handler)
	// set the data
	startPop := "255"
	growthRate := "10"
	popVal := url.Values{}
	popVal.Set("startPopulation", startPop)
	popVal.Add("growthRate", growthRate)
	client := new(http.Client)
	_, err := client.PostForm(server.URL+"/population", popVal)
	if err != nil {
		t.Log("problem posting values")
		t.Fatal(err)
	}
	getResponse, err := http.Get(server.URL + "/population")
	if err != nil {
		t.Log("problem getting values")
		t.Fatal(err)
	}
	t.Log(getResponse)
	defer server.Close()
}
