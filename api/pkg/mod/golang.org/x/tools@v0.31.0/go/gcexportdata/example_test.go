// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.7 && gc && !android && !ios && (unix || aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || plan9 || windows)

package gcexportdata_test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/gcexportdata"
)

// ExampleRead uses gcexportdata.Read to load type information for the
// "fmt" package from the fmt.a file produced by the gc compiler.
func ExampleRead() {
	// Find the export data file.
	filename, path := gcexportdata.Find("fmt", "")
	if filename == "" {
		log.Fatalf("can't find export data for fmt")
	}
	fmt.Printf("Package path:       %s\n", path)

	// Open and read the file.
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r, err := gcexportdata.NewReader(f)
	if err != nil {
		log.Fatalf("reading export data %s: %v", filename, err)
	}

	// Decode the export data.
	fset := token.NewFileSet()
	imports := make(map[string]*types.Package)
	pkg, err := gcexportdata.Read(r, fset, imports, path)
	if err != nil {
		log.Fatal(err)
	}

	// We can see all the names in Names.
	members := pkg.Scope().Names()
	foundPrintln := false
	for _, member := range members {
		if member == "Println" {
			foundPrintln = true
			break
		}
	}
	fmt.Print("Package members:    ")
	if foundPrintln {
		fmt.Println("Println found")
	} else {
		fmt.Println("Println not found")
	}

	// We can also look up a name directly using Lookup.
	println := pkg.Scope().Lookup("Println")
	// go 1.18+ uses the 'any' alias
	typ := strings.ReplaceAll(println.Type().String(), "interface{}", "any")
	fmt.Printf("Println type:       %s\n", typ)
	posn := fset.Position(println.Pos())
	// make example deterministic
	posn.Line = 123
	fmt.Printf("Println location:   %s\n", slashify(posn))

	// Output:
	//
	// Package path:       fmt
	// Package members:    Println found
	// Println type:       func(a ...any) (n int, err error)
	// Println location:   $GOROOT/src/fmt/print.go:123:1
}

// ExampleNewImporter demonstrates usage of NewImporter to provide type
// information for dependencies when type-checking Go source code.
func ExampleNewImporter() {
	const src = `package myrpc

// choosing a package that doesn't change across releases
import "net/rpc"

const serverError rpc.ServerError = ""
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "myrpc.go", src, 0)
	if err != nil {
		log.Fatal(err)
	}

	packages := make(map[string]*types.Package)
	imp := gcexportdata.NewImporter(fset, packages)
	conf := types.Config{Importer: imp}
	pkg, err := conf.Check("myrpc", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err)
	}

	// object from imported package
	pi := packages["net/rpc"].Scope().Lookup("ServerError")
	fmt.Printf("type %s.%s %s // %s\n",
		pi.Pkg().Path(),
		pi.Name(),
		pi.Type().Underlying(),
		slashify(fset.Position(pi.Pos())),
	)

	// object in source package
	twopi := pkg.Scope().Lookup("serverError")
	fmt.Printf("const %s %s = %s // %s\n",
		twopi.Name(),
		twopi.Type(),
		twopi.(*types.Const).Val(),
		slashify(fset.Position(twopi.Pos())),
	)

	// Output:
	//
	// type net/rpc.ServerError string // $GOROOT/src/net/rpc/client.go:20:1
	// const serverError net/rpc.ServerError = "" // myrpc.go:6:7
}

func slashify(posn token.Position) token.Position {
	posn.Filename = filepath.ToSlash(posn.Filename) // for MS Windows portability
	return posn
}
