// Gosane lints Go programs. It merges multiple linters:
// nogenerics, noinits, noiter, noshortif, noinlineerr,
// gochecknoglobals and forbidigo.
package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"4d63.com/gochecknoglobals/checknoglobals"
	"github.com/AlwxSin/noinlineerr"

	"github.com/alexandregv/gosane/internal/forbidigo"
	"github.com/alexandregv/gosane/nogenerics"
	"github.com/alexandregv/gosane/noinits"
	"github.com/alexandregv/gosane/noiter"
	"github.com/alexandregv/gosane/noshortif"
)

func main() {
	multichecker.Main(
		noiter.Analyzer(),
		nogenerics.Analyzer(),
		noinlineerr.NewAnalyzer(),
		noinits.Analyzer(),
		noshortif.Analyzer(),
		checknoglobals.Analyzer(),
		forbidigo.Analyzer(),
	)
}
