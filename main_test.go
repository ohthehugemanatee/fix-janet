package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func TestFindScript(t *testing.T) {
	dirtyHTML, err := ioutil.ReadFile("test/19.html")
	if err != nil {
		log.Fatalf("Could not read dirty HTML file. Error: %q", err)
	}
	dirtyHTMLString := string(dirtyHTML)
	cleanHTML, _ := removeScriptFromHTML(dirtyHTMLString)
	if strings.Contains(cleanHTML, "coinhive") {
		t.Error("Coinhive external script include was not removed from the HTML")
	}
	if strings.Contains(cleanHTML, "new CoinHive.Anonymous") {
		t.Error("Coinhive inline script was not removed from the HTML")
	}

}

func TestWriteFile(t *testing.T) {
	fileName := "testFile.tmp"
	fileContents := "You came in that thing? You're braver than I thought."
	writeStringToFile(fileName, fileContents)
	defer os.Remove(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		t.Errorf("Failed opening created file. Error: %v", err)
	}
	realContents, _ := ioutil.ReadAll(file)
	realContentsString := string(realContents)
	if fileContents != realContentsString {
		t.Errorf("Contents of the file are incorrect. Got %v", realContentsString)
	}
}
