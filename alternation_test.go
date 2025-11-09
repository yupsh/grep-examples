package grep_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_alternation() {
	// echo -e "apple\nbanana\ncherry\napricot" | grep "apple\|cherry"
	gloo.MustRun(
		Grep("apple|cherry", strings.NewReader("apple\nbanana\ncherry\napricot")),
	)
	// Output:
	// apple
	// cherry
}
