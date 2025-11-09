package grep_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

// This example demonstrates case-insensitive matching in log files
func ExampleGrep_fromFile_caseInsensitive() {
	// cat testdata/mixed_case.txt | grep -i "error"
	gloo.MustRun(
		Grep("error", IgnoreCase, gloo.File("testdata/mixed_case.txt")),
	)
	// Output:
	// Error: failed to connect
	// error: retry attempt
	// ERROR: timeout occurred
}
