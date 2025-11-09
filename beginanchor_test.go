package grep_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_beginAnchor() {
	// echo -e "hello world\nworld hello\nhello" | grep "^hello"
	gloo.MustRun(
		Grep("^hello", strings.NewReader("hello world\nworld hello\nhello")),
	)
	// Output:
	// hello world
	// hello
}
