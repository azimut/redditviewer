package main

import (
	"flag"
	"fmt"

	"github.com/azimut/redditviewer/human"
	"github.com/azimut/redditviewer/printer"
	"github.com/azimut/redditviewer/request"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/tidwall/gjson"
)

var timeout int
var uri string
var max_width int

func init() {
	flag.IntVar(&timeout, "t", 5, "timeout after seconds")
	flag.IntVar(&max_width, "w", 120, "max width")
	flag.StringVar(&uri, "u", "", "url")
	color.NoColor = false
}

func main() {
	flag.Parse()

	if uri == "" {
		panic(fmt.Errorf("-u parameter not provided"))
	}

	data, err := request.GetFromParam(timeout, uri)
	if err != nil {
		panic(err)
	}

	width, _, err := terminal.GetSize(0)
	if err != nil {
		panic(err)
	}
	width = human.Min(width, max_width)

	post := gjson.Get(data, "0.data.children.0.data")
	printer.Print_Header(post, width)

	num_comments := post.Get("num_comments").Int()
	if num_comments > 0 {
		comments := gjson.Get(data, "1.data.children.#.data")
		author := post.Get("author").String()
		printer.Print_Posts(comments, author, width)
	}
}
