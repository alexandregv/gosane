package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/alexandregv/gosane/nogenerics"
	"github.com/alexandregv/gosane/noiter"
)

func main() {
	multichecker.Main(
		noiter.Analyzer(),
		nogenerics.Analyzer(),
	)
}
