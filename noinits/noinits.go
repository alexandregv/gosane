// Package noinits reimplements gochecknoinits but using the go/analysis standard.
// golangci-lint also [reimplements it but doesn't export it].
//
// [reimplements it but doesn't export it]: https://github.com/golangci/golangci-lint/blob/main/pkg/golinters/gochecknoinits/gochecknoinits.go
package noinits

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "gochecknoinits", // Use the same name for compatibility, like nolint:gochecknoinits directives
		Doc:      "Checks that no init functions are present in Go code",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (any, error) {
	insp, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, nil
	}

	insp.Preorder([]ast.Node{(*ast.FuncDecl)(nil)}, func(node ast.Node) {
		funcDecl := node.(*ast.FuncDecl)
		if funcDecl.Name.Name == "init" && funcDecl.Recv == nil {
			pass.Reportf(funcDecl.Pos(), "init functions are not allowed")
		}
	})

	return nil, nil
}
