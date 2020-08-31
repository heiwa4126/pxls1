package main

import (
	"fmt"
	"log"
	"os"

	"github.com/heiwa4126/pxls1/lib"
)

var (
	// Version ...
	Version string = "v9.9.9"
	// Revision ...
	Revision string = "9999999"
)

func main() {

	jsonPath, outFile, ymlfmt := lib.ParseCli(Version, Revision)
	var err error

	if ymlfmt {
		// `-y` mode
		fmt.Fprintf(os.Stderr, "JSONPath=%s, ExcelFile=%s\n", jsonPath, outFile)
		err = lib.RunYAML(jsonPath, outFile)
	} else {
		// normal mode
		fmt.Fprintf(os.Stderr, "JSONPath=%s, ExcelFile=%s\n", jsonPath, outFile)
		err = lib.Run(jsonPath, outFile)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stderr, "convert completed.\n")
}
