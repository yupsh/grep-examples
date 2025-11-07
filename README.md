# Grep Command Examples

This directory contains comprehensive examples for the yupsh grep command, organized from simplest to most sophisticated. Each example is in its own test file and demonstrates a specific feature or pattern.

## Organization

Examples are organized by complexity level to help you learn progressively:

### Basic Pattern Matching

Simple examples that demonstrate basic string and pattern matching.

1. **[basicmatch_test.go](basicmatch_test.go)** - Basic substring matching
   ```bash
   # Equivalent: grep "ap"
   ```

2. **[multiline_test.go](multiline_test.go)** - Matching across multiple lines
   ```bash
   # Equivalent: grep "line"
   ```

3. **[fixedstrings_test.go](fixedstrings_test.go)** - Literal string matching (no regex)
   ```bash
   # Equivalent: grep -F "file[txt"
   ```

### Regular Expression Patterns

Examples using regex patterns for advanced matching.

4. **[regex_test.go](regex_test.go)** - Basic regex with character classes
   ```bash
   # Equivalent: grep "test[0-9]+"
   ```

5. **[beginanchor_test.go](beginanchor_test.go)** - Match at line beginning with `^`
   ```bash
   # Equivalent: grep "^hello"
   ```

6. **[endanchor_test.go](endanchor_test.go)** - Match at line end with `$`
   ```bash
   # Equivalent: grep "txt$"
   ```

7. **[alternation_test.go](alternation_test.go)** - Multiple patterns with `|`
   ```bash
   # Equivalent: grep "apple\|cherry"
   ```

8. **[emptylines_test.go](emptylines_test.go)** - Match empty lines
   ```bash
   # Equivalent: grep -v "^$"
   ```

### Flag Options

Examples demonstrating grep's various flags and options.

9. **[ignorecase_test.go](ignorecase_test.go)** - Case-insensitive matching
   ```bash
   # Equivalent: grep -i "APPLE"
   ```

10. **[invertmatch_test.go](invertmatch_test.go)** - Invert match (show non-matching lines)
    ```bash
    # Equivalent: grep -v "ap"
    ```

11. **[linenumber_test.go](linenumber_test.go)** - Show line numbers
    ```bash
    # Equivalent: grep -n "ap"
    ```

12. **[wholeword_test.go](wholeword_test.go)** - Match whole words only
    ```bash
    # Equivalent: grep -w "cat"
    ```

### Combined Features

Advanced examples combining multiple flags and complex patterns.

13. **[combinedflags_test.go](combinedflags_test.go)** - Multiple flags together
    ```bash
    # Equivalent: grep -in "error"
    ```

14. **[caseinsensitiveregex_test.go](caseinsensitiveregex_test.go)** - Case-insensitive regex
    ```bash
    # Equivalent: grep -i "error [0-9]+"
    ```

### Reading from Files

Real-world examples that read from test data files instead of inline strings. These demonstrate practical usage patterns.

15. **[fromfile_basic_test.go](fromfile_basic_test.go)** - Basic pattern matching from file
    ```bash
    # Equivalent: cat testdata/fruits.txt | grep "ap"
    ```

16. **[fromfile_logerrors_test.go](fromfile_logerrors_test.go)** - Filter errors from log file
    ```bash
    # Equivalent: cat testdata/log_entries.txt | grep "ERROR"
    ```

17. **[fromfile_linenumbers_test.go](fromfile_linenumbers_test.go)** - Find with line numbers in code
    ```bash
    # Equivalent: cat testdata/code_snippets.txt | grep -n "var"
    ```

18. **[fromfile_caseinsensitive_test.go](fromfile_caseinsensitive_test.go)** - Case-insensitive search in logs
    ```bash
    # Equivalent: cat testdata/mixed_case.txt | grep -i "error"
    ```

19. **[fromfile_invertmatch_test.go](fromfile_invertmatch_test.go)** - Filter out comments from config
    ```bash
    # Equivalent: cat testdata/config_file.txt | grep -v "^#"
    ```

20. **[fromfile_fixedstrings_test.go](fromfile_fixedstrings_test.go)** - Match literal strings with special chars
    ```bash
    # Equivalent: cat testdata/filenames.txt | grep -F "test[123]"
    ```

21. **[fromfile_wholeword_test.go](fromfile_wholeword_test.go)** - Match whole words from file
    ```bash
    # Equivalent: cat testdata/words.txt | grep -w "cat"
    ```

22. **[fromfile_emptylines_test.go](fromfile_emptylines_test.go)** - Remove empty lines
    ```bash
    # Equivalent: cat testdata/empty_lines.txt | grep -v "^$"
    ```

23. **[fromfile_regex_test.go](fromfile_regex_test.go)** - Regex patterns on file content
    ```bash
    # Equivalent: cat testdata/log_entries.txt | grep "ERROR.*timeout"
    ```

24. **[fromfile_combined_test.go](fromfile_combined_test.go)** - Multiple flags with file input
    ```bash
    # Equivalent: cat testdata/mixed_case.txt | grep -in "error"
    ```

See [testdata/README.md](testdata/README.md) for details on the test data files.

## Running Examples

Run all examples:
```bash
go test -v
```

Run a specific example:
```bash
go test -v -run ExampleGrep_basicMatch
```

## Writing Style Guide

Each test file follows this structure:

1. **Package declaration**: `package grep_test`
2. **Imports**: Only import what's needed for that specific example
3. **Example function**: Named `ExampleGrep_<feature>` with:
   - Comment showing equivalent traditional grep command
   - Call to `yup.MustRun()` with the grep command
   - `// Output:` comment with expected output

### Template

```go
package grep_test

import (
	"strings"

	yup "github.com/yupsh/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_myFeature() {
	// echo -e "input" | grep [flags] "pattern"
	yup.MustRun(
		Grep("pattern", [flags...], strings.NewReader("input")),
	)
	// Output:
	// expected output
}
```

## Key Features

### Pattern Matching

The first argument to `Grep()` is the pattern:
- **String literal**: `Grep("hello", ...)` - matches "hello" anywhere in line
- **Regular expression**: `Grep("^test[0-9]+", ...)` - full regex support
- **Fixed strings**: `Grep("file[txt", FixedStrings, ...)` - literal matching

### Flags

Common flags available:
- `IgnoreCase` - Case-insensitive matching (like `-i`)
- `LineNumber` - Show line numbers (like `-n`)
- `Invert` - Show non-matching lines (like `-v`)
- `WholeWord` - Match whole words only (like `-w`)
- `FixedStrings` - Treat pattern as literal string (like `-F`)

Flags can be combined:
```go
Grep("error", IgnoreCase, LineNumber, Invert, reader)
```

### Input

Pass input as `io.Reader`:

**From string:**
```go
strings.NewReader("line1\nline2\nline3")
```

**From file:**
```go
yup.File("testdata/data.txt")
```

## Common Patterns

### Filter log files
```go
Grep("ERROR", IgnoreCase, strings.NewReader(logContent))
```

### Find exact matches
```go
Grep("exact", WholeWord, FixedStrings, strings.NewReader(text))
```

### Exclude lines
```go
Grep("debug", Invert, strings.NewReader(output))
```

### Search with line numbers
```go
Grep("TODO", LineNumber, strings.NewReader(code))
```

## Learning Path

1. Start with **Basic Pattern Matching** (1-3) to understand how grep works
2. Move to **Regular Expression Patterns** (4-8) to learn regex syntax
3. Explore **Flag Options** (9-12) for different matching modes
4. Apply **Combined Features** (13-14) for advanced use cases
5. Practice with **Reading from Files** (15-24) for real-world scenarios

## See Also

- [Grep Command Documentation](../../../grep/README.md)
- [yupsh Framework Documentation](../../../framework/DESIGN.md)
- [Go Examples Documentation](https://go.dev/blog/examples)

