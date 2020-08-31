package lib

import (
	"flag"
	"fmt"
	"os"
)

func ParseCli(ver string, rev string) (string, string, bool) {

	// flags
	help := flag.Bool("h", false, "ヘルプの表示")
	version := flag.Bool("v", false, "バージョンの表示")
	yamlfmt := flag.Bool("y", false, "YAMLモード")

	flag.Usage = func() {
		fmt.Println(`JSONの変換
Usage:
 pxls1 [JSON files directory] [output Excel file]
 pxls1 -y [JSON files directory] [output update db file (YAML)]
 pxls1 [flags]

flags:`)
		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()

	if *version {
		fmt.Printf("pxls1 %s (%s)\n", ver, rev)
		os.Exit(2)
	}

	if *help {
		flag.Usage()
		os.Exit(2)
	}

	jsonPath := "./testdata/1"
	outFile := "./Book1.xlsx"
	if *yamlfmt {
		outFile = "./updates_db.yaml"
	}

	if len(args) >= 1 {
		jsonPath = args[0]
	}
	if len(args) >= 2 {
		outFile = args[1]
	}

	return jsonPath, outFile, *yamlfmt
}
