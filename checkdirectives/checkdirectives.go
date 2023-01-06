package checkdirectives

import (
	"regexp"

	"golang.org/x/tools/go/analysis"
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "gocheckdirectives",
		Doc:  "Checks that go directive commenst (//go:) are valid.",
		Run:  run,
	}
}

var reg = regexp.MustCompile(`^\s*(//(\s*)go:([a-z_]+))`)

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		for _, group := range file.Comments {
			for _, comment := range group.List {
				matches := reg.FindStringSubmatch(comment.Text)
				if len(matches) == 0 {
					continue
				}
				// Leading whitespace will cause the go directive to be ignored
				// by the compiler with no error, causing it not to work. This
				// is an easy mistake.
				if len(matches[2]) > 0 {
					pass.ReportRangef(comment, "go directive contains leading space: %s", matches[1])
				}
				// If the directive is unknown it will be ignored by the
				// compiler with no error. This is an easy mistake to make,
				// especially if you typo a directive.
				if !isKnown(matches[3]) {
					pass.ReportRangef(comment, "unrecognized go directive: %s", matches[1])
				}
			}
		}
	}
	return nil, nil
}

func isKnown(directive string) bool {
	for _, k := range known {
		if directive == k {
			return true
		}
	}
	return false
}

var known = []string{
	// Found by running the following command on the source of go.
	// git grep -o -E -h '//go:[a-z_]+' -- ':!**/*_test.go' ':!test/' ':!**/testdata/**' | sort -u
	"binary",
	"build",
	"buildsomethingelse",
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
