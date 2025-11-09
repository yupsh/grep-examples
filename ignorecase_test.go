package grep_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_ignoreCase() {
	// echo -e "Apple\nBanana\napricot" | grep -i "APPLE"
	gloo.MustRun(
		Grep("APPLE", IgnoreCase, strings.NewReader("Apple\nBanana\napricot")),
	)
	// Output:
	// Apple
}
