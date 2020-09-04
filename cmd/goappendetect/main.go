package main

import (
	"github.com/shoooooman/goappendetect"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(goappendetect.Analyzer) }

