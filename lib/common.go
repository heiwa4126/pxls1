package lib

import (
	"fmt"
	"sort"
	"strings"
)

// TailIs はsの末尾がsubstrであるときtrueを返す
func TailIs(s, substr string) bool {
	if len(s) < len(substr) {
		return false
	}
	return s[len(s)-len(substr):] == substr
}

// version.archを分解する
func VerArch(verarch string) (ver string, arch string, err error) {
	for _, arch = range []string{"x86_64", "i686", "noarch"} {
		if TailIs(verarch, "."+arch) {
			ver = verarch[0 : len(verarch)-len(arch)-1]
			return
		}
	}
	err = fmt.Errorf("Unknown arch. '%s'", verarch)
	return
}

// 文字列スライスの大文字小文字無視のソート
func isort(a []string) []string {
	sort.Slice(a, func(i, j int) bool { return strings.ToLower(a[i]) < strings.ToLower(a[j]) })
	return a
}

// 文字列スライスの大文字小文字無視のソート
func psort(a []Pkg) []Pkg {
	sort.Slice(a, func(i, j int) bool { return a[i].Compare(&a[j]) < 0 })
	return a
}
