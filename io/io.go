package io

import (
	"io/ioutil"
	"regexp"
	"testing"
)

var invalidPathCharRegex = regexp.MustCompile("([^a-zA-Z0-9])")

// TestDirectory returns a temporary directory for this test only. Calling TestDirectory multiple times for the same
// instance of t returns a new directory every time.
func TestDirectory(t *testing.T) string {
	if dir, err := ioutil.TempDir("", normalizeTestName(t)); err != nil {
		t.Fatal(err)
		return ""
	} else {
		return dir
	}
}

func normalizeTestName(t *testing.T) string {
	return invalidPathCharRegex.ReplaceAllString(t.Name(), "_")
}
