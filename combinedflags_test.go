package grep_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_combinedFlags() {
	// echo -e "Error: failed\nWARNING: check\nerror: retry" | grep -in "error"
	yup.MustRun(
		Grep("error", IgnoreCase, LineNumber, strings.NewReader("Error: failed\nWARNING: check\nerror: retry")),
	)
	// Output:
	// 1:Error: failed
	// 3:error: retry
}

