package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"kasmctl/exec"
	"kasmctl/get"
	"kasmctl/push"
	"net/http"
	"testing"
	"time"
)

type TestAuth struct {
	API_KEY        string `json:"api_key"`
	API_KEY_SECRET string `json:"api_key_secret"`
}

//This package tests our functions by spinning up a local HTTP server and responding to requests with a 200 OK. This is invoked with 'go test'
func TestAPI(t *testing.T) {
	//Start up a local HTTP server
	server := http.Server{Addr: ":8080"}
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			//As this entire API uses POST requests, if we don't receive a POST request, we will return a t.Error
			if r.Method != "POST" {
				t.Error("Expected POST request")
			}
			//Ensure the POST request includes the correct key and secret in the body
			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Error("Error reading body")
			}
			fmt.Println(string(r.Method) + " " + r.URL.Path + " " + string(body))
			//Unmarshall the body into a TestAuth struct
			var testauth TestAuth
			err = json.Unmarshal(body, &testauth)
			if err != nil {
				t.Error("Error unmarshalling body")
			}
			//Ensure that the key and secret are correct - but only if we have received a request on a path that isn't /
			if r.URL.Path != "/" {
				if testauth.API_KEY != "testkey" || testauth.API_KEY_SECRET != "testsecret" {
					t.Error("Test key and secret incorrect for URL " + r.URL.Path + " expected testkey and testsecret. Got " + testauth.API_KEY + " and " + testauth.API_KEY_SECRET)
				}
			}
		})
		server.ListenAndServe()
	}()
	//Check that the server is running, maxing out at 500ms
	for i := 0; i < 5; i++ {
		//Make a quick POST request to the server
		if _, err := http.Post("http://localhost:8080/", "application/json", bytes.NewBuffer([]byte(`{"key": "testkey", "secret": "testsecret"}`))); err == nil {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
	//Make all our requests to the local HTTP server
	get.Test("http://localhost:8080", "testkey", "testsecret", false, "testtarget")
	push.Test("http://localhost:8080", "testkey", "testsecret", false, "testtarget")
	exec.Test("http://localhost:8080", "testkey", "testsecret", false, "testtarget")

}
