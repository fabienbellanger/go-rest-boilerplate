package lib

import (
	"testing"
)

// TestIsDatebaseConfigCorrect : Test si la configuration de la base de données est correcte
func TestIsDatebaseConfigCorrect(t *testing.T) {
	Config.Database.Driver = ""
	Config.Database.Name = ""
	Config.Database.Host = ""
	isCorrect := IsDatabaseConfigCorrect()

	if isCorrect {
		t.Errorf("IsDatabaseConfigCorrect - got: %t, want: %t.", isCorrect, false)
	}

	Config.Database.Driver = "aaa"
	Config.Database.Name = ""
	Config.Database.Host = ""
	isCorrect = IsDatabaseConfigCorrect()

	if isCorrect {
		t.Errorf("IsDatabaseConfigCorrect - got: %t, want: %t.", isCorrect, false)
	}

	Config.Database.Driver = "aaa"
	Config.Database.Name = "aaa"
	Config.Database.Host = ""
	isCorrect = IsDatabaseConfigCorrect()

	if isCorrect {
		t.Errorf("IsDatabaseConfigCorrect - got: %t, want: %t.", isCorrect, false)
	}

	Config.Database.Driver = "aaa"
	Config.Database.Name = ""
	Config.Database.Host = "localhost"
	isCorrect = IsDatabaseConfigCorrect()

	if isCorrect {
		t.Errorf("IsDatabaseConfigCorrect - got: %t, want: %t.", isCorrect, false)
	}

	Config.Database.Driver = ""
	Config.Database.Name = "aaa"
	Config.Database.Host = ""
	isCorrect = IsDatabaseConfigCorrect()

	if isCorrect {
		t.Errorf("IsDatabaseConfigCorrect - got: %t, want: %t.", isCorrect, false)
	}

	Config.Database.Driver = ""
	Config.Database.Name = "aaa"
	Config.Database.Host = "localhost"
	isCorrect = IsDatabaseConfigCorrect()

	if isCorrect {
		t.Errorf("IsDatabaseConfigCorrect - got: %t, want: %t.", isCorrect, false)
	}

	Config.Database.Driver = ""
	Config.Database.Name = ""
	Config.Database.Host = "localhost"
	isCorrect = IsDatabaseConfigCorrect()

	if isCorrect {
		t.Errorf("IsDatabaseConfigCorrect - got: %t, want: %t.", isCorrect, false)
	}

	Config.Database.Driver = "aaa"
	Config.Database.Name = "aaa"
	Config.Database.Host = "localhost"
	isCorrect = IsDatabaseConfigCorrect()

	if !isCorrect {
		t.Errorf("IsDatabaseConfigCorrect - got: %t, want: %t.", !isCorrect, false)
	}
}
