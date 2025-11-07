package grep_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/grep"
)

// This example demonstrates filtering out comments from a config file
func ExampleGrep_fromFile_invertMatch() {
	// cat testdata/config_file.txt | grep -v "^#"
	yup.MustRun(
		Grep("^#", Invert, yup.File("testdata/config_file.txt")),
	)
	// Output:
	// host=localhost
	// port=5432
	// timeout=30
	// debug=false
}

