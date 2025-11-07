package grep_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

// This example demonstrates finding variables with line numbers in code
func ExampleGrep_fromFile_lineNumbers() {
	// cat testdata/code_snippets.txt | grep -n "var"
	yup.MustRun(
		Grep("var", LineNumber, yup.File("testdata/code_snippets.txt")),
	)
	// Output:
	// 2:  var count = 0;
	// 3:  for (var i = 0; i < items.length; i++) {
}

