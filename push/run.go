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
	default:
		fmt.Println("Invalid target")
		admin.Help()
	}
}
