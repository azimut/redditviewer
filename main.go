package main

import (
	"flag"

	"github.com/azimut/redditviewer/printer"
	"github.com/azimut/redditviewer/request"

	"github.com/tidwall/gjson"
)

// TODO: polymorfism? for hostname
// TODO: validate that url.Path is a comment url
// TODO: free json string?

func main() {
	//dat, err := ioutil.ReadFile("/home/sendai/.json.2")
	var timeout int
	flag.IntVar(&timeout, "t", 5, "timeout after seconds")
	flag.Parse()
	dat, err := request.GetFromParam(timeout)
	if err != nil {
		panic(err)
	}
	post := gjson.Get(dat, "0.data.children.0.data")
	printer.Print_Header(post)
	num_comments := post.Get("num_comments").Int()
	if num_comments > 0 {
		comments := gjson.Get(dat, "1.data.children.#.data")
		printer.Print_Posts(comments)
	}
}
