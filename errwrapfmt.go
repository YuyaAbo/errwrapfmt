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
	Doc:  "finds invalid error wrap format",
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
	ins := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.BasicLit)(nil),
	}

	ins.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.BasicLit:
			if wrapReg.MatchString(n.Value) && !validReg.MatchString(n.Value) {
				pass.Reportf(n.Pos(), "invalid error wrap format")
			}
		}
	})

	return nil, nil
}
