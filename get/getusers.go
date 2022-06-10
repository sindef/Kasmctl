package get

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserDetails struct {
	Username     string `json:"username"`
	User_ID      string `json:"user_id"`
	Locked       bool   `json:"locked"`
	Disabled     bool   `json:"disabled"`
	Last_Session string `json:"last_session"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Phone        string `json:"phone"`
	Organization string `json:"organization"`
	Groups       struct {
		Name     string `json:"name"`
		Group_ID string `json:"group_id"`
	} `json:"groups"`
	Kasms              []string `json:"kasms"`
	Two_Factor_Enabled bool     `json:"two_factor"`
	Created            string   `json:"created"`
}

//For some reason this was not liking being created as an array of the above struct.. not sure why.
type User struct {
	User struct {
		Username     string `json:"username"`
		User_ID      string `json:"user_id"`
		Locked       bool   `json:"locked"`
		Disabled     bool   `json:"disabled"`
		Last_Session string `json:"last_session"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Phone        string `json:"phone"`
		Organization string `json:"organization"`
		Groups       struct {
			Name     string `json:"name"`
			Group_ID string `json:"group_id"`
		} `json:"groups"`
		Kasms              []string `json:"kasms"`
		Two_Factor_Enabled bool     `json:"two_factor"`
		Created            string   `json:"created"`
	} `json:"user"`
}

type Users struct {
	Users []UserDetails `json:"users"`
}

type Attr struct {
	User_Attributes struct {
		SSH_PUBLIC_KEY       string `json:"ssh_public_key"`
		Show_Tips            bool   `json:"show_tips"`
		User_ID              string `json:"user_id"`
		Toggle_Control_Panel bool   `json:"toggle_control_panel"`
		Chat_SFX             bool   `json:"chat_sfx"`
		User_Attributes_ID   string `json:"user_attributes_id"`
		Default_Image        string `json:"default_image"`
		Auto_Login_Kasm      bool   `json:"auto_login_kasm"`
	} `json:"user_attributes"`
}

//Returns a user ID from a username - this is exported to be used by the other packages
func GetUserID(url string, key string, secret string, notls bool, username string) string {
	body := GetUser(url, key, secret, notls, username)
	var user User
	json.Unmarshal(body, &user)
	return user.User.User_ID
}

//Returns a JSON response from the API - gets a single user and returns the response as a []byte which can be unmarshalled into a struct
func GetUser(url string, key string, secret string, notls bool, username string) []byte {
	getuserurl := url + "/api/public/get_user"

	js := []byte(`{
		"api_key": "` + key + `",
		"api_key_secret": "` + secret + `",
		"target_user": {
			"username": "` + username + `"
		}
	}`)
	req, err := http.NewRequest("POST", getuserurl, bytes.NewBuffer(js))
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
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func SingleUser(url string, key string, secret string, notls bool, username string) {
	body := GetUser(url, key, secret, notls, username)

	//Unmarshall the JSON into a struct - we can then print the fields we want
	var user User
	json.Unmarshal(body, &user)
	//Print all of the user's data from the response (this is what we unmarshalled into our struct)
	fmt.Println("Username: " + user.User.Username)
	fmt.Println("User ID: " + user.User.User_ID)
	fmt.Println("Locked: " + fmt.Sprint(user.User.Locked))
	fmt.Println("Disabled: " + fmt.Sprint(user.User.Disabled))
	fmt.Println("Last Session: " + user.User.Last_Session)
	fmt.Println("First Name: " + user.User.FirstName)
	fmt.Println("Last Name: " + user.User.LastName)
	fmt.Println("Phone: " + user.User.Phone)
	fmt.Println("Organization: " + user.User.Organization)
	fmt.Println("Group: " + user.User.Groups.Name)
	fmt.Println("Group ID: " + user.User.Groups.Group_ID)
	fmt.Println("Kasms: " + fmt.Sprint(user.User.Kasms))
	fmt.Println("Two Factor Enabled: " + fmt.Sprint(user.User.Two_Factor_Enabled))
	fmt.Println("Created: " + user.User.Created)

}

//Returns all of the users in the system
func AllUsers(url string, key string, secret string, notls bool, verbose bool) {
	getuserurl := url + "/api/public/get_users"

	js := []byte(`{"api_key":"` + key + `","api_key_secret":"` + secret + `"}`)
	req, err := http.NewRequest("POST", getuserurl, bytes.NewBuffer(js))
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
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//Unmarshall the JSON.
	var users Users
	json.Unmarshal(body, &users)
	//Print all of the user's data - I used println to make it easier to read. Printf would be better for formatting, but I found the code to be a bit harder to read.
	for _, user := range users.Users {
		if verbose {
			fmt.Println("Username: " + user.Username)
			fmt.Println("User ID: " + user.User_ID)
			fmt.Println("Locked: " + fmt.Sprint(user.Locked))
			fmt.Println("Disabled: " + fmt.Sprint(user.Disabled))
			fmt.Println("Last Session: " + user.Last_Session)
			fmt.Println("First Name: " + user.FirstName)
			fmt.Println("Last Name: " + user.LastName)
			fmt.Println("Phone: " + user.Phone)
			fmt.Println("Organization: " + user.Organization)
			fmt.Println("Group: " + user.Groups.Name)
			fmt.Println("Group ID: " + user.Groups.Group_ID)
			fmt.Println("Kasms: " + fmt.Sprint(user.Kasms))
			fmt.Println("Two Factor Enabled: " + fmt.Sprint(user.Two_Factor_Enabled))
			fmt.Println("Created: " + user.Created)
			fmt.Println("")
		} else {
			fmt.Println("Username: " + user.Username)
			fmt.Println("User ID: " + user.User_ID)
			fmt.Println("")
		}
	}

}

func SingleUserAttr(url string, key string, secret string, notls bool, username string) {
	user := GetUser(url, key, secret, notls, username)
	//We have returned a JSON object, but we only need the user_ID field for the next request so we'll unmarshall it into a struct of User then pull out the user_ID field
	var userd User
	json.Unmarshal(user, &userd)
	userid := userd.User.User_ID
	//Now we can make the request to get the user's attributes
	js := []byte(`{
		"api_key": "` + key + `",
		"api_key_secret": "` + secret + `",
		"target_user": {
			"user_id": "` + userid + `"
		}
	}`)
	attrurl := url + "/api/public/get_attributes"
	req, err := http.NewRequest("POST", attrurl, bytes.NewBuffer(js))
	if err != nil {
		panic(err)
	}
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

	//Unmarshall the JSON into a struct - we can then print the fields we want - Structs are found at the top of this file
	var userattr Attr
	json.Unmarshal(body, &userattr)

	fmt.Println("Username: " + username)
	fmt.Println("User_ID: " + userattr.User_Attributes.User_ID)
	fmt.Println("User Attributes ID: " + userattr.User_Attributes.User_Attributes_ID)
	fmt.Println("SSH Public Key: " + userattr.User_Attributes.SSH_PUBLIC_KEY)
	fmt.Println("Show Tips Enabled: " + fmt.Sprint(userattr.User_Attributes.Show_Tips))
	fmt.Println("Toggle Control Panel Enabled: " + fmt.Sprint(userattr.User_Attributes.Toggle_Control_Panel))
	fmt.Println("Chat_SFX_Enabled: " + fmt.Sprint(userattr.User_Attributes.Chat_SFX))
	fmt.Println("Default Image: " + userattr.User_Attributes.Default_Image)
	fmt.Println("Auto-Login Enabled: " + fmt.Sprint(userattr.User_Attributes.Auto_Login_Kasm))
	fmt.Println("")
}
