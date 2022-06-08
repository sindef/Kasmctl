package admin

import (
	"fmt"
	"os"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

//Reads the environmental variables for the url, secret and key and store this as a configmap
func Getenv() map[string]string {
	config := make(map[string]string)
	//Read the config file - the default is $HOME/.kasmctl/config
	if _, err := os.Stat(os.Getenv("HOME") + "/.kasmctl/config"); err != nil {
		fmt.Println(Yellow + "Warning: Config file not found at " + os.Getenv("HOME") + "/.kasmctl/config, using environmental variables" + Reset)
		config["KASM_URL"] = os.Getenv("KASM_URL")
		config["KASM_SECRET"] = os.Getenv("KASM_SECRET")
		config["KASM_KEY"] = os.Getenv("KASM_KEY")
		if config["KASM_URL"] == "" || config["KASM_SECRET"] == "" || config["KASM_KEY"] == "" {
			panic("Missing environmental variables or config file for KASM_URL, KASM_SECRET or KASM_KEY")
		}
		//Write the config file
		WriteConfig(config)
	} else {
		config = readConfig()
	}
	if !strings.HasPrefix(config["KASM_URL"], "https://") {
		panic("KASM_URL must start with https:// - http is not supported")
	}
	return config
}

func Help() {
	var helpMessage string = `
Kasmctl is designed to interact with our Kasm API, it only includes basic operations so far. 

Usage: kasmctl [OPTIONS] [OPERATION] [TARGET] 

	Operations:
	get 		Display one or many resources
	push 		Push a new resource to the API, such as a new user or a new group
	delete 		Delete a resource from the Kasm Server - this is not reversible
	exec 		Execute a command on a Kasm session - images are identified by their ID found with 'kasmctl get sessions'

	Targets:
	users		Create, list, update, or delete users.
	groups		Create, list, update, or delete groups.
	sessions	Create, list, update, or terminate sessions. Sessions can also have commands executed on them.

	Options:
	--help		Display this help message
	--version	Display the version of kasmctl
	--verbose	Display verbose output (for 'get users')
`
	fmt.Println(helpMessage)
	os.Exit(0)
}
