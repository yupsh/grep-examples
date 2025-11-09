package grep_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

// This example demonstrates whole word matching to avoid partial matches
func ExampleGrep_fromFile_wholeWord() {
	// cat testdata/words.txt | grep -w "cat"
	gloo.MustRun(
		Grep("cat", WholeWord, FixedStrings, gloo.File("testdata/words.txt")),
	)
	// Output:
	// cat
}
