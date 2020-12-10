package main

import (
	"flag"
	"fmt"

	"github.com/azimut/redditviewer/printer"
	"github.com/azimut/redditviewer/request"
	"github.com/fatih/color"

	"github.com/tidwall/gjson"
)

var timeout int
var uri string

func init() {
	flag.IntVar(&timeout, "t", 5, "timeout after seconds")
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

	post := gjson.Get(data, "0.data.children.0.data")
	printer.Print_Header(post)

	num_comments := post.Get("num_comments").Int()
	if num_comments > 0 {
		comments := gjson.Get(data, "1.data.children.#.data")
		author := post.Get("author").String()
		printer.Print_Posts(comments, author)
	}
}
