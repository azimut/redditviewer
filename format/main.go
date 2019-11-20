package format

import (
	"strings"

	"github.com/bbrks/wrap"
	"golang.org/x/crypto/ssh/terminal"
)

const SPACES int = 3

func Max_Int(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Wrap_Line(post string, depth int) (string, error) {
	width, _, err := terminal.GetSize(0)
	if err != nil {
		return "", err
	}
	return wrap.Wrap(post, width-(depth*SPACES)), nil
}

func Wrap_Post(post string, depth int) (string, error) {
	width, _, err := terminal.GetSize(0)
	if err != nil {
		return "", err
	}
	return wrap.Wrap(post, width-(depth*SPACES)-2), nil
}

func Indent_Post(post string, depth int) string {
	var post_indented []string
	indent := strings.Repeat(" ", Max_Int((depth*SPACES)-1, 0))
	for _, s := range strings.Split(post, "\n") {
		post_indented = append(post_indented, indent+"｜"+s)
	}
	return strings.Join(post_indented, "\n")
}

func Format_Post(post string, depth int) (string, error) {
	splitted, err := Wrap_Post(post, depth)
	if err != nil {
		return "", (err)
	}
	splitted = strings.TrimSpace(splitted)
	return Indent_Post(splitted, depth), nil
}
