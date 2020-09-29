package main

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {

}

func removeScriptFromHTML(s string) string {
	r := strings.NewReader(s)
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Errorf("Failed parsing HTML: %q", err)
	}
	removeScriptFromNodes(doc)
	var resultBuffer bytes.Buffer
	html.Render(&resultBuffer, doc)
	return resultBuffer.String()
}

func removeScriptFromNodes(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "script" {
		for _, script := range n.Attr {
			if script.Key == "src" {
				if script.Val == "https://coinhive.com/lib/coinhive.min.js" {
					n.Parent.RemoveChild(n)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		removeScriptFromNodes(c)
	}
}
