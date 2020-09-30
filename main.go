package main

import (
	"fmt"
	"log"
	"os"

	"github.com/heiwa4126/pxls1/cli"
	"github.com/heiwa4126/pxls1/lib"
)

func run() error {
	if *cli.YamlFmt {
		// `-y` mode
		fmt.Fprintf(os.Stderr, "JSONPath=%s, YamlFile=%s\n", cli.JsonPath, cli.OutFile)
		return lib.RunYAML(cli.JsonPath, cli.OutFile)
	}
	// normal mode
	fmt.Fprintf(os.Stderr, "JSONPath=%s, ExcelFile=%s\n", cli.JsonPath, cli.OutFile)
	return lib.Run(cli.JsonPath, cli.OutFile)
}

func main() {

	cli.ParseCli()

	if err := run(); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stderr, "*** convert completed.")
}
