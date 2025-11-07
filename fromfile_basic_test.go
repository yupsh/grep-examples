package grep_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleGrep_fromFile_basic() {
	// cat testdata/fruits.txt | grep "ap"
	yup.MustRun(
		Grep("ap", yup.File("testdata/fruits.txt")),
	)
	// Output:
	// apple
	// apricot
}

