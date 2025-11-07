package grep_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_invertMatch() {
	// echo -e "apple\nbanana\napricot" | grep -v "ap"
	yup.MustRun(
		Grep("ap", Invert, strings.NewReader("apple\nbanana\napricot")),
	)
	// Output:
	// banana
}

