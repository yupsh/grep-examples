package grep_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_multiline() {
	// echo -e "line one\nline two\nline three" | grep "line"
	yup.MustRun(
		Grep("line", strings.NewReader("line one\nline two\nline three")),
	)
	// Output:
	// line one
	// line two
	// line three
}

