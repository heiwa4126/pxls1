package main

import (
	"fmt"
	"log"
	"os"

	"github.com/heiwa4126/pxls1/cli"
	"github.com/heiwa4126/pxls1/lib"
)

func main() {

	var err error

	cli.ParseCli()

	if *cli.YamlFmt {
		// `-y` mode
		fmt.Fprintf(os.Stderr, "JSONPath=%s, YamlFile=%s\n", cli.JsonPath, cli.OutFile)
		err = lib.RunYAML(cli.JsonPath, cli.OutFile)
	} else {
		// normal mode
		fmt.Fprintf(os.Stderr, "JSONPath=%s, ExcelFile=%s\n", cli.JsonPath, cli.OutFile)
		err = lib.Run(cli.JsonPath, cli.OutFile)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stderr, "convert completed.")
}
