package goappendetect_test

import (
	"testing"

	"github.com/shoooooman/goappendetect"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, goappendetect.Analyzer, "a")
}

