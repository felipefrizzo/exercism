// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode/utf8"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var acronym string

	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", "")
	s = strings.ReplaceAll(s, "   ", " ")

	a := strings.Split(s, " ")
	for _, str := range a {
		d, _ := utf8.DecodeRuneInString(str)
		acronym = acronym + string(d)
	}
	return strings.ToUpper(acronym)
}
