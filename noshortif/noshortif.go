// Package noshortif implements [go/analysis.Analyzer] and forbids the short-if syntax.
package noshortif

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "noshortif",
		Doc:      "Checks that the short-if syntax is not used",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (any, error) {
	insp, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, nil
	}

	insp.Preorder([]ast.Node{(*ast.IfStmt)(nil)}, func(node ast.Node) {
		ifStmt := node.(*ast.IfStmt)

		assignStmt, ok := ifStmt.Init.(*ast.AssignStmt)
		if ok {
			pass.Reportf(assignStmt.Pos(), "short-if syntax is not allowed")
		}
	})

	return nil, nil
}
