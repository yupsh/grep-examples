package grep_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_alternation() {
	// echo -e "apple\nbanana\ncherry\napricot" | grep "apple\|cherry"
	yup.MustRun(
		Grep("apple|cherry", strings.NewReader("apple\nbanana\ncherry\napricot")),
	)
	// Output:
	// apple
	// cherry
}

