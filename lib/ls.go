package lib

import (
	"path/filepath"
	"sort"
	"strings"
)

func stem(path1 string) string {
	base := filepath.Base(path1)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

func Host2File(host string, base string) string {
	return filepath.Join(base, host+".json")
}

// Ls は ....
func Ls(searchPath string) ([]string, error) {

	rc, err := filepath.Glob(searchPath + "/*.json")
	if err != nil {
		// I/Oエラーのみerrになる
		return nil, err
	}

	hosts := make([]string, len(rc))

	for i, v := range rc {
		hosts[i] = stem(v)
	}

	sort.Slice(hosts, func(i, j int) bool { return strings.ToLower(hosts[i]) < strings.ToLower(hosts[j]) })

	return hosts, nil
}
