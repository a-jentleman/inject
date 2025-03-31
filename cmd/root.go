package main

import (
	"fmt"
	"go/ast"
	"go/types"

	"github.com/spf13/cobra"
	"golang.org/x/tools/go/packages"
)

const pkgPath = "github.com/a-jentleman/inject"

var rootCmd = &cobra.Command{
	Use: "inject",
	RunE: func(cmd *cobra.Command, args []string) error {
		pkgs, err := packages.Load(&packages.Config{Mode: packages.LoadAllSyntax, Context: cmd.Context()}, args...)
		processPackages(pkgs)
		return err
	},
}

func processPackages(pkgs []*packages.Package) error {
	for _, pkg := range pkgs {
		err := processPackage(pkg)
		if err != nil {
			return err
		}
	}
	return nil
}

func processPackage(pkg *packages.Package) error {
	for _, file := range pkg.Syntax {
		err := processFile(pkg, file)
		if err != nil {
			return err
		}
	}
	return nil
}

func processFile(pkg *packages.Package, file *ast.File) error {
	importSpec := findImportSpec(file, pkgPath)
	if importSpec == nil {
		return nil
	}

	injectPkg, ok := pkg.Imports[pkgPath]
	if !ok {
		panic("inject package not found")
	}

	_, f := findFuncDefObject(injectPkg, "Inject")
	if f == nil {
		panic("inject.Inject not found")
	}

	return nil
}

func findFuncDefObject(pkg *packages.Package, name string) (*ast.Ident, *types.Func) {
	for ident, object := range pkg.TypesInfo.Defs {
		if ident.Name != name {
			continue
		}

		return ident, object.(*types.Func)
	}

	return nil, nil
}

func findImportSpec(file *ast.File, path string) *ast.ImportSpec {
	v := &findPkgVisitor{path: fmt.Sprintf("%q", path)}
	ast.Walk(v, file)
	return v.spec
}

type findPkgVisitor struct {
	path string
	spec *ast.ImportSpec
}

func (f *findPkgVisitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return nil
	}

	if f.spec != nil {
		return nil
	}

	s, ok := node.(*ast.ImportSpec)
	if !ok {
		return f
	}

	if s.Path.Value != f.path {
		return f
	}

	f.spec = s
	return nil
}
