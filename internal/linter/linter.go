package linter

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/almeidazs/gowther/internal/rules"
)

type Linter struct {
	Write  bool
	Unsafe bool
	Fset   *token.FileSet
	Config *rules.Config
}

func New(write, unsafe bool, config *rules.Config) *Linter {
	return &Linter{
		Write:  write,
		Unsafe: unsafe,
		Fset:   token.NewFileSet(),
		Config: config,
	}
}

func (l *Linter) ProcessPath(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				name := info.Name()
				if name == "vendor" || name == ".git" {
					return filepath.SkipDir
				}

				return nil
			}

			if strings.HasSuffix(p, ".go") && !strings.Contains(p, "vendor/") {
				return l.ProcessFile(p)
			}

			return nil
		})
	}

	return l.ProcessFile(path)
}

func (l *Linter) ProcessFile(filename string) error {
	src, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if _, err = parser.ParseFile(l.Fset, filename, src, parser.ParseComments); err != nil {
		return fmt.Errorf("parse error in %s: %v", filename, err)
	}

	return nil
}
