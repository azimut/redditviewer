package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/url"
	"redditviewer/format"
)

// TODO: polymorfism? for hostname
// TODO: validate that url.Path is a comment url
// TODO: free json string?
func ruler() {
	s := "https://www.reddit.com/r/politics/comments/9wqvmc/federal_judge_finds_georgia_county_violated_civil/.json"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("host:", u.Host)
	fmt.Println("path:", u.Path)
}

func Childrens(r gjson.Result) {
	Format_Post(r)
	for _, v := range r.Get("replies.data.children.#.data").Array() {
		Childrens(v)
	}
}
func Parents(r string) {
	for _, c := range gjson.Get(r, "1.data.children.#.data").Array() {
		Childrens(c)
	}
}

func Format_Post(r gjson.Result) {
	depth := int(r.Get("depth").Int())
	resp, _ :=
		format.Format_Line(
			fmt.Sprintf("%s %s %s\n",
				r.Get("score").String(),
				r.Get("author").String(),
				r.Get("created_utc").String()),
			depth)
	fmt.Println(resp)
	resp, _ =
		format.Format_Line(
			r.Get("body").String(),
			depth)
	fmt.Println(resp)
	fmt.Println()
}

func main() {
	dat, err := ioutil.ReadFile("/home/sendai/.json.2")
	if err != nil {
		panic(err)
	}

	post := gjson.Get(string(dat), "0.data.children.0.data")
	n_comments := post.Get("num_comments").Int()
	fmt.Println("title:", post.Get("title"))
	fmt.Println("url:", post.Get("url"))
	fmt.Println("selftext:", post.Get("selftext"))
	fmt.Printf("%d - %s - %d Comment(s)\n",
		post.Get("ups").Int(),
		post.Get("author"),
		n_comments)
	fmt.Println()
	if n_comments > 0 {
		Parents(string(dat))
	}
	// comments := gjson.Get(string(dat), "1.data.children").Array()
	// fmt.Println(comments[0])
}
