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
	"os"
	"strings"

	"golang.org/x/tools/refactor/importgraph"
)

var (
	// TODO: join paths properly
	VENDORS = []string{"/vendor", "/GoDeps"}
)

func printGraph(pkgs []string) error {
	graph, _, errs := importgraph.Build(&build.Default)
	for _, err := range errs {
		fmt.Fprintln(os.Stderr, err)
		return fmt.Errorf("Failed to build import graph.")
	}

	fmt.Println("digraph {")
	
	for fromPkg, imports := range graph {
		if !includePkg(fromPkg, pkgs) {
			continue
		}
		for toPkg := range imports {
			fmt.Printf("  %q -> %q;\n", fromPkg, toPkg)
		}
	}

	fmt.Println("}")
	return nil
}

func isVendoredPkg(pkg string, include string) bool {
	for _, vendor := range VENDORS {
		if strings.HasPrefix(pkg, include + vendor) {
			return true
		}
	}
	return false
}

func includePkg(pkg string, includes []string) bool {
	for _, include := range includes {
		if strings.HasPrefix(pkg, include) && !isVendoredPkg(pkg, include) {
			return true
		}
	}
	return false
}

func main() {
	flag.Parse()
	err := printGraph(flag.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
}
