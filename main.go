package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	targetFile := os.Args[1]
	file, err := os.Open(targetFile)
	if err != nil {
		log.Fatalf("Failed opening target file. Error: %v", err)
	}
	fileContents, _ := ioutil.ReadAll(file)
	fileContentsString := string(fileContents)
	cleanedHTML, err := removeScriptFromHTML(fileContentsString)
	if err != nil {
		log.Fatalf("Failed removing script from HTML. Error: %v", err)
	}
	err = writeStringToFile(targetFile, cleanedHTML)
	if err != nil {
		log.Fatalf("Failed writing cleaned result back to file. Error: %v", err)
	}
}

func removeScriptFromHTML(s string) (string, error) {
	r := strings.NewReader(s)
	doc, err := html.ParseWithOptions(r)
	if err != nil {
		return "", err
	}
	removeScriptFromNodes(doc)
	var resultBuffer bytes.Buffer
	html.Render(&resultBuffer, doc)
	return resultBuffer.String(), nil
}

func removeScriptFromNodes(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, div := range n.Attr {
			if div.Key == "id" {
				if div.Val == "block-block-4" {
					n.Parent.RemoveChild(n)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		removeScriptFromNodes(c)
	}
}

func writeStringToFile(fileName string, s string) error {
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.WriteString(s)
	if err != nil {
		return err
	}
	return nil
}
