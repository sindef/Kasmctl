package push

import (
	"fmt"
	"strings"
	"unicode"
)

//Calculate the levenshtein distance between two strings
func levenshtein(a, b string) int {
	//Calculate the length of the strings
	la := len(a)
	lb := len(b)
	//Create a matrix of the length of the strings
	d := make([][]int, la+1)
	for i := 0; i <= la; i++ {
		d[i] = make([]int, lb+1)
	}
	//Initialize the first row and column to 0
	for i := 0; i <= la; i++ {
		d[i][0] = i
	}
	for j := 0; j <= lb; j++ {
		d[0][j] = j
	}
	//Calculate the distance between the strings
	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			if a[i-1] == b[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				d[i][j] = min(d[i-1][j-1]+1, min(d[i][j-1]+1, d[i-1][j]+1))
			}
		}
	}
	//Print the distance
	fmt.Println(d[la][lb])
	//Return the distance
	return d[la][lb]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*Check if the password has at least 8 characters, 1 uppercase, 1 number, and 1 special character
This is a very basic check, but it should be good enough for our purposes, especially given we're using plaintext here.*/
func checkPassword(password string, username string) bool {

	//Check the levenshtein distance between the password and the username
	if levenshtein(password, username) < 6 {
		fmt.Println(Red+"Password for user", username, "is too similar to the username"+Reset)
		return false
	}
	//Check the length of the password
	letters := 0
	var num, upper, special bool
	//The list of characters that are not allowed in the password - thanks for this, Kasm.
	badchars := []string{"|", "-", "_", ";", ":", "`", ",", ".", "?"}
	for i, c := range password {
		switch {
		case unicode.IsNumber(c):
			num = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			for _, badchar := range badchars {
				if strings.Contains(string(c), badchar) {
					fmt.Println(Red+"Bad character:", string(c), "at position", i, "in password for user", username+Reset)
					return false
				}
			}
			special = true
		case unicode.IsSpace(c):
			fmt.Println(Red+"Space Character in password for user", username, ". Not a valid password"+Reset)
			return false
		case unicode.IsLetter(c):
			letters++

		}
	}
	if letters < 8 {
		fmt.Println(Red+"Password for user", username, "is too short"+Reset)
		return false
	}
	if num && upper && special {
		return true
	}
	//Return false if the password does not meet the requirements
	fmt.Println(Red+"\nPassword for user", username, "does not meet the password criteria for Kasm."+Yellow)
	fmt.Println("Password must contain at least 1 number, 1 uppercase and 1 special character" + Reset)
	return false
}
