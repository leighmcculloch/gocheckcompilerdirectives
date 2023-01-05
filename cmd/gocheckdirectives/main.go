package main

import (
	"github.com/leighmcculloch/gocheckdirectives/checkdirectives"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(checkdirectives.Analyzer)
}
