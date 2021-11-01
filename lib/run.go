package lib

import (
	"fmt"

	"github.com/heiwa4126/pxls1/cli"
)

// Run は ....
func Run(jsonDir string, excelFile string) error {

	// get list
	hosts, err := Ls(jsonDir)
	if err != nil {
		return err
	}
	if len(hosts) == 0 {
		return fmt.Errorf("No JSON files")
	}

	// loop by hosts
	e1 := NewExcel1()

	for _, host := range hosts {
		pkgs, err2 := ReadJSON7(Host2File(host, jsonDir))
		if err2 != nil {
			return err2
		}
		timestamp, err2 := ReadLogFile(Host2LogFile(host, jsonDir))
		if err2 != nil {
			return err2
		}
		pkgdb, err2 := CsvRead(*cli.CsvPath + "/" + host + ".csv")
		if err2 != nil {
			return err2
		}

		e1.AddHost(host, pkgs, pkgdb, timestamp)
	}

	// write Excel file
	return e1.Finish(excelFile)
}
