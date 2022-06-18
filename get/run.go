package get

import (
	"fmt"
	"kasmctl/conf"
	"kasmctl/test"
)

//Entry point for GET functions
func Run(target []string) {
	config := conf.Getenv()
	url, key, secret, notls := test.TestConfig(config)
	if len(target) > 0 {
		switch target[0] {
		case "user":
			//Usage: kasmctl get user <username> or kasmctl get user <username> --attributes
			if len(target) > 2 {
				if target[2] == "attributes" {
					SingleUserAttr(url, key, secret, notls, target[1])
				} else {
					fmt.Println("Invalid argument")
					fmt.Println("Usage: kasmctl get user <username> [attributes]")
				}
			} else if len(target) < 1 {
				fmt.Println("Please specify a user")
				fmt.Println("kasmctl get user <user>")
			} else {
				SingleUser(url, key, secret, notls, target[1])
			}

		case "users":
			//Usage: kasmctl get users or kasmctl get users --verbose
			if len(target) > 1 {
				if target[1] == "--verbose" {
					AllUsers(url, key, secret, notls, true)
				}
			} else {
				AllUsers(url, key, secret, notls, false)
			}
		case "groups":
			fmt.Println("Not implemented yet")
		case "sessions":
			//Usage: kasmctl get sessions <username> or kasmctl get sessions
			if len(target) > 1 {
				if len(target) > 2 {
					fmt.Println("Invalid argument")
					fmt.Println("Usage: kasmctl get sessions <username>")
				} else {
					GetKasmUser(url, key, secret, notls, target[1])
				}
			} else {
				GetAllKasms(url, key, secret, notls)
			}

		case "images":
			//Usage: kasmctl get images
			if len(target) == 2 && target[1] == "--verbose" {
				GetImages(url, key, secret, notls, true)
			} else if len(target) == 1 {
				GetImages(url, key, secret, notls, false)
			} else {
				fmt.Println("Invalid argument")
				fmt.Println("Usage: kasmctl get images [--verbose]")
			}
		default:
			fmt.Println("Invalid target")
			conf.Help()
		}
	} else {
		fmt.Println("No target provided - Please provide a target")
		conf.Help()
	}
}

func Test(url string, key string, secret string, notls bool, target string) {
	GetAllKasms(url, key, secret, notls)
	GetKasms(url, key, secret, notls)
	GetKasmUser(url, key, secret, notls, target)
	GetKasmID(url, key, secret, notls, target)
	GetUser(url, key, secret, notls, target)
	GetUserID(url, key, secret, notls, target)
	AllUsers(url, key, secret, notls, false)
	GetImages(url, key, secret, notls, false)
}
