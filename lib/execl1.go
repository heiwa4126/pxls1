package lib

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Excel1 struct {
	f *excelize.File
}

// a little bit wordy...

func NewExcel1() *Excel1 {
	return &Excel1{excelize.NewFile()}
}

func (e1 *Excel1) AddHost(host string, packages []Pkg) {
	e1.f.NewSheet(host)
	for i, v := range packages {
		e1.f.SetCellValue(host, fmt.Sprintf("A%d", i+1), v.Name)
		e1.f.SetCellValue(host, fmt.Sprintf("B%d", i+1), v.Version)
		e1.f.SetCellValue(host, fmt.Sprintf("C%d", i+1), v.Arch)
		e1.f.SetCellValue(host, fmt.Sprintf("D%d", i+1), v.String())
	}
}

func (e1 *Excel1) Finish(filename string) error {
	e1.f.DeleteSheet("Sheet1")
	return e1.f.SaveAs(filename)
}
