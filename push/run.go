package push

import (
	"fmt"
	"kasmctl/admin"
	"kasmctl/test"
	"os"
)

//Entry point for PUSH functions
func Run(target []string) {
	config := admin.Getenv()
	url, key, secret, notls := test.TestConfig(config)
	switch target[0] {
	case "user":
		singleUser(url, key, secret, notls)
	case "users":
		//Check file exists
		if _, err := os.Stat(target[1]); err == nil {
			bulkuser(target[1], url, key, secret, notls)
		} else {
			fmt.Println("Please provide a valid file")
		}
	case "group":
		//Check the length of the target slice
		if len(target) > 3 {
			if target[1] == "add" {
				AddUserToGroup(url, key, secret, notls, target[2], target[3])
			} else if target[1] == "remove" {
				RemoveUserFromGroup(url, key, secret, notls, target[2], target[3])
			} else {
				fmt.Println("Invalid command")
				fmt.Println("kasmctl push group [add/remove] <username> <groupid>")
			}
		} else {
			fmt.Println("Invalid command")
			fmt.Println("kasmctl push group [add/remove] <username> <groupid>")
		}
	default:
		fmt.Println("Invalid target")
		admin.Help()
	}
}
func Test(url string, key string, secret string, notls bool, target string) {
	//Have to redo these so I can run the tests
	// singleUser(url, key, secret, notls)
	// bulkuser(target, url, key, secret, notls)
	AddUserToGroup(url, key, secret, notls, "testuser", "1")
	RemoveUserFromGroup(url, key, secret, notls, "testuser", "1")

}
