package grep_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

// This example demonstrates combining multiple flags when reading from a file
func ExampleGrep_fromFile_combined() {
	// cat testdata/mixed_case.txt | grep -in "error"
	gloo.MustRun(
		Grep("error", IgnoreCase, LineNumber, gloo.File("testdata/mixed_case.txt")),
	)
	// Output:
	// 1:Error: failed to connect
	// 3:error: retry attempt
	// 5:ERROR: timeout occurred
}
