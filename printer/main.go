package printer

import (
	"fmt"
	"redditviewer/format"
	"redditviewer/human"
	"strings"

	"github.com/tidwall/gjson"
)

func Childrens(r gjson.Result) {
	Print_Post(r)
	for _, v := range r.Get("replies.data.children.#.data").Array() {
		Childrens(v)
	}
}
func Print_Posts(r gjson.Result) {
	for _, c := range r.Array() {
		Childrens(c)
	}
}

func Print_Post(r gjson.Result) {
	depth := int(r.Get("depth").Int())
	unix_human := human.Unix_Time(r.Get("created_utc").Int())
	resp, _ :=
		format.Format_Post(
			r.Get("body").String(),
			depth)
	fmt.Println(resp)
	resp, _ =
		format.Format_Post(
			fmt.Sprintf("%s(%s) - %s\n",
				r.Get("author").String(),
				r.Get("score").String(),
				unix_human),
			depth)
	fmt.Println(resp)
	fmt.Println()
}

func Print_Header(r gjson.Result) {
	fmt.Println()
	// TODO: check error
	selftext, _ := format.Wrap_Line(r.Get("selftext").String(), 0)
	unix_human := human.Unix_Time(r.Get("created_utc").Int())
	fmt.Println("title:", r.Get("title"))
	fmt.Println("url:", r.Get("url"))
	if len(strings.TrimSpace(selftext)) != 0 {
		fmt.Println()
		fmt.Println(selftext)
	}
	fmt.Printf("(%d)%s - %s - %d Comment(s)\n",
		r.Get("ups").Int(),
		r.Get("author"),
		unix_human,
		r.Get("num_comments").Int())
	fmt.Println()
}
