package grep_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

// This example demonstrates regex pattern matching in log files
func ExampleGrep_fromFile_regex() {
	// cat testdata/log_entries.txt | grep "202[0-9]-[0-9][0-9]-[0-9][0-9]"
	gloo.MustRun(
		Grep("ERROR.*timeout", gloo.File("testdata/log_entries.txt")),
	)
	// Output:
	// 2024-01-15 10:25:45 ERROR Database timeout
}
