package swaggervalidatecheck

import (
	"golang.org/x/tools/go/analysis"
)

func New() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "swaggervalidate",
		Doc:  "check handler functions for call swagger body validate",
		Run:  run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	return nil, nil
}
