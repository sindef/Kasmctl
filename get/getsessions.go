package get

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Kasm struct {
	Expiration_Date string `json:"expiration_date"`
	Container_IP    string `json:"container_ip"`
	Server          struct {
		Port      int    `json:"port"`
		Hostname  string `json:"hostname"`
		Zone_Name string `json:"zone_name"`
		Provider  string `json:"provider"`
	} `json:"server"`
	User struct {
		Username string `json:"username"`
	} `json:"user"`
	Start_Date      string  `json:"start_date"`
	PointofPresence string  `json:"point_of_presence"`
	Token           string  `json:"token"`
	Image_id        string  `json:"image_id"`
	View_only_token string  `json:"view_only_token"`
	Cores           float64 `json:"cores"`
	Hostname        string  `json:"hostname"`
	Kasm_id         string  `json:"kasm_id"`
	Image           struct {
		Image_id      string `json:"image_id"`
		Name          string `json:"name"`
		Image_src     string `json:"image_src"`
		Friendly_name string `json:"friendly_name"`
	}
	Is_persistent_profile string `json:"is_persistent_profile"`
	Memory                int    `json:"memory"`
	Operational_status    string `json:"operational_status"`
	Container_id          string `json:"container_id"`
	Port                  int    `json:"port"`
	Keepalive_date        string `json:"keepalive_date"`
	User_id               string `json:"user_id"`
	Share_id              string `json:"share_id"`
	Host                  string `json:"host"`
	Server_id             string `json:"server_id"`
}

type Kasms struct {
	Kasms []Kasm `json:"kasms"`
}

func GetKasms(url string, key string, secret string, notls bool) []byte {
	getuserurl := url + "/api/public/get_kasms"

	js := []byte(`{"api_key":"` + key + `","api_key_secret":"` + secret + `"}`)
	req, err := http.NewRequest("POST", getuserurl, bytes.NewBuffer(js))
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

func GetKasmID(url string, key string, secret string, notls bool, sessionid string) string {
	body := GetKasms(url, key, secret, notls)

	//Unmarshall the JSON into a struct - we can then print the fields we want
	var kasms Kasms
	json.Unmarshal(body, &kasms)
	//Loop through each Kasm and if the user.Username matches the username we want, print the fields
	for _, kasm := range kasms.Kasms {
		if kasm.Kasm_id == sessionid {
			fmt.Println("Session ID matches user: " + kasm.User.Username)
			return kasm.User_id
		}
	}
	return ""
}

func getKasmUser(url string, key string, secret string, notls bool, username string) {
	body := GetKasms(url, key, secret, notls)

	//Unmarshall the JSON into a struct - we can then print the fields we want
	var kasms Kasms
	json.Unmarshal(body, &kasms)
	//Loop through each Kasm and if the user.Username matches the username we want, print the fields
	for _, kasm := range kasms.Kasms {
		if kasm.User.Username == username {
			fmt.Println("Kasm ID: " + kasm.Kasm_id)
			fmt.Println("Container ID: " + kasm.Container_id)
			fmt.Println("Server ID: " + kasm.Server_id)
			fmt.Println("User ID: " + kasm.User_id)
			fmt.Println("Share ID: " + kasm.Share_id)
			fmt.Println("Host: " + kasm.Host)
			fmt.Println("Port:", kasm.Port)
			fmt.Println("Memory:", kasm.Memory)
			fmt.Println("Cores:", kasm.Cores)
			fmt.Println("Operational Status: " + kasm.Operational_status)
			fmt.Println("Keepalive Date: " + kasm.Keepalive_date)
			fmt.Println("Start Date: " + kasm.Start_Date)
			fmt.Println("Expiration Date: " + kasm.Expiration_Date)
			fmt.Println("Point of Presence: " + kasm.PointofPresence)
			fmt.Println("Image ID: " + kasm.Image_id)
			fmt.Println("Image Name: " + kasm.Image.Name)
			fmt.Println("Image Friendly Name: " + kasm.Image.Friendly_name)
			fmt.Println("Image Source: " + kasm.Image.Image_src)
			fmt.Println("Hostname: " + kasm.Hostname)
			fmt.Println("View Only Token:", kasm.View_only_token)
			fmt.Println("")
		}
	}

}

//Prints all kasm IDs, memory, cores and associated user on the server
func GetAllKasms(url string, key string, secret string, notls bool) {
	body := GetKasms(url, key, secret, notls)
	var kasms Kasms
	json.Unmarshal(body, &kasms)
	for _, kasm := range kasms.Kasms {
		fmt.Println("Kasm ID: " + kasm.Kasm_id)
		fmt.Println("Share ID: " + kasm.Share_id)
		fmt.Println("Memory:", kasm.Memory)
		fmt.Println("Cores:", kasm.Cores)
		fmt.Println("Image Friendly Name: " + kasm.Image.Friendly_name)
		fmt.Println("View Only Token: " + kasm.View_only_token)
		fmt.Println("User Name: " + kasm.User.Username)
		fmt.Println("User ID: " + kasm.User_id)
		fmt.Println("")
	}
}
