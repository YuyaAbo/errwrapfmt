package errwrapfmt

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/YuyaAbo/errwrapfmt"
)

func main() {
	unitchecker.Main(errwrapfmt.Analyzer)
}
