package printer

import (
	"fmt"
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/azimut/redditviewer/format"
	"github.com/azimut/redditviewer/human"
	"github.com/fatih/color"

	"github.com/tidwall/gjson"
)

func Print_Posts(r gjson.Result, op string) {
	for _, c := range r.Array() {
		childrens(c, op)
	}
}

func childrens(r gjson.Result, op string) {
	print_post(r, op)
	for _, v := range r.Get("replies.data.children.#.data").Array() {
		childrens(v, op)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func print_post(r gjson.Result, op string) {
	// Comment
	depth := int(r.Get("depth").Int())
	comment := markdown.Render(r.Get("body").String(), 80, max(3*depth, 1))
	fmt.Print(string(comment))
	// Check if author is op
	author := r.Get("author").String()
	yellow := color.New(color.FgYellow).SprintFunc()
	if author == op {
		author = yellow(author)
	}
	// Footer
	unix_human := human.Unix_Time(r.Get("created_utc").Int())
	score := r.Get("score").Int()
	reply, _ :=
		format.Format_Post(
			fmt.Sprintf("%s(%d) - %s\n",
				author,
				score,
				unix_human),
			depth)
	fmt.Println(reply)
	fmt.Println()
}

func Print_Header(r gjson.Result) {
	title := r.Get("title")
	url := r.Get("url")
	fmt.Printf("\ntitle: %s\nurl: %s\n", title, url)

	selftext := r.Get("selftext").String()
	if len(strings.TrimSpace(selftext)) > 0 {
		resp := markdown.Render(selftext, 80, 3)
		fmt.Printf("\n%s\n", string(resp))
	}

	upvotes := r.Get("ups").Int()
	author := r.Get("author")
	unix_human := human.Unix_Time(r.Get("created_utc").Int())
	comments := r.Get("num_comments").Int()
	fmt.Printf("%s(%d) - %s - %d Comment(s)\n\n",
		author,
		upvotes,
		unix_human,
		comments)
}
