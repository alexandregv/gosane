// Package forbidigo wraps [forbidigo] with custom parameters for global-state functions.
//
// [forbidigo]: https://github.com/ashanbrown/forbidigo
package forbidigo

import (
	"fmt"

	"golang.org/x/tools/go/analysis"

	"github.com/ashanbrown/forbidigo/v2/forbidigo"
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "forbidigo",
		Doc:  "Forbids the use of specific identifiers",
		Run: makeRun([]string{
			`^http\.ListenAndServe`,

			`^viper\.Get`,
			`^viper\.Set`,
			`^viper\.Bind`,
		}),
	}
}

func makeRun(patterns []string) func(*analysis.Pass) (any, error) {
	return func(pass *analysis.Pass) (any, error) {
		linter, err := forbidigo.NewLinter(
			patterns,
			forbidigo.OptionAnalyzeTypes(true),
			forbidigo.OptionExcludeGodocExamples(true),
		)
		if err != nil {
			return nil, fmt.Errorf("forbidigo: %w", err)
		}

		config := forbidigo.RunConfig{
			Fset:      pass.Fset,
			TypesInfo: pass.TypesInfo,
		}

		for _, file := range pass.Files {
			issues, err := linter.RunWithConfig(config, file)
			if err != nil {
				return nil, fmt.Errorf("forbidigo: %w", err)
			}

			for _, issue := range issues {
				pass.Reportf(issue.Pos(), "%s", issue.Details())
			}
		}

		return nil, nil
	}
}
