go-fixpkg
=========

This is a starter project that can be copied to produce custom
`gofix`-like commands.  It is intended for those familiar with how much
trouble can be caused by changing the public APIs of Go packages, and
can help automate some of the work for those who decide to do it anyway.

Getting Started
---------------

Assuming you have a Go package with a public API that is changing

1.	Copy (or clone, submodule add, ...) this repository into a
	subdirectory of your package (e.g. "gomypkgfix").  Alternatively,
	simply copy these files from the Go source tree:
	- `src/cmd/fix/fix.go`
	- `src/cmd/fix/main.go`
	- `src/cmd/fix/main_test.go`
	- `LICENSE`
2.	In your package's documentation, include instructions somewhat
	like these:
	
	> Package "mypkg" has made some public changes that will break old
	> code.  Fortunately, a command has been written based on `go fix`
	> that will upgrade your code for you.  After backing up your code,
	> run these commands:
	>
	> ```
	> go get <path to mypkg>/gomypkgfix
	> cd $YOUR_PACKAGE
	> gomypkgfix .
	> ```
3.	Make sure those instructions actually work without error.

So far, your `gomypkgfix` command doesn't do anything.  To change this
you have to include a working fix.  There are some examples included in
this repository (or there will be soon), but you can find more in the Go
source tree under `src/cmd/fix`.

Writing a Fix
-------------

Each fix is a pair of files: the fix itself, and the tests that validate
the fix (test-driven development is highly recommended).  Here is a very
small fix, `minimal.go`:

```go
package main

import (
	"go/ast"
)

func init() {
	register(minimalFix)
}

var minimalFix = fix{
	"minimal",
	"2013-07-21", // when this fix was published
	minimal,
	`
	Don't change anything.
`,
}

func minimal(f *ast.File) bool {
	// This method MUST return true if and only if it has
	// made changes to f.
	return false
}
```

The `main_test.go` file includes a pretty nifty testing harness: instead of writing your own testing code, you just register some test cases from `init()`.  Here's `minimal_test.go`:

```go
package main

func init() {
	addTestCases(minimalTests, minimal)
}

var minimalTests = []testCase{
{
	Name: "minimal.0",
	In: `package main
`
	Out: `package main
`
}
```

Test-driven development is the recommended workflow:

1. Add (or modify an existing) test case.
2. Run the tests, see the new one fail.
3. Modify the code until all tests pass.
4. Repeat until the tests verify the requirements have been met.

Publishing a Fix
----------------

Before publishing, double check the following:

- You have tested against other code bases to make sure the unit tests haven't missed anything.

Tips and Tricks
---------------

One of the subtleties is that if a package is imported with an alias, the alias must be used in the new type expression.  Conveniently, the files copied from Go's `src/cmd/fix` provide a convenience function `importSpec` that finds this alias.

```go
spec := importSpec(f, "<canonical path to mypkg>")
if spec == nil {
	return false
}
mypkg := "<un-aliased name of mypkg>"
if spec.Name != nil {
	mypkg = spec.Name.Name
}
```
