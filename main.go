package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	tagstats := make(map[string]int)
	var charCount int

	if 2 == len(os.Args) {

		url := os.Args[1]
		resp, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}

		// Heavily modified form of example on golang.org/x/net/html

		doc, err := html.Parse(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var walkHTML func(*html.Node)

		// Recursive function to work through HTML and act on nodes
		walkHTML = func(readHead *html.Node) {
			// Why readHead? It reflects what's going on: read, advance, read, etc.

			// The original used single character variables (n, c, etc),
			// which makes them harder to work with in an IDE/editor.

			// Works on current node
			if readHead.Type == html.ElementNode {
				for _, tag := range readHead.Attr {
					tagstats[tag.Key]++
					charCount += len(tag.Val)
				}
			}
			// Goes ever-deeper with each iteration
			for childNode := readHead.FirstChild; childNode != nil; childNode = childNode.NextSibling {
				walkHTML(childNode)
			}
		}

		// Walk the doc (get it)
		walkHTML(doc)
		fmt.Printf("Map of tag counts: %v\n", tagstats)
		fmt.Printf("Total characters of tag content: %d", charCount)
	} else {
		log.Printf("Usage: %s <url>", os.Args[1])

	}

}
