# Test Template Guide for yupsh Grep Command

This directory demonstrates how to write tests for yupsh commands following best practices. Use this as a reference when creating examples for other commands.

## Directory Structure

```
examples/commands/grep/
├── README.md                    # Comprehensive guide organized by complexity
├── TEMPLATE_GUIDE.md           # This file
├── [feature1]_test.go          # One test file per feature
├── [feature2]_test.go
├── ...
├── go.mod
└── go.sum
```

## File Organization Principles

### 1. One Feature Per File

✅ **DO**: Create separate test files for each distinct feature
```
basicmatch_test.go          # Simple substring matching
ignorecase_test.go          # Case-insensitive flag
regex_test.go               # Regular expression patterns
```

❌ **DON'T**: Combine multiple features in one large file
```
command_test.go             # Everything together
```

### 2. Descriptive Filenames

Use lowercase, descriptive names that match the feature:
- `basicmatch_test.go` - not `test1_test.go`
- `ignorecase_test.go` - not `case_test.go`
- `combinedflags_test.go` - not `flags_test.go`

### 3. Complexity Organization

Organize examples in the README from simple to sophisticated:

1. **Basic Pattern Matching** - Simple string searches
2. **Regular Expression Patterns** - Advanced regex features
3. **Flag Options** - Individual flags
4. **Combined Features** - Multiple flags together

## Test File Template

Every test file should follow this exact structure:

```go
package grep_test

import (
	"strings"

	gloo "github.com/yupsh/framework"
	. "github.com/yupsh/grep"
)

func ExampleGrep_[featureName]() {
	// echo -e "input" | grep [flags] "pattern"
	gloo.MustRun(
		Grep("pattern",[flags...], strings.NewReader("input")),
)
// Output:
// expected output
}
```

## Documentation Guidelines

### Package-Level Documentation

For **test packages** (`*_test`), package-level documentation goes in the **README.md**, not in a `doc.go` file. Test packages should not have package comments in the Go files.

**Test files:**
```go
package grep_test  // No package comment needed

import (...)
```

The README.md serves as the comprehensive documentation for the example package.

### Example Comments

Every example function should have:
1. A comment showing the traditional command equivalent
2. Clear, minimal input
3. Expected output in `// Output:` comment

```go
func ExampleGrep_ignoreCase() {
	// echo -e "Apple\nBanana\napricot" | grep -i "APPLE"
	gloo.MustRun(
		Grep("APPLE", IgnoreCase, strings.NewReader("Apple\nBanana\napricot")),
	)
	// Output:
	// Apple
}
```

## README Structure

Every command's README should follow this structure:

```markdown
# [Command Name] Examples

Brief introduction paragraph.

## Organization

Description of how examples are organized.

### [Category 1 Name]

Brief description of this category.

1. **[filename_test.go](filename_test.go)** - Brief description
   ```bash
   # Equivalent: traditional command
   ```

2. **[filename2_test.go](filename2_test.go)** - Brief description
   ```bash
   # Equivalent: traditional command
   ```

### [Category 2 Name]

...

## Running Examples

Instructions for running tests.

## Writing Style Guide

Command-specific style guidelines.

## Key Features

Important features for this command.

## Learning Path

Recommended order for learning.

## See Also

Links to related documentation.
```

## Code Style Guidelines

### 1. Use MustRun for Examples

✅ **DO**:
```go
gloo.MustRun(
	Grep("pattern", strings.NewReader("input")),
)
```

❌ **DON'T**:
```go
err := gloo.Run(Grep("pattern", strings.NewReader("input")))
if err != nil {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
}
```

### 2. Minimal Imports

Only import what you need for each example:

✅ **DO**:
```go
import (
	"strings"

	gloo "github.com/yupsh/framework"
	. "github.com/yupsh/grep"
)
```

❌ **DON'T**:
```go
import (
	"fmt"
	"os"
	"io"
	"strings"
	// ... everything
)
```

### 3. Clear, Minimal Examples

Each example should demonstrate **one concept** clearly:

✅ **DO**: Simple, focused example
```go
func ExampleGrep_invertMatch() {
	gloo.MustRun(
		Grep("ap", Invert, strings.NewReader("apple\nbanana\napricot")),
	)
	// Output:
	// banana
}
```

❌ **DON'T**: Complex example doing multiple things
```go
func ExampleGrep_everything() {
	// 50 lines of setup
	// Multiple concepts mixed together
	// Unclear what's being demonstrated
}
```

### 4. Use Dot Import for Command Package

Always use dot import for the command package to make examples cleaner:

✅ **DO**:
```go
import (
	. "github.com/yupsh/grep"
)

func ExampleGrep_feature() {
	gloo.MustRun(
		Grep("pattern", IgnoreCase, reader),
	)
}
```

❌ **DON'T**:
```go
import (
	grep "github.com/yupsh/grep"
)

func ExampleGrep_feature() {
	gloo.MustRun(
		grep.Grep("pattern", grep.IgnoreCase, reader),
	)
}
```

### 5. Use Descriptive Function Names

Function names should be `ExampleGrep_<feature>` with clear feature names:

✅ **DO**:
```go
func ExampleGrep_ignoreCase() { ... }
func ExampleGrep_invertMatch() { ... }
func ExampleGrep_combinedFlags() { ... }
```

❌ **DON'T**:
```go
func ExampleGrep_test1() { ... }
func ExampleGrep_feature() { ... }
func ExampleGrep() { ... }  // Too generic
```

## Testing Your Examples

Before considering your examples complete:

1. **All tests pass**: `go test -v`
2. **No linter errors**: Check with your linter
3. **Examples are runnable**: Each can be executed independently
4. **Output is correct**: All `// Output:` comments match actual output
5. **Documentation is clear**: README is comprehensive and well-organized
6. **Godoc renders properly**: Comments follow godoc conventions

## Checklist for New Command Examples

- [ ] Created separate test file for each feature
- [ ] No package-level godoc in test files (use README.md instead)
- [ ] Added equivalent command comment to each example
- [ ] All examples use `gloo.MustRun()`
- [ ] Organized imports (stdlib, then framework, then command)
- [ ] Created comprehensive README.md
- [ ] Organized README from simple to sophisticated
- [ ] All tests pass
- [ ] No linter errors
- [ ] Used descriptive filenames
- [ ] Examples demonstrate one concept each
- [ ] Clear `// Output:` comments
- [ ] Used dot import for command package

## Benefits of This Structure

1. **Easy to Find** - Each feature is in its own file
2. **Easy to Learn** - README provides learning path
3. **Easy to Maintain** - Changes to one feature don't affect others
4. **Easy to Test** - Run specific examples: `go test -run ExampleGrep_feature`
5. **Easy to Document** - Godoc automatically generates API docs
6. **Easy to Copy** - Use any example as a starting point
7. **Professional** - Follows Go best practices

## Example Commands Using This Template

- `examples/commands/awk/` - The original reference implementation
- `examples/commands/grep/` - This directory
- _(Add other commands as they adopt this pattern)_

## Questions?

If you're unsure about how to structure examples for a new command:
1. Look at the `awk` and `grep` examples as reference
2. Follow the patterns in this guide
3. Keep it simple and focused
4. One feature per file, one concept per example

---

**Remember**: These examples are not just tests—they're documentation, tutorials, and the first place users will look to learn your command. Make them excellent!
