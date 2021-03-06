package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)

	if HasEvenDistinctCharacters(s) {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}

func HasEvenDistinctCharacters(s string) bool {
	var c []string
	for _, r := range s {
		c = append(c, string(r))
	}
	sort.Strings(c)

	var sortedString string

	for _, r := range c {
		sortedString += r
	}

	var sum int
	var lastCharacter rune

	for _, r := range sortedString {
		if r != lastCharacter {
			sum++
		}
		lastCharacter = r
	}

	if sum%2 == 0 {
		return true
	} else {
		return false
	}
}
