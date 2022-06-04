package errwrapfmt

import (
	"go/ast"
	"regexp"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "errwrapfmt",
	Doc:  "errwrapfmt finds invalid error wrap format",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var (
	wrapReg  = regexp.MustCompile(`%w`)
	validReg = regexp.MustCompile(`: %w`)
)

func run(pass *analysis.Pass) (any, error) {
	in := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// target node is only string literal
	types := []ast.Node{
		(*ast.BasicLit)(nil),
	}

	in.Preorder(types, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.BasicLit:
			if wrapReg.MatchString(n.Value) && !validReg.MatchString(n.Value) {
				pass.Reportf(n.Pos(), "invalid error wrap format")
			}
		}
	})

	return nil, nil
}
