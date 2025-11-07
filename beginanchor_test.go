package grep_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_beginAnchor() {
	// echo -e "hello world\nworld hello\nhello" | grep "^hello"
	yup.MustRun(
		Grep("^hello", strings.NewReader("hello world\nworld hello\nhello")),
	)
	// Output:
	// hello world
	// hello
}

