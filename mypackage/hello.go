package mypackage

import (
	"strings"
)

func Mysplit(sentence string) []string {
	// http://golang.org/pkg/strings/#Join
	return strings.Split(sentence, " ")
}

func Myjoin(parts []string) string {
	// http://golang.org/pkg/strings/#Split
	return strings.Join(parts, "|")
}
