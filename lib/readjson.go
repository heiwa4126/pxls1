package lib

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"sort"
	"strings"
)

// ReadJSON は ....
func ReadJSON(jsonFile string) ([]string, error) {

	fp, err := os.Open(jsonFile)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	dec := json.NewDecoder(bufio.NewReader(fp))

	var m []string
	err = dec.Decode(&m)
	if err != nil && err != io.EOF {
		return nil, err
	}

	// 空だったらすぐ終わり
	if len(m) == 0 {
		return m, nil
	}

	j := 0
	// remove "Downloaded: "
	for _, v := range m {
		if strings.Index(v, "Downloaded: ") == 0 {
			m[j] = v[12:]
			j++
		}
	}
	m = m[0:j]

	// sort ignore-case
	sort.Slice(m, func(i, j int) bool { return strings.ToLower(m[i]) < strings.ToLower(m[j]) })

	// all over
	return m, nil
}
