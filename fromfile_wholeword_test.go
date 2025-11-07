package grep_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

// This example demonstrates whole word matching to avoid partial matches
func ExampleGrep_fromFile_wholeWord() {
	// cat testdata/words.txt | grep -w "cat"
	yup.MustRun(
		Grep("cat", WholeWord, FixedStrings, yup.File("testdata/words.txt")),
	)
	// Output:
	// cat
}

