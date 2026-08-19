package noshortif

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestNoShortIf(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer(), "noshortif")
}
