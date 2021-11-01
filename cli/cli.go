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

	JsonPath string = "./test/7"
	OutFile  string = "./Book1.xlsx"
	YamlFmt  *bool
	CsvPath  *string
)

func ParseCli() {

	// flags
	help := flag.Bool("h", false, "ヘルプの表示")
	version := flag.Bool("v", false, "バージョンの表示")
	YamlFmt = flag.Bool("y", false, "YAMLモード")
	CsvPath = flag.String("r", "/home/heiwa/works/ansible/t2/var/rpmqa_csv", "rpmqa_csvのパス")

	flag.Usage = func() {
		fmt.Println(`JSONの変換
Usage:
 pxls1 <JSON files directory> <output Excel file(xlsx) -r <rpmqa CSVPATH>>
 pxls1 -y <JSON files directory> <output update db file (YAML)>
 pxls1 [flags]

flags:`)
		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()

	if *version {
		fmt.Println("pxls1 v" + Version + " (" + Revision + ")")
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
