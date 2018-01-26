package main

import (
	"fmt"
	"regexp"
	"testing"
)

var (
	superstring = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	substring   = "FGHIJ"
	index       = 5
	regex       = regexp.MustCompile(fmt.Sprintf("^.{%d}%s", index, substring))
)

func BenchmarkSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = superstring[5:len(substring)] == substring
	}
}

func BenchmarkRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = regex.MatchString(superstring)
	}
}
