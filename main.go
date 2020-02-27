package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	_"io"
	_"io/ioutil"
	"log"
	_"golang.org/x/net/html"
	"os"
	_"reflect"
)

type Link struct {
	Href string
	Text string
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer // is a variable-sized buffer of bytes with read and write methods
	w := io.Writer(&buf) // creates writer
	html.Render(w, n)  // returns a flush that writes any buffered data to the underlying stream
	return buf.String()
}


func main(){
	file, err := os.Open("ex1.html") // open the file

	if err != nil {
		log.Fatal(err)
	}

	z, err := html.Parse(file) // returns for the html from the given reader

	if err != nil {
		log.Fatal(err)
	}

	var links []Link
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					links = append(links, Link{Href: a.Val, Text: renderNode(n.FirstChild)})
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(z)

	_ = file.Close()

	// io.lib works with methods for working with streams of bytes
}
