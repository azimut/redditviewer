package main

import (
	"fmt"
	//	"github.com/rivo/tview"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/url"
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

func Comments(n_comments int64) string {
	switch n_comments {
	case 0:
		return "No comments"
	case 1:
		return fmt.Sprintf("%d comment\n", n_comments)
	default:
		return fmt.Sprintf("%d comments\n", n_comments)
	}
}

func childrens(r gjson.Result) {
	body := r.Get("body").String()
	fmt.Println("acomment:", body)
	for _, v := range r.Get("replies.data.children.#.data").Array() {
		childrens(v)
	}
}

func cm(r string) {
	for _, c := range gjson.Get(r, "1.data.children.#.data").Array() {
		childrens(c)
	}
}

func main() {
	dat, err := ioutil.ReadFile("/home/sendai/.json.2")
	if err != nil {
		panic(err)
	}
	///gjson.Valid()
	fmt.Println(string(dat)[0:20])
	post := gjson.Get(string(dat), "0.data.children.0.data")
	fmt.Println("title:", post.Get("title"))
	fmt.Println("url:", post.Get("url"))
	fmt.Println("permalink:", post.Get("permalink"))
	fmt.Println("selftext:", post.Get("selftext"))
	fmt.Println("ups:", post.Get("ups"))
	fmt.Println("author:", post.Get("author"))
	n_comments := post.Get("num_comments").Int()
	fmt.Println(Comments(n_comments))
	cm(string(dat))
	// comments := gjson.Get(string(dat), "1.data.children").Array()
	// fmt.Println(comments[0])
	//
	// rootDir := "."
	// root := tview.
	// 	NewTreeNode(rootDir)
	// tree := tview.
	// 	NewTreeView().
	// 	SetRoot(root).
	// 	SetCurrentNode(root)

	// // A helper function which adds the files and directories of the given path
	// // to the given target node.
	// add := func(target *tview.TreeNode, path string) {
	// 	files, err := ioutil.ReadDir(path)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	for _, file := range files {
	// 		node := tview.NewTreeNode(file.Name()).
	// 			SetReference(filepath.Join(path, file.Name())).
	// 			SetSelectable(file.IsDir())
	// 		target.AddChild(node)
	// 	}
	// }

	// // Add the current directory to the root node.
	// add(root, rootDir)
	// //
	// if err := tview.NewApplication().SetRoot(tree, true).Run(); err != nil {
	// 	panic(err)
	// }
}
