package conf

import (
	"os"
	"runtime"
	"testing"
)

//Write a quick test config file to the /tmp/ or (if on Windows), the %localappdata%\Temp folder with some test vars
func TestConfig(t *testing.T) {
	url := "http://localhost:8080"
	key := "testkey"
	secret := "testsecret"
	dir := ""
	//Check if we are on a Unix fs or Windows
	if runtime.GOOS == "windows" {
		dir = "%localappdata%\\Temp\\KasmctlTest"
	} else {
		dir = "/tmp/kasmctltest"
	}
	//Create the directory if it doesn't exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}
	//Write the config file
	file, err := os.Create(dir + "/config")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString("url=" + url + "\n")
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString("key=" + key + "\n")
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString("secret=" + secret + "\n")
	if err != nil {
		panic(err)
	}
	configMap := readConfig(dir + "/config")
	if configMap["url"] != url {
		t.Error("Expected url to be " + url + " but got " + configMap["url"])
	}
	if configMap["key"] != key {
		t.Error("Expected key to be " + key + " but got " + configMap["key"])
	}
	if configMap["secret"] != secret {
		t.Error("Expected secret to be " + secret + " but got " + configMap["secret"])
	}

}
