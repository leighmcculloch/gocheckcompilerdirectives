package checkdirectives_test

import (
	"go/ast"
	"testing"

	"github.com/leighmcculloch/gocheckdirectives/checkdirectives"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestRun(t *testing.T) {
	testdata := analysistest.TestData()
	analyzer := checkdirectives.Analyzer()
	analysistest.Run(t, testdata, analyzer)
}

func FuzzRun(f *testing.F) {
	analyzer := checkdirectives.Analyzer()
	f.Add("hello world")
	f.Add("go:generate echo hello world")
	f.Add("go:embed")
	f.Add("go:embod")
	f.Add("  go:embod")
	f.Add("  go:embed")
	f.Fuzz(func(t *testing.T, comment string) {
		pass := analysis.Pass{
			Report: func(d analysis.Diagnostic) {},
			Files: []*ast.File{
				{Name: &ast.Ident{}, Comments: []*ast.CommentGroup{{List: []*ast.Comment{
					{Text: "//" + comment},
				}}}},
			},
		}
		_, err := analyzer.Run(&pass)
		if err != nil {
			t.Fatal(err)
		}
	})
}

func BenchmarkRun(b *testing.B) {
	analyzer := checkdirectives.Analyzer()
	pass := analysis.Pass{
		Report: func(d analysis.Diagnostic) {},
		Files: []*ast.File{
			{Name: &ast.Ident{}, Comments: []*ast.CommentGroup{{List: []*ast.Comment{
				{Text: "// some other comment"},
				{Text: "//go:generate echo hello world"},
				{Text: "// some other comment"},
				{Text: "//go:generate echo hello world"},
				{Text: "// some other comment"},
				{Text: "//go:generate echo hello world"},
				{Text: "// some other comment"},
				{Text: "//go:generate echo hello world"},
				{Text: "// some other comment"},
				{Text: "//go:generate echo hello world"},
				{Text: "// some other comment"},
			}}}},
		},
	}
	for i := 0; i < b.N; i++ {
		_, err := analyzer.Run(&pass)
		if err != nil {
			b.Fatal(err)
		}
	}
}
