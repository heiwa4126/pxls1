package lib

import (
	"sort"
	"strings"
)

// 文字列スライスの大文字小文字無視のソート
func isort(a []string) []string {
	sort.Slice(a, func(i, j int) bool { return strings.ToLower(a[i]) < strings.ToLower(a[j]) })
	return a
}
