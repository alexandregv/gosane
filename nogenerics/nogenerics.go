// Package nogenerics implements [go/analysis.Analyzer] and
// flags usage of generics.
package nogenerics

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "nogenerics",
		Doc:      "Forbids the use of generics (type parameters)",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (any, error) {
	insp, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, nil
	}

	insp.Preorder([]ast.Node{(*ast.FuncType)(nil), (*ast.TypeSpec)(nil)}, func(node ast.Node) {
		switch n := node.(type) {
		case *ast.FuncType:
			if n.TypeParams != nil && n.TypeParams.NumFields() > 0 {
				pass.Reportf(n.Pos(), "generics are not allowed (type parameters in function)")
			}
		case *ast.TypeSpec:
			if n.TypeParams != nil && n.TypeParams.NumFields() > 0 {
				pass.Reportf(n.Pos(), "generics are not allowed (type parameters in type declaration)")
			}
		}
	})

	return nil, nil
}
