package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/alexandregv/gosane/noshortif"
)

func main() {
	singlechecker.Main(noshortif.Analyzer())
}
