package lib

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Excel1 struct {
	f *excelize.File
}

// a little bit wordy...

func NewExcel1() *Excel1 {
	return &Excel1{excelize.NewFile()}
}

func (e1 *Excel1) AddHost(host string, packages []Pkg, pkgdb PkgDB, timestamp string) {
	e1.f.NewSheet(host)
	e1.f.SetCellValue(host, "A1", "パッケージ名")
	e1.f.SetCellValue(host, "B1", "アーキテクチャー")
	e1.f.SetCellValue(host, "C1", "現行バージョン")
	e1.f.SetCellValue(host, "D1", fmt.Sprintf("更新バージョン(%s時点)", timestamp))

	for i, v := range packages {
		j := i + 2
		ver := v.Version
		n := strings.Index(ver, ":")
		if n >= 0 {
			ver = ver[(n + 1):]
		}

		e1.f.SetCellValue(host, fmt.Sprintf("A%d", j), v.Name)
		e1.f.SetCellValue(host, fmt.Sprintf("B%d", j), v.Arch)
		e1.f.SetCellValue(host, fmt.Sprintf("C%d", j), pkgdb.Ver(v.Name, v.Arch))
		e1.f.SetCellValue(host, fmt.Sprintf("D%d", j), ver)
		// e1.f.SetCellValue(host, fmt.Sprintf("D%d", j), v.ToString())
	}
}

func (e1 *Excel1) Finish(filename string) error {
	e1.f.DeleteSheet("Sheet1")
	return e1.f.SaveAs(filename)
}
