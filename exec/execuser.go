package exec

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"kasmctl/get"
	"net/http"
)

func LogoutUser(url string, key string, secret string, notls bool, username string) {
	userid := get.GetUserID(url, key, secret, notls, username)
	uri := url + "/api/public/logout_user"
	//Json has to be a byte in the form of:
	js := []byte(`{
		"api_key": "` + key + `",
		"api_key_secret": "` + secret + `",
		"target_user": {
			"user_id": "` + userid + `"
		}
	}`)
	//Make the post request to logout the user
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(js))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	if notls {
		client = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//We don't actually get a response if we logout, so we just print the response ONLY if we get an error
	if resp.StatusCode != 200 {
		fmt.Println(string(body))
	} else {
		fmt.Println("User " + username + " has been logged out")
	}
}
