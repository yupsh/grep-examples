package grep_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_basicMatch() {
	// echo -e "apple\nbanana\napricot\ncherry" | grep "ap"
	yup.MustRun(
		Grep("ap", strings.NewReader("apple\nbanana\napricot\ncherry")),
	)
	// Output:
	// apple
	// apricot
}

