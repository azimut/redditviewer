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
	s := string(dat)
	post := gjson.Get(s, "0.data.children.0.data")
	comments := gjson.Get(s, "1.data.children.#.data")
	n_comments := post.Get("num_comments").Int()
	printer.Print_Header(post)
	if n_comments > 0 {
		printer.Print_Posts(comments)
	}
}
