package lib

import (
	"testing"
)

func TestInitConfig(t *testing.T) {
	InitConfig("../config.toml")

	timezone := Config.Database.Timezone != "" && Config.Database.Timezone != "Europe%2FParis"
	if timezone {
		t.Errorf("InitConfig - got: %t, want: %t.", timezone, false)
	}
}

// TestIsDatebaseConfigCorrect
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

// TestIsServerConfigCorrect
func TestIsServerConfigCorrect(t *testing.T) {
	Config.Server.AllowOrigins = []string{}
	isCorrect := IsServerConfigCorrect()

	if isCorrect {
		t.Errorf("IsServerConfigCorrect - got: %t, want: %t.", isCorrect, false)
	}
}
