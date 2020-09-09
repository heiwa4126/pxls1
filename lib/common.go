package lib

import "fmt"

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
