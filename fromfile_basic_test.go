package grep_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleGrep_fromFile_basic() {
	// cat testdata/fruits.txt | grep "ap"
	gloo.MustRun(
		Grep("ap", gloo.File("testdata/fruits.txt")),
	)
	// Output:
	// apple
	// apricot
}
