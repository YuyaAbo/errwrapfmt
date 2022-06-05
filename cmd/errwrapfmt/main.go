package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/YuyaAbo/errwrapfmt"
)

func main() { singlechecker.Main(errwrapfmt.Analyzer) }
