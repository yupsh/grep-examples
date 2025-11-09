package grep_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_basicMatch() {
	// echo -e "apple\nbanana\napricot\ncherry" | grep "ap"
	gloo.MustRun(
		Grep("ap", strings.NewReader("apple\nbanana\napricot\ncherry")),
	)
	// Output:
	// apple
	// apricot
}
