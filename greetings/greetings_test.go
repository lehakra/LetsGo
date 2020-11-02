package greetings

import (
	"regexp"
	"testing"
)

// Tests a valid call to Hello and checks valid return value
func TestHelloName(t *testing.T) {
	name := "Romil"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Romil") = %q, %v, want match for %q, nil`, msg, err, want)
	}
}

// Test Hello with empty string
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
