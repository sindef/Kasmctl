package conf

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Read config file - returns a map with the config values
func readConfig(configFile string) map[string]string {
	config := make(map[string]string)
	//Read the config file - either the environmental variable KASM_CONFIG or the default is $HOME/.kasmctl/config
	if os.Getenv("KASM_CONFIG") != "" {
		configFile = os.Getenv("KASM_CONFIG")
	}
	if configFile == "" {
		configFile = os.Getenv("HOME") + "/.kasmctl/config"
	}
	file, err := os.Open(configFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//As this is a .conf file in plaintext, we just need to read the lines and split them by the = sign. Any line that doesn't have a = sign or starts with a # is ignored.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") && !strings.HasPrefix(line, "#") {
			split := strings.Split(line, "=")
			//Read only the first value
			config[split[0]] = split[1]
		}
	}
	return config
}

func WriteConfig(config map[string]string) {
	//Write the config file - $HOME/.kasmctl/config
	configFile := os.Getenv("HOME") + "/.kasmctl/config"
	fmt.Println("Environmental variables are all set, writing new config file")
	//Create the directory if it doesn't exist
	if _, err := os.Stat(os.Getenv("HOME") + "/.kasmctl"); os.IsNotExist(err) {
		os.Mkdir(os.Getenv("HOME")+"/.kasmctl", 0755)
	}
	cfile, err := os.Create(configFile)
	if err != nil {
		panic(err)
	}
	defer cfile.Close()
	file, err := os.OpenFile(configFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//Write the config file
	for key, value := range config {
		_, err := file.WriteString(key + "=" + value + "\n")
		if err != nil {
			panic(err)
		}
	}
}
