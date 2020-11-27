package main

import (
	"flag"

	"github.com/azimut/redditviewer/printer"
	"github.com/azimut/redditviewer/request"

	"github.com/tidwall/gjson"
)

var timeout int
var uri string

func init() {
	flag.IntVar(&timeout, "t", 5, "timeout after seconds")
	flag.StringVar(&uri, "u", "", "url")
}

func main() {
	flag.Parse()

	data, err := request.GetFromParam(timeout, uri)
	if err != nil {
		panic(err)
	}

	post := gjson.Get(data, "0.data.children.0.data")
	printer.Print_Header(post)

	num_comments := post.Get("num_comments").Int()
	if num_comments > 0 {
		comments := gjson.Get(data, "1.data.children.#.data")
		printer.Print_Posts(comments)
	}
}
