package grep_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_invertMatch() {
	// echo -e "apple\nbanana\napricot" | grep -v "ap"
	gloo.MustRun(
		Grep("ap", Invert, strings.NewReader("apple\nbanana\napricot")),
	)
	// Output:
	// banana
}
