// goimportgraph - generate a `.dot` file from an import graph of a package
//
// Usage
//
//     goimportgraph <go package path> 
//

package main

import (
	"flag"
	"fmt"
	"go/build"

	"golang.org/x/tools/refactor/importgraph"
)

func printGraph(pkg string) error {
	// TODO: pkg is unused
	graph, _, errs := importgraph.Build(&build.Default)
	for _, err := range errs {
		// TODO: print all errors
		if err != nil {
			return err
		}
	}

	fmt.Println("digraph {")
	
	for fromPkg, imports := range graph {
		// TODO: filter out packages
		for toPkg := range imports {
			fmt.Printf("  %q -> %q;\n", fromPkg, toPkg)
		}
	}

	fmt.Println("}")
	return nil
}


func main() {
	flag.Parse()
	for _, pkg := range flag.Args() {
		// TODO: handle errors
		printGraph(pkg)
	}
}
