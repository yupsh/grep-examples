package grep_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_fixedStrings() {
	// echo -e "file.txt\nfile[txt\ntest.doc" | grep -F "file[txt"
	gloo.MustRun(
		Grep("file[txt", FixedStrings, strings.NewReader("file.txt\nfile[txt\ntest.doc")),
	)
	// Output:
	// file[txt
}
