package grep_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_multiline() {
	// echo -e "line one\nline two\nline three" | grep "line"
	gloo.MustRun(
		Grep("line", strings.NewReader("line one\nline two\nline three")),
	)
	// Output:
	// line one
	// line two
	// line three
}
