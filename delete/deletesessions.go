package delete

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"kasmctl/get"
	"net/http"
)

func DestroyKasm(url string, apikey string, secret string, notls bool, user string, kasmid string) {
	uri := url + "/api/public/destroy_kasm"
	uid := get.GetUserID(url, apikey, secret, notls, user)
	kid := get.GetKasm(url, apikey, secret, notls, user)
	//If kasmid is not specified, check the user kasms, and ask the user to choose one. If there are no kasms, return.
	if kasmid == "" {
		if len(kid) == 0 {
			fmt.Println("User has no kasms")
			return
		}
		fmt.Println("Choose a kasm to destroy:")
		for i, v := range kid {
			fmt.Println(i, "-", v)
		}
		var input int
		fmt.Scanln(&input)
		//If input is out of range, return.
		if input < 0 || input >= len(kid) {
			fmt.Println("Invalid input")
			return
		}
		kasmid = kid[input]
	}
	json := []byte(`{
		"api_key":"` + apikey + `",
		"api_key_secret": "` + secret + `",
		"kasm_id": "` + kasmid + `",
		"user_id": "` + uid + `"
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
		fmt.Println("Kasm destroyed successfully")
	} else {
		fmt.Println("Unexpected error")
	}
}
