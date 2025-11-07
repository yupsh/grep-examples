package grep_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_regex() {
	// echo -e "test123\ntest\nabc456" | grep "test[0-9]+"
	yup.MustRun(
		Grep("test[0-9]+", strings.NewReader("test123\ntest\nabc456")),
	)
	// Output:
	// test123
}

