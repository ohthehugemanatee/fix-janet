package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {

}

func removeScriptFromHTML(s string) string {
	r := strings.NewReader(s)
	doc, err := html.ParseWithOptions(r)
	if err != nil {
		fmt.Errorf("Failed parsing HTML: %q", err)
	}
	removeScriptFromNodes(doc)
	var resultBuffer bytes.Buffer
	html.Render(&resultBuffer, doc)
	return resultBuffer.String()
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

func writeStringToFile(fileName string, s string) {
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(s)
	if err != nil {
		fmt.Println(err)
		return
	}
}
