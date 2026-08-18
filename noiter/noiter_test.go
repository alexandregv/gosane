package noiter

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestNoIter(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer(), "noiter")
}
