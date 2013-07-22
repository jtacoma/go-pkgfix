// +build samples

package main

import (
	"go/ast"
)

func init() {
	register(fiximportFix)
}

var fiximportFix = fix{
	"fiximport",
	"2013-07-21", // when this fix was published
	fiximport,
	`
    Fix import statements to match the new canonical import path.
`,
}

func fiximport(f *ast.File) bool {
	spec := importSpec(f, "example.com/mypkg")
	if spec == nil {
		return false
	}
	spec.Path.Value = "\"example.com/mypkg2\""
	return true
}
