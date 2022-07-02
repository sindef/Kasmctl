package main

import (
	"fmt"
	"kasmctl/conf"
	"kasmctl/delete"
	"kasmctl/exec"
	"kasmctl/get"
	"kasmctl/push"
	"os"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

func main() {
	//If there are no arguments, print the welcome message:
	if len(os.Args) == 1 {
		fmt.Printf(Green + "\nWelcome to kasmctl!\n" + Reset +
			"kasmctl is a command line interface for interacting with the Kasm API. It controls a Kasm instance using API calls\n" +
			"To get started, create a configuration file at ~/.kasmctl/config, or set the following environmental variables:\n\n" +
			"KASM_SECRET: Your API secret\n" +
			"KASM_KEY: Your API key\n" +
			"KASM_URL: The URL of the Kasm instance you want to control\n" +
			"\n" +
			"For more information, see the README.md file in the kasmctl repository on gitlab or reach out to me on Teams.\n" +
			"You can also run `kasmctl --help` to see a list of commands and targets.\n")
		return
	}
	//Parse the operation line arguments
	target := os.Args[2:]
	switch os.Args[1] {
	case "get", "--get", "-g":
		get.Run(target)
	case "push", "--push", "-p":
		push.Run(target)
	case "delete", "del", "rm", "--delete", "--del", "--rm", "-d":
		delete.Run(target)
	case "exec", "--exec", "-e", "execute":
		exec.Run(target)
	case "help", "--help", "-h":
		//Print the help message
		conf.Help()
	case "version", "--version", "-v":
		//Print the version
		fmt.Println("kasmctl v0.4")
	default:
		fmt.Println("Invalid operation" + os.Args[1])

		os.Exit(1)
	}

}
