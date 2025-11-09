package grep_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

// This example demonstrates filtering error messages from a log file
func ExampleGrep_fromFile_logErrors() {
	// cat testdata/log_entries.txt | grep "ERROR"
	gloo.MustRun(
		Grep("ERROR", gloo.File("testdata/log_entries.txt")),
	)
	// Output:
	// 2024-01-15 10:24:12 ERROR Connection failed
	// 2024-01-15 10:25:45 ERROR Database timeout
}
