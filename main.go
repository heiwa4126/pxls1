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

	// parse command line

	jsonPath := "./testdata/1"
	excelFile := "./Book1.xlsx"

	if len(os.Args) >= 2 {
		jsonPath = os.Args[1]
	}
	if len(os.Args) >= 3 {
		excelFile = os.Args[2]
	}

	switch jsonPath {
	case "-v":
		fmt.Printf("pxls1 %s (%s)\n", Version, Revision)
		os.Exit(2)
	case "-h":
		fmt.Printf("usage:\n\tpxls1 [-h|-v] <JSON files directory> <output Excel file>\n")
		os.Exit(2)
	}

	// main
	fmt.Fprintf(os.Stderr, "JSONPath=%s, ExcelFile=%s\n", jsonPath, excelFile)

	err := lib.Run(jsonPath, excelFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stderr, "completed.\n")
}
