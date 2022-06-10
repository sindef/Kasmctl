package push

//This is very dirty - this was the original quick script I made (albeit somewhat butchered) to add bulk users. I'll redo this at some point

import (
	"bytes"
	"crypto/tls"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"unicode"
)

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

//Take the CSV file, and return this as a map with the first row as the keys
func parseCSV(file string) (map[string][]string, error) {
	csvFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	//Create the map from the csvdata
	csvData := make(map[string][]string)
	for i, row := range rawCSVdata {
		if i == 0 {
			for _, key := range row {
				csvData[key] = []string{}
			}
		} else {
			for i, key := range row {
				csvData[rawCSVdata[0][i]] = append(csvData[rawCSVdata[0][i]], key)
			}
		}
	}
	return csvData, nil
}

//Adds a user via a POST request to the Kasm API
func addUser(csvData map[string]string, url string, apikey string, secret string, notls bool) string {
	//Verify the password is valid
	if !checkPassword(csvData["password"], csvData["username"]) {
		os.Exit(1)
	}
	//Join the url
	url = url + "/api/public/create_user"
	var jreq = createuserJson(csvData, apikey, secret)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jreq))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	//Set http default transport to skip certificate verification
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
	//If the response returns a user_id and the username matches the one we sent, we're good
	if strings.Contains(string(body), "user_id") && strings.Contains(string(body), csvData["username"]) {
		return Green + "Successfully created user: " + csvData["username"] + Reset
	} else {
		return string(body)
	}
}

/*Check if the password has at least 8 characters, 1 uppercase, 1 number, and 1 special character
This is a very basic check, but it should be good enough for our purposes, especially given we're using plaintext here.*/
func checkPassword(password string, username string) bool {
	letters := 0
	var num, upper, special bool
	//For some reason, punctuation is not accepted - so we need to reject it (will look into this if time allows)
	badchars := []string{"!", "(", ")", "-", "_", "=", "+", "[", "]", "{", "}", ";", ":", "'", ",", ".", "?"}
	for i, c := range password {
		switch {
		case unicode.IsNumber(c):
			num = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			for _, badchar := range badchars {
				if strings.Contains(string(c), badchar) {
					fmt.Println(Red+"Bad character:", string(c), "at position", i, "in password for user", username+Reset)
					return false
				}
			}
			special = true
		case unicode.IsSpace(c):
			fmt.Println(Red+"Space Character in password for user", username, ". Not a valid password"+Reset)
			return false
		case unicode.IsLetter(c):
			letters++

		}
	}
	if letters < 8 {
		fmt.Println(Red+"Password for user", username, "is too short"+Reset)
		return false
	}
	if num && upper && special {
		return true
	}
	//Return which variable was not true
	fmt.Println(Red+"\nPassword for user", username, "does not meet the password criteria for Kasm."+Yellow)
	fmt.Println("Password must contain at least 1 number, 1 uppercase and 1 special character" + Reset)
	return false
}

func createuserJson(csvData map[string]string, apikey string, secret string) []byte {
	json := `{
		"api_key":"` + apikey + `",
		"api_key_secret": "` + secret + `",
		"target_user": {
			"username" : "` + csvData["username"] + `",
			"first_name" : "` + csvData["first_name"] + `",
			"last_name" : "` + csvData["last_name"] + `",
			"locked": false,
			"disabled": false,
			"organization": "` + csvData["organization"] + `",
			"phone": "` + csvData["phone"] + `",
			"password": "` + csvData["password"] + `"
		}
	}`
	return []byte(json)
}

func bulkuser(file string, url string, apikey string, secret string, notls bool) {
	csvData, err := parseCSV(file)
	if err != nil {
		panic(err)
	}
	//For the length of the map, create a new map with the keys and values of each row - call addUser for each row
	for i := 0; i < len(csvData["username"]); i++ {
		data := make(map[string]string)
		for key, value := range csvData {
			data[key] = value[i]
		}
		fmt.Println(addUser(data, url, apikey, secret, notls))
	}
}

func singleUser(url string, apikey string, secret string, notls bool) {
	//Get the data from the user
	fmt.Printf("Enter the username: ")
	var username string
	fmt.Scanln(&username)
	fmt.Printf("Enter the first name: ")
	var firstname string
	fmt.Scanln(&firstname)
	fmt.Printf("Enter the last name: ")
	var lastname string
	fmt.Scanln(&lastname)
	fmt.Printf("Enter the organization: ")
	var organization string
	fmt.Scanln(&organization)
	fmt.Printf("Enter the phone number: ")
	var phone string
	fmt.Scanln(&phone)
	fmt.Printf("Enter the initial password: ")
	var password string
	fmt.Scanln(&password)
	//Create the map from the user input
	csvData := make(map[string]string)
	csvData["username"] = username
	csvData["first_name"] = firstname
	csvData["last_name"] = lastname
	csvData["organization"] = organization
	csvData["phone"] = phone
	csvData["password"] = password
	//Add the user
	fmt.Println(addUser(csvData, url, apikey, secret, notls))
}
