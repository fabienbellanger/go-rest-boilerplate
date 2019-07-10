package models

import (
	"testing"
)

// TestGetFullname tests GetFullname function
func TestGetFullname(t *testing.T) {
	user := User{Lastname: "Lastname", Firstname: "Firstname"}
	fullname := user.GetFullname()
	fullnameWanted := "Firstname Lastname"
	if fullname != fullnameWanted {
		t.Errorf("GetFullname - got: %s, want: %s.", fullname, fullnameWanted)
	}

	user = User{Lastname: "", Firstname: "Firstname"}
	fullname = user.GetFullname()
	fullnameWanted = "Firstname"
	if fullname != fullnameWanted {
		t.Errorf("GetFullname - got: %s, want: %s.", fullname, fullnameWanted)
	}

	user = User{Lastname: "Lastname", Firstname: ""}
	fullname = user.GetFullname()
	fullnameWanted = "Lastname"
	if fullname != fullnameWanted {
		t.Errorf("GetFullname - got: %s, want: %s.", fullname, fullnameWanted)
	}

	user = User{Lastname: "", Firstname: ""}
	fullname = user.GetFullname()
	fullnameWanted = ""
	if fullname != fullnameWanted {
		t.Errorf("GetFullname - got: %s, want: %s.", fullname, fullnameWanted)
	}
}

// TestCheckLogin
func TestCheckLogin(t *testing.T) {
	var username, password string
	var user, userWanted User

	username = ""
	password = ""
	user, _ = CheckLogin(username, password)
	if user != userWanted {
		t.Errorf("CheckLogin(%s, %s) - got: %v, want: %v", username, password, user, userWanted)
	}
}
