package grep_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_endAnchor() {
	// echo -e "test.txt\ntest.log\nfile.txt" | grep "txt$"
	yup.MustRun(
		Grep("txt$", strings.NewReader("test.txt\ntest.log\nfile.txt")),
	)
	// Output:
	// test.txt
	// file.txt
}

