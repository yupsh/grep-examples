package grep_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

// This example demonstrates literal string matching with special regex characters
func ExampleGrep_fromFile_fixedStrings() {
	// cat testdata/filenames.txt | grep -F "test[123]"
	yup.MustRun(
		Grep("test[123]", FixedStrings, yup.File("testdata/filenames.txt")),
	)
	// Output:
	// test[123].log
}

