package main

import (
	"github.com/leighmcculloch/gocheckcompilerdirectives/checkcompilerdirectives"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(checkcompilerdirectives.Analyzer())
}
