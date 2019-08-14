package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/wtetsu/sillycd/pkg/sillycd"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Wrong argument")
		os.Exit(1)
	}

	var specifiedName = os.Args[1]
	var target string

	if filepath.IsAbs(specifiedName) {
		target = specifiedName
	} else {
		target = "./" + specifiedName
	}

	foundDirectory := sillycd.PickOutDirectory(target)

	if foundDirectory == "" {
		os.Exit(1)
	}
	fmt.Println(foundDirectory)
}
