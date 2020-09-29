package cli

// parse commandline

import (
	"flag"
	"fmt"
	"os"
)

var (
	// Version ...
	Version string = "v9.9.9"
	// Revision ...
	Revision string = "9999999"

	JsonPath string = "./testdata/7"

	OutFile string = "./Book1.xlsx"

	YamlFmt *bool
)

func ParseCli() {

	// flags
	help := flag.Bool("h", false, "ヘルプの表示")
	version := flag.Bool("v", false, "バージョンの表示")
	YamlFmt = flag.Bool("y", false, "YAMLモード")

	flag.Usage = func() {
		fmt.Println(`JSONの変換
Usage:
 pxls1 <JSON files directory> <output Excel file(xlsx)>
 pxls1 -y <JSON files directory> <output update db file (YAML)>
 pxls1 [flags]

flags:`)
		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()

	if *version {
		fmt.Println("pxls1 " + Version + " (" + Revision + ")")
		os.Exit(2)
	}

	if *help {
		flag.Usage()
		os.Exit(2)
	}

	if *YamlFmt {
		OutFile = "./updates_db.yaml"
	}

	if len(args) >= 1 {
		JsonPath = args[0]
	}
	if len(args) >= 2 {
		OutFile = args[1]
	}
}
