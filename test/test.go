package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

//Function to test that the configuration is valid - make a GET request to the API and check that the response is 200 OK
func TestConfig(config map[string]string) (url string, key string, secret string, notls bool) {

	getusersurl := config["KASM_URL"] + "/api/public/get_users"
	json := []byte(`{"api_key":"` + config["KASM_KEY"] + `","api_key_secret":"` + config["KASM_SECRET"] + `"}`)
	req, err := http.NewRequest("POST", getusersurl, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	//Make a GET request to the API (kinda - it's a POST request, but we're just getting a list of users - /api/public/get_users)
	//Print our post request in plain text
	client := &http.Client{}
	resp, err := client.Do(req)
	//If we receive a TLS error, set the TLS flag to true, but throw a warning
	if err != nil {
		if strings.Contains(err.Error(), "x509: certificate signed by unknown authority") {
			notls = true
			fmt.Println(Yellow + "WARNING: Using an insecure connection. This is not recommended." + Reset)
		} else {
			panic(err)
		}
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//If the response is 200 OK, we're good
	if resp.StatusCode == 200 {
		fmt.Println(Green + "Successfully connected to Kasm API" + Reset)
	} else {
		fmt.Println(Red + "Failed to connect to Kasm API" + Reset)
		fmt.Println(string(body))
		os.Exit(1)
	}
	return config["KASM_URL"], config["KASM_KEY"], config["KASM_SECRET"], notls
}
