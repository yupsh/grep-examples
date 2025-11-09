package grep_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_wholeWord() {
	// echo -e "cat\ncatch\ncatapult" | grep -w "cat"
	gloo.MustRun(
		Grep("cat", WholeWord, FixedStrings, strings.NewReader("cat\ncatch\ncatapult")),
	)
	// Output:
	// cat
}
