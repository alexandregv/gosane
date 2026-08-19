// Package noiter implements [go/analysis.Analyzer] and flags usage of iter.
package noiter

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "noiter",
		Doc:      "Forbids the use of iterators/generators (range-over-func and iter package)",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (any, error) {
	insp, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, nil
	}

	for _, file := range pass.Files {
		for _, imp := range file.Imports {
			path := imp.Path.Value
			if path == `"iter"` {
				pass.Reportf(imp.Pos(), "iter package is not allowed")
			}
		}
	}

	insp.Preorder([]ast.Node{(*ast.RangeStmt)(nil)}, func(node ast.Node) {
		rangeStmt := node.(*ast.RangeStmt)
		typ := pass.TypesInfo.TypeOf(rangeStmt.X)
		if typ != nil {
			_, ok := typ.Underlying().(*types.Signature)
			if ok {
				pass.Reportf(rangeStmt.Pos(), "range-over-func is not allowed")
			}
		}
	})

	return nil, nil
}
