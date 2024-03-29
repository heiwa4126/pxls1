package lib

import (
	"path/filepath"
	"strings"
)

func stem(path1 string) string {
	base := filepath.Base(path1)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

func Host2File(host string, base string) string {
	return filepath.Join(base, host+".json")
}

func Host2LogFile(host string, base string) string {
	return filepath.Join(base, host+".log")
}

// Ls は ....
func Ls(searchPath string) ([]string, error) {

	rc, err := filepath.Glob(searchPath + "/*.json")
	if err != nil {
		// I/Oエラーのみerrになる
		return nil, err
	}

	hosts := make([]string, len(rc))

	j := 0
	for _, v := range rc {
		if !strings.HasSuffix(v, "_i686.json") {
			hosts[j] = stem(v)
			j++
		}
	}

	return isort(hosts[:j]), nil
}
