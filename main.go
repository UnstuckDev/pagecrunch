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
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode {
				for _, tag := range n.Attr {
					tagstats[tag.Key]++
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(doc)
		fmt.Printf("%v\n", tagstats)
	} else {
		log.Printf("Usage: %s <url>", os.Args[1])

	}

}
