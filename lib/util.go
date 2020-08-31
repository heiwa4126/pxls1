package lib

import (
	"sort"
	"strings"
)

func isort(a []string) []string {
	sort.Slice(a, func(i, j int) bool { return strings.ToLower(a[i]) < strings.ToLower(a[j]) })
	return a
}
