package main

import (
	"fmt"
	"kasmctl/admin"
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
		admin.Help()
	case "version", "--version", "-v":
		//Print the version
		fmt.Println("kasmctl v0.2")
	default:
		fmt.Println("Invalid operation" + os.Args[1])

		os.Exit(1)
	}

}
