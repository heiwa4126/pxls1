package lib

import (
	"encoding/csv"
	"io"
	"os"

	"log"
)

type PkgDBKey [2]string
type PkgDB map[PkgDBKey]string

func NewPkgDBKey(pkg, arch string) PkgDBKey {
	return [2]string{pkg, arch}
}
func (p PkgDB) Ver(pkg, arch string) string {
	return p[NewPkgDBKey(pkg, arch)]
}

/// CSVファイルを読んで、mapにして返す。
func CsvRead(csvfile string) (PkgDB, error) {
	m := make(PkgDB)
	fp, err := os.Open(csvfile)
	if err != nil {
		// ファイルがない場合
		// return nil, err
		log.Printf("CSVFile %s not found.\n", csvfile)
		return m, nil
	}
	defer fp.Close()

	r := csv.NewReader(fp)
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		// fmt.Printf("%#v\n", row)
		pkg := row[0]
		arch := row[3]
		ver := row[1] + "-" + row[2]

		// fmt.Printf("%s,%s,%s\n", pkg, arch, ver)
		m[NewPkgDBKey(pkg, arch)] = ver
	}

	return m, nil
}
