package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/alexandregv/gosane/noiter"
)

func main() {
	singlechecker.Main(noiter.Analyzer())
}
