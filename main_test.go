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
		t.Error("Coinhive external script include was not removed from the HTML")
	}
	if strings.Contains(cleanHTML, "new CoinHive.Anonymous") {
		t.Error("Coinhive inline script was not removed from the HTML")
	}

}
