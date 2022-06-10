package push

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"kasmctl/get"
	"net/http"
)

//Weirdly, there is seemingly no option to get groups, so we only are making functionality to add or remove users from groups.
//We will get the USER ID from the username and the get.GetUserID function.
//Group ID needs to be passed from the command line.

//Add a user to group, USER ID will be taken from the username, group ID needs to be passed from the command line.
func AddUserToGroup(url string, key string, secret string, notls bool, username string, groupid string) {
	addusertogroupurl := url + "/api/public/add_user_group"

	js := []byte(`{
		"api_key": "` + key + `",
		"api_key_secret": "` + secret + `",
		"target_user": {
			"user_id": "` + get.GetUserID(url, key, secret, notls, username) + `"
		},
		"target_group": {
			"group_id": "` + groupid + `"
		}
	}`)
	req, err := http.NewRequest("POST", addusertogroupurl, bytes.NewBuffer(js))
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	//If status code is 200, then we are good.
	if resp.StatusCode == 200 {
		fmt.Println("User added to group")
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

//Remove a user from group, USER ID will be taken from the username, group ID needs to be passed from the command line.
func RemoveUserFromGroup(url string, key string, secret string, notls bool, username string, groupid string) {
	rmuserurl := url + "/api/public/remove_user_group"
	js := []byte(`{
		"api_key": "` + key + `",
		"api_key_secret": "` + secret + `",
		"target_user": {
			"user_id": "` + get.GetUserID(url, key, secret, notls, username) + `"
		},
		"target_group": {
			"group_id": "` + groupid + `"
		}
	}`)
	req, err := http.NewRequest("POST", rmuserurl, bytes.NewBuffer(js))
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	//If response code is 200, return a successful message.
	if resp.StatusCode == 200 {
		fmt.Println("User removed from group")
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
