package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"4d63.com/gochecknoglobals/checknoglobals"
	"github.com/AlwxSin/noinlineerr"

	"github.com/alexandregv/gosane/nogenerics"
	"github.com/alexandregv/gosane/noiter"
)

func main() {
	multichecker.Main(
		noiter.Analyzer(),
		nogenerics.Analyzer(),
		noinlineerr.NewAnalyzer(),
		checknoglobals.Analyzer(),
	)
}
