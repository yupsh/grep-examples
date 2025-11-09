package grep_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_emptyLines() {
	// echo -e "line1\n\nline3\n\nline5" | grep -v "^$"
	gloo.MustRun(
		Grep("^$", Invert, strings.NewReader("line1\n\nline3\n\nline5")),
	)
	// Output:
	// line1
	// line3
	// line5
}
