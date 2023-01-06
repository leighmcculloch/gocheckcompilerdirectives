package checkdirectives

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestCheckNoGlobals(t *testing.T) {
	testdata := analysistest.TestData()
	analyzer := Analyzer()
	analysistest.Run(t, testdata, analyzer)
}
