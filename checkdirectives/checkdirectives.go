package checkdirectives

import (
	"errors"
	"go/ast"
	"regexp"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "gocheckdirectives",
	Doc:  "Checks that go directive commenst (//go:) are valid.",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			comment, ok := n.(*ast.Comment)
			if !ok {
				return true
			}

			if isDirectiveIncorrectLeadingWhitespace(comment.Text) {
				pass.ReportRangef(comment, "comment contains directive")
			}

			return true
		})
	}
	return nil, errors.New("not implemented yet")
}

func isDirectiveIncorrectLeadingWhitespace(comment string) bool {
	matched, err := regexp.MatchString(`//\s+go:`, comment)
	if err != nil {
		panic("regex invalid: " + err.Error())
	}
	return matched
}

// func getNode(fset *token.FileSet, node any) string {
// 	s := strings.Builder{}
// 	_ = printer.Fprint(&s, fset, node)
// 	return s.String()
// }
