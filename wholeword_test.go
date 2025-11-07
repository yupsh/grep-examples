package grep_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_wholeWord() {
	// echo -e "cat\ncatch\ncatapult" | grep -w "cat"
	yup.MustRun(
		Grep("cat", WholeWord, FixedStrings, strings.NewReader("cat\ncatch\ncatapult")),
	)
	// Output:
	// cat
}

