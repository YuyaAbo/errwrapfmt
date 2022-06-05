package errwrapfmt_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/YuyaAbo/errwrapfmt"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, errwrapfmt.Analyzer)
}
