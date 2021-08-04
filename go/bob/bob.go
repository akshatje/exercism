// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var answer string = ""
	if remark=="" {
		answer ="Fine. Be that way!"
	}
	c1:= func(remark string) bool{
		if strings.Contains(remark,"?") {
			answer = "Sure."
			return true
		}
		return false
	}

	c2:= func(remark string) bool {
		for _,char:= range remark {
			if(unicode.IsLower(char)){return false}
	}
		answer="Whoa, chill out!"
		return  true
}

	if c1(remark) && c2(remark) == true {
		answer="Calm down, I know what I'm doing!"
	}else {answer="Whatever."}
	c1(remark)
	c2(remark)
	return answer
}
