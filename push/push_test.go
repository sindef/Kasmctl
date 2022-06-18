package push

import "testing"

func TestPush(t *testing.T) {
	testPassword := "testpassword"
	testUsername := "username"
	basicCheck := checkPassword(testPassword, testUsername)
	if basicCheck {
		t.Error("Expected basic password to be false but got true")
	}
	testPassword = "t#3sTpW0rD$$as91sa"
	complexCheck := checkPassword(testPassword, testUsername)
	if !complexCheck {
		t.Error("Expected complex password to be true but got false")
	}
	//Test the similarity check
	testPassword = "username"
	similarityCheck := checkPassword(testPassword, testUsername)
	if similarityCheck {
		t.Error("Expected similarity check to be false but got true")
	}

}
