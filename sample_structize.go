// +build samples

package main

import (
	"go/ast"
)

func init() {
	register(mypkgstructizeFix)
}

var mypkgstructizeFix = fix{
	"mypkgstructize",
	"2013-07-21",
	mypkgstructize,
	`
	Use structs instead of interfaces.
`,
}

var mypkgstructized = map[string]bool{
	"Widget": true,
}

func mypkgstructize(f *ast.File) bool {
	spec := importSpec(f, "example.com/mypkg")
	if spec == nil {
		return false
	}
	mypkg := "mypkg"
	if spec.Name != nil {
		mypkg = spec.Name.Name
	}

	fixed := false
	walk(f, func(n interface{}) {
		switch node := n.(type) {
		case *ast.ArrayType:
			t := mypkgstructizetype(mypkg, node.Elt)
			if t != nil {
				node.Elt = t
				fixed = true
			}
		case *ast.CompositeLit:
			// This is irrelevant only because the original type is an
			// interface i.e. cannot be the type of a composite literal.
		case *ast.Ellipsis:
			t := mypkgstructizetype(mypkg, node.Elt)
			if t != nil {
				node.Elt = t
				fixed = true
			}
		case *ast.Field:
			t := mypkgstructizetype(mypkg, node.Type)
			if t != nil {
				node.Type = t
				fixed = true
			}
		case *ast.MapType:
			t := mypkgstructizetype(mypkg, node.Key)
			if t != nil {
				node.Key = t
				fixed = true
			}
			t = mypkgstructizetype(mypkg, node.Value)
			if t != nil {
				node.Value = t
				fixed = true
			}
		case *ast.Object:
			// Does something need to be done here with node.Type?
			// What does it take to trigger this case?
		case *ast.TypeAssertExpr:
			t := mypkgstructizetype(mypkg, node.Type)
			if t != nil {
				node.Type = t
				fixed = true
			}
		case *ast.TypeSpec:
			t := mypkgstructizetype(mypkg, node.Type)
			if t != nil {
				node.Type = t
				fixed = true
			}
		case *ast.ValueSpec:
			t := mypkgstructizetype(mypkg, node.Type)
			if t != nil {
				node.Type = t
				fixed = true
			}
		}
	})
	return fixed
}

func mypkgstructizetype(mypkg string, n ast.Expr) ast.Expr {
	s, ok := n.(*ast.SelectorExpr)
	if ok {
		p, ok := s.X.(*ast.Ident)
		if ok && p.Name == mypkg {
			if mypkgstructized[s.Sel.Name] {
				return &ast.StarExpr{
					X: &ast.SelectorExpr{
						X:   ast.NewIdent(mypkg),
						Sel: ast.NewIdent(s.Sel.Name),
					},
				}
			}
		}
	}
	return nil
}
