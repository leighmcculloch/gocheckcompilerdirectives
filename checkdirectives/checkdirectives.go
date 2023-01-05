package checkdirectives

import (
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
		for _, group := range file.Comments {
			for _, comment := range group.List {
				if isIncorrectLeadingWhitespace(comment.Text) {
					pass.ReportRangef(comment, "go directive contains leading space: %q", comment.Text)
				}
				if isUnrecognized(comment.Text) {
					pass.ReportRangef(comment, "unrecognized go directive: %q", comment.Text)
				}
			}
		}
	}
	return nil, nil
}

func isIncorrectLeadingWhitespace(comment string) bool {
	matched, err := regexp.MatchString(`//\s+go:`, comment)
	if err != nil {
		panic("regex invalid: " + err.Error())
	}
	return matched
}

func isUnrecognized(comment string) bool {
	r := regexp.MustCompile(`//\s*go:([a-z_]+)`)
	matches := r.FindStringSubmatch(comment)
	if len(matches) == 0 {
		return false
	}
	for _, k := range known {
		if matches[1] == k {
			return false
		}
	}
	return true
}

var known = []string{
	// Found by running the following command on the source of go.
	// git grep -o -E -h '//go:[a-z_]+' -- ':!**/*_test.go' ':!test/' ':!**/testdata/**' | sort -u
	"binary",
	"build",
	"buildsomethingelse",
	"cgo_",
	"cgo_dynamic_linker",
	"cgo_export_dynamic",
	"cgo_export_static",
	"cgo_import_dynamic",
	"cgo_import_static",
	"cgo_ldflag",
	"cgo_unsafe_args",
	"embed",
	"generate",
	"linkname",
	"name",
	"nocheckptr",
	"noescape",
	"noinline",
	"nointerface",
	"norace",
	"nosplit",
	"notinheap",
	"nowritebarrier",
	"nowritebarrierrec",
	"systemstack",
	"uintptrescapes",
	"uintptrkeepalive",
	"yeswritebarrierrec",
}
