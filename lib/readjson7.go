package lib

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"
)

// ReadJSON7 は ....
func ReadJSON7(jsonFile string) ([]string, error) {

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

	m := make([]string, len(s))
	j := 0
	for _, v := range s {
		x := strings.Index(v[1], " from ")
		if x > 0 {
			m[j] = v[0] + "-" + v[1][0:x]
			j++
		}
	}

	return isort(m[:j]), nil
}
