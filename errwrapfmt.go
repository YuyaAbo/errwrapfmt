package errwrapfmt

import (
	"go/ast"

	"regexp"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "errwrapfmt",
	Doc:  "errwrapfmt finds invalid error wrap format",
	Run:  run,
}

var (
	wrapReg  = regexp.MustCompile(`%w`)
	validReg = regexp.MustCompile(`: %w`)
)

func run(pass *analysis.Pass) (any, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			// target node is only string literal
			bl, ok := n.(*ast.BasicLit)
			if !ok {
				return true
			}

			if wrapReg.MatchString(bl.Value) && !validReg.MatchString(bl.Value) {
				pass.Reportf(bl.Pos(), "invalid error wrap format")
			}

			return true
		})
	}
	return nil, nil
}
