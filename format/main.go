package format

import (
	"github.com/bbrks/wrap"
	"golang.org/x/crypto/ssh/terminal"
	"strings"
)

func Limit_Line(post string, depth int) (string, error) {
	width, _, err := terminal.GetSize(0)
	if err != nil {
		return "", err
	}
	return wrap.Wrap(post, width-depth*4), nil
}

func Indent_Line(post string, depth int) string {
	var post_indented []string
	indent := strings.Repeat(" ", depth*3)
	for _, s := range strings.Split(post, "\n") {
		post_indented = append(post_indented, indent+"｜"+s)
	}
	return strings.Join(post_indented, "\n")
}

func Format_Line(post string, depth int) (string, error) {
	splitted, err := Limit_Line(post, depth)
	if err != nil {
		return "", (err)
	}
	splitted = strings.TrimSpace(splitted)
	return Indent_Line(splitted, depth), nil
}
