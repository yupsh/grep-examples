package grep_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

// This example demonstrates removing empty lines from a file
func ExampleGrep_fromFile_emptyLines() {
	// cat testdata/empty_lines.txt | grep -v "^$"
	yup.MustRun(
		Grep("^$", Invert, yup.File("testdata/empty_lines.txt")),
	)
	// Output:
	// line one
	// line three
	// line five
}

