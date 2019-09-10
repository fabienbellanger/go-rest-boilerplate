package user

import (
	"testing"

	"github.com/fabienbellanger/go-rest-boilerplate/models"
)

// TestCheckLogin
func TestCheckLogin(t *testing.T) {
	var username, password string
	var user, userWanted models.User

	username = ""
	password = ""
	user, _ = CheckLogin(username, password)
	if user != userWanted {
		t.Errorf("CheckLogin(%s, %s) - got: %v, want: %v", username, password, user, userWanted)
	}
}
