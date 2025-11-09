package grep_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_lineNumber() {
	// echo -e "apple\nbanana\napricot" | grep -n "ap"
	gloo.MustRun(
		Grep("ap", LineNumber, strings.NewReader("apple\nbanana\napricot")),
	)
	// Output:
	// 1:apple
	// 3:apricot
}
