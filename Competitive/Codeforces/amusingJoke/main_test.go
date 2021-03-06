package main

import "testing"

func TestCanItBePermuted(t *testing.T) {
	var input = []struct {
		s1     string
		s2     string
		output string
		result bool
	}{
		{"SANTACLAUS", "DEDMOROZ", "SANTAMOROZDEDCLAUS", true},
		{"PAPAINOEL", "JOULUPUKKI", "JOULNAPAOILELUPUKKI", false},
		{"BABBONATALE", "FATHERCHRISTMAS", "BABCHRISTMASBONATALLEFATHER", false},
	}

	for _, test := range input {
		if CanItBePermuted(test.s1, test.s2, test.output) != test.result {
			t.Fatal("Error")
		}
	}
}
