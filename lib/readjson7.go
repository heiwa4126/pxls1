package lib

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

func readMainJson(jsonFile string) ([][2]string, error) {
	fp, err := os.Open(jsonFile)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	dec := json.NewDecoder(bufio.NewReader(fp))

	var s [][2]string
	err = dec.Decode(&s)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return s, nil
}

func readI686Json(i686File string) (map[string]int, error) {
	fp, err := os.Open(i686File)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	dec := json.NewDecoder(bufio.NewReader(fp))

	var s []string
	err = dec.Decode(&s)
	if err != nil && err != io.EOF {
		return nil, err
	}

	st := make(map[string]int)
	for _, v := range s {
		st[v] = 0
	}

	return st, nil
}

// ReadLogFile は 指定されたファイル(実はタイムスタンプが入ってるだけ)を読み込む
func ReadLogFile(logfile string) (string, error) {
	bs, err := os.ReadFile(logfile)
	if err != nil {
		// return "", err
		log.Printf("Log file %s not found.", logfile)
		return "unknown", nil
	}

	return string(bs), nil
}

// ReadJSON7 は ....
func ReadJSON7(jsonFile string) ([]Pkg, error) {

	s, err := readMainJson(jsonFile)
	if err != nil {
		return nil, err
	}

	s2, err := readI686Json(jsonFile[0:len(jsonFile)-5] + "_i686.json")
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%+v\n", s2)

	m := make([]Pkg, len(s)*2)
	j := 0
	for _, v := range s {
		// exclude slony1 package
		if strings.HasPrefix(v[0], "slony1") ||
			strings.HasPrefix(v[0], "postgresql94") {
			continue
		}

		x := strings.Index(v[1], " from ")
		if x > 0 {
			ver, arch, err := VerArch(v[1][0:x])
			if err != nil {
				return nil, err
			}
			p := Pkg{v[0], ver, arch}
			m[j] = p
			j++

			// i686 special
			_, ok := s2[p.Name]
			if ok && p.Arch == "x86_64" {
				m[j] = p
				m[j].Arch = "i686"
				j++
			}
		}
	}

	return psort(m[:j]), nil
}
