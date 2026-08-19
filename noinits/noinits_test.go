package noinits

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestNoInits(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer(), "noinits")
}
