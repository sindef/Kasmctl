package exec

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kasmctl/get"
	"net/http"
)

//I'm only capturing the time for now, but to capture the rest of the response, create the struct here and unmarshall into this struct.
type Response struct {
	Current_Time string `json:"current_time"`
}

type Kasm struct {
	Kasm []Response `json:"kasm"`
}

func ExecCommand(url string, key string, secret string, notls bool, kasmid string, command string) {
	execurl := url + "/api/public/exec_command_kasm"
	//Get the userid matching the sessionid
	user := get.GetKasmID(url, key, secret, notls, kasmid)
	fmt.Println("Executing command: " + command + " on kasm: " + kasmid + " as user: root")
	fmt.Printf("Confirm? (y/n): ")
	var input string
	fmt.Scanln(&input)
	if input != "y" {
		fmt.Println("Aborting")
		return
	}

	js := []byte(`{
		"api_key": "` + key + `",
		"api_key_secret": "` + secret + `",
		"kasm_id": "` + kasmid + `",
		"user_id": "` + user + `",
		"exec_config":
		{
			"cmd": "` + command + `",
			"environment": {},
			"workdir": "/home/kasm-user",
			"privileged": true,
			"user": "root"
		}
	}`)
	req, err := http.NewRequest("POST", execurl, bytes.NewBuffer(js))
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
	//Print the field "current time" from the JSON response if the response is successful
	if resp.StatusCode == 200 {
		var res Response
		json.Unmarshal(body, &res)
		fmt.Println("Command executed successfully at: " + res.Current_Time)
	} else {
		fmt.Println("Command failed")
		fmt.Println(string(body))
	}
}
