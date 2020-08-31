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

func (e1 *Excel1) AddHost(host string, packages []string) {
	e1.f.NewSheet(host)
	for i, v := range packages {
		e1.f.SetCellValue(host, fmt.Sprintf("A%d", i+1), v)
	}
}

func (e1 *Excel1) Finish(filename string) error {
	e1.f.DeleteSheet("Sheet1")
	return e1.f.SaveAs(filename)
}
