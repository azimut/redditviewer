package printer

import (
	"fmt"
	"strings"

	"github.com/azimut/redditviewer/format"
	"github.com/azimut/redditviewer/human"
	"github.com/fatih/color"

	"github.com/tidwall/gjson"
)

func Childrens(r gjson.Result, op string) {
	Print_Post(r, op)
	for _, v := range r.Get("replies.data.children.#.data").Array() {
		Childrens(v, op)
	}
}
func Print_Posts(r gjson.Result, op string) {
	for _, c := range r.Array() {
		Childrens(c, op)
	}
}

func Print_Post(r gjson.Result, op string) {
	depth := int(r.Get("depth").Int())
	unix_human := human.Unix_Time(r.Get("created_utc").Int())
	resp, _ :=
		format.Format_Post(
			r.Get("body").String(),
			depth)
	fmt.Println(resp)
	author := r.Get("author").String()
	yellow := color.New(color.FgYellow).SprintFunc()
	if author == op {
		author = yellow(author)
	}
	resp, _ =
		format.Format_Post(
			fmt.Sprintf("%s(%s) - %s\n",
				author,
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
