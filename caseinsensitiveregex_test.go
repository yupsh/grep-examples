package grep_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_caseInsensitiveRegex() {
	// echo -e "ERROR 123\nwarning 456\nError 789" | grep -i "error [0-9]+"
	yup.MustRun(
		Grep("error [0-9]+", IgnoreCase, strings.NewReader("ERROR 123\nwarning 456\nError 789")),
	)
	// Output:
	// ERROR 123
	// Error 789
}

