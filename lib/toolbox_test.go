package lib

import (
	"testing"
)

// TestUcfirst
func TestUcfirst(t *testing.T) {
	s1 := "test chaîne Avec et sans majuscule"
	s1u := Ucfirst(s1)

	if s1u != "Test chaîne Avec et sans majuscule" {
		t.Errorf("Ucfirst - got: %q, want: %q.", s1u, "Test chaîne Avec et sans majuscule")
	}

	s2 := "Test chaîne Avec et sans majuscule"
	s2u := Ucfirst(s2)

	if s2u != "Test chaîne Avec et sans majuscule" {
		t.Errorf("Ucfirst - got: %q, want: %q.", s2u, "Test chaîne Avec et sans majuscule")
	}

	s3 := "été"
	s3u := Ucfirst(s3)

	if s3u != "Été" {
		t.Errorf("Ucfirst - got: %q, want: %q.", s3u, "Été")
	}
}

// TestInArray
func TestInArray(t *testing.T) {
	tInt := []int{10, 56, 23, 85}
	found, index := InArray(56, tInt)
	foundWanted, indexWanted := true, 1

	if found != foundWanted || indexWanted != indexWanted {
		t.Errorf("Ucfirst - got: %t, %d, want: %t, %d.", found, index, foundWanted, indexWanted)
	}

	tInt = []int{10, 56, 23, 85}
	found, index = InArray(589, tInt)
	foundWanted, indexWanted = true, -1

	if found == foundWanted || indexWanted > -1 {
		t.Errorf("Ucfirst - got: %t, %d, want: %t, %d.", found, index, foundWanted, indexWanted)
	}

	tString := []string{"45", "ghgh", "kl7878"}
	found, index = InArray("kl7878", tString)
	foundWanted, indexWanted = true, 2

	if found != foundWanted || index != indexWanted {
		t.Errorf("Ucfirst - got: %t, %d, want: %t, %d.", found, index, foundWanted, indexWanted)
	}

	tString = []string{"45", "ghgh", "kl7878"}
	found, index = InArray(589, tString)
	foundWanted, indexWanted = false, -1

	if found != foundWanted || index != indexWanted {
		t.Errorf("Ucfirst - got: %t, %d, want: %t, %d.", found, index, foundWanted, indexWanted)
	}
}
