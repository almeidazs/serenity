package bestpractices

import (
	"go/ast"

	"github.com/serenitysz/serenity/internal/rules"
)

type NoDeferInLoopRule struct{}

func (d *NoDeferInLoopRule) Name() string {
	return "no-defer-in-loop"
}

func (d *NoDeferInLoopRule) Targets() []ast.Node {
	return []ast.Node{
		(*ast.ForStmt)(nil),   // tradicional loops
		(*ast.RangeStmt)(nil), // loops into slices, maps, channels
	}
}

func (d *NoDeferInLoopRule) Run(runner *rules.Runner, node ast.Node) {
	if runner.ShouldStop != nil && runner.ShouldStop() {
		return
	}

	bp := runner.Cfg.Linter.Rules.BestPractices
	if bp == nil || (bp.Use != nil && !*bp.Use) || bp.NoDeferInLoop == nil {
		return
	}

	var body *ast.BlockStmt

	switch n := node.(type) {
	case *ast.RangeStmt:
		body = n.Body

	case *ast.ForStmt:
		body = n.Body
	}

	if body == nil {
		return
	}

	ast.Inspect(body, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.FuncLit:
			return false

		case *ast.ForStmt, *ast.RangeStmt:
			return false
		case *ast.DeferStmt:

			maxIssues := rules.GetMaxIssues(runner.Cfg)
			if maxIssues > 0 && maxIssues >= int16(len(*runner.Issues)) {
				return false
			}

			*runner.Issues = append(*runner.Issues, rules.Issue{
				ID:       rules.NoDeferInLoopID,
				Pos:      runner.Fset.Position(t.Pos()),
				Severity: rules.ParseSeverity(bp.NoDeferInLoop.Severity),
			})
		}

		return true
	})
}
