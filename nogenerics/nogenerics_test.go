package nogenerics

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestNoGenerics(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer(), "nogenerics")
}
