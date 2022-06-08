package get

import (
	"fmt"
	"kasmctl/admin"
	"kasmctl/test"
)

//Entry point for GET functions
func Run(target []string) {
	config := admin.Getenv()
	url, key, secret, notls := test.TestConfig(config)
	if len(target) > 0 {
		switch target[0] {
		case "user":
			//Usage: kasmctl get user <username> or kasmctl get user <username> --attributes
			if len(target) > 2 {
				if target[2] == "attributes" {
					singleUserAttr(url, key, secret, notls, target[1])
				} else {
					fmt.Println("Invalid argument")
					fmt.Println("Usage: kasmctl get user <username> [attributes]")
				}
			} else if len(target) < 1 {
				fmt.Println("Please specify a user")
				fmt.Println("kasmctl get user <user>")
			} else {
				singleUser(url, key, secret, notls, target[1])
			}

		case "users":
			//Usage: kasmctl get users or kasmctl get users --verbose
			if len(target) > 1 {
				if target[1] == "--verbose" {
					allUsers(url, key, secret, notls, true)
				}
			} else {
				allUsers(url, key, secret, notls, false)
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
					getKasmUser(url, key, secret, notls, target[1])
				}
			} else {
				GetAllKasms(url, key, secret, notls)
			}

		case "images":
			fmt.Println("Not implemented yet")
		default:
			fmt.Println("Invalid target")
			admin.Help()
		}
	} else {
		fmt.Println("No target provided - Please provide a target")
		admin.Help()
	}
}
