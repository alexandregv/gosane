package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/alexandregv/gosane/nogenerics"
)

func main() {
	singlechecker.Main(nogenerics.Analyzer())
}
