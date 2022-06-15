package delete

import (
	"fmt"
	"kasmctl/admin"
	"kasmctl/test"
)

func Run(target []string) {
	if len(target) < 1 {
		fmt.Println("Please specify a target")
		return
	}
	config := admin.Getenv()
	url, key, secret, notls := test.TestConfig(config)
	switch target[0] {
	case "user", "users":
		if len(target) < 2 {
			fmt.Println("Please specify a user")
			return
		}
		if len(target) > 3 {
			fmt.Println("Too many arguments")
			return
		}
		if len(target) == 2 {
			DeleteUser(url, key, secret, notls, target[1], false)
		}
		if len(target) == 3 {
			if target[2] == "--force" {
				DeleteUser(url, key, secret, notls, target[1], true)
			} else {
				fmt.Println("Unknown argument: " + target[2])
				fmt.Println("Usage: kasmctl delete user <username> [--force]")
			}
		}
	case "sessions", "session", "kasm", "kasms":
		if len(target) < 2 {
			fmt.Println("Please specify a user")
			fmt.Println("Usage: kasmctl delete sessions <user> [<kasmid>]")
			return
		}
		if len(target) > 3 {
			fmt.Println("Too many arguments")
			fmt.Println("Usage: kasmctl delete sessions <user> [<kasmid>]")
			return
		}
		if len(target) == 2 {
			DestroyKasm(url, key, secret, notls, target[1], "")
		}
		if len(target) == 3 {
			DestroyKasm(url, key, secret, notls, target[1], target[2])
		}
	default:
		fmt.Println("Unknown target: " + target[0])
	}

}
func Test(url string, key string, secret string, notls bool, target string) {
	DestroyKasm(url, key, secret, notls, target, "")
	DeleteUser(url, key, secret, notls, target, false)
}
