package delete

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"kasmctl/get"
	"net/http"
)

func DeleteUser(url string, apikey string, secret string, notls bool, user string, force bool) {
	uri := url + "/api/public/delete_user"
	uid := get.GetUserID(url, apikey, secret, notls, user)
	if uid == "" {
		fmt.Println("User not found")
		return
	}
	//Confirm user wants to delete
	if !force {
		fmt.Print("Are you sure you want to delete user " + user + "? (y/n): ")
		var input string
		fmt.Scanln(&input)
		if input != "y" {
			fmt.Println("Aborting")
			return
		}
	}
	//Convert force to string
	f := "false"
	if force {
		f = "true"
	}
	json := []byte(`{
		"api_key":"` + apikey + `",
		"api_key_secret": "` + secret + `",
		"target_user": {
			"user_id": "` + uid + `"
		},
		"force": ` + f + `
	}`)
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(json))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	if notls {
		client = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	//If response code is 200, return a successful message.
	if resp.StatusCode == 200 {
		fmt.Println("User deleted")
	} else {
		fmt.Println("Unexpected error:")
		//Print the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))
	}
}
