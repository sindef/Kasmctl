package exec

import (
	"fmt"
	"kasmctl/admin"
	"kasmctl/test"
)

//kasmctl exec user logout <username>
//kasmctl exec command <sessionid> <command>

//This is the entry point for EXEC functions
func Run(target []string) {
	config := admin.Getenv()
	url, key, secret, notls := test.TestConfig(config)
	switch target[0] {
	case "user":
		if len(target) > 1 {
			if target[1] == "logout" {
				if len(target) > 2 {
					LogoutUser(url, key, secret, notls, target[2])
				} else {
					fmt.Println("Please specify a user")
					fmt.Println("kasmctl exec user logout <username>")
				}
			}
		} else {
			fmt.Println("Invalid argument")
			fmt.Println("Usage: kasmctl exec user logout <username>")
		}
	case "command":
		if len(target) > 2 {
			ExecCommand(url, key, secret, notls, target[2], target[1])
		} else {
			fmt.Println("Please specify a sessionid and command")
			fmt.Println("kasmctl exec command <sessionid> <command>")
		}
	default:
		fmt.Println("Invalid target")
		admin.Help()
	}
}
func Test(url string, key string, secret string, notls bool, target string) {
	// ExecCommand(url, key, secret, notls, "1", "echo test")
	// LogoutUser(url, key, secret, notls, "testuser")
}
