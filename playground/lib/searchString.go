package main

import (
	"strings"
)

func bandNameGenerator(word string) string {
	// Happy coding
	first := word[:1]
	last := word[len(word)-1:]

	if first == last {
		return strings.Title(word) + word[1:]
	}
	return "The " + strings.Title(word)
}

func mainBand() {
	println(bandNameGenerator("test"))
}
