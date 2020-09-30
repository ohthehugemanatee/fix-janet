package main

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestFindScript(t *testing.T) {
	dirtyHTML, err := ioutil.ReadFile("test/19.html")
	if err != nil {
		log.Fatalf("Could not read dirty HTML file. Error: %q", err)
	}
	dirtyHTMLString := string(dirtyHTML)
	cleanHTML := removeScriptFromHTML(dirtyHTMLString)
	if strings.Contains(cleanHTML, "coinhive") {
		t.Error("Coinhive text was not removed from the HTML")
	}
}
