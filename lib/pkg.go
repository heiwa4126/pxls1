package lib

import (
	"fmt"
	"sort"
	"strings"
)

type Pkg struct {
	Name    string
	Version string
	Arch    string
}

const (
	ARCH_x86    = "x86_64"
	ARCH_i686   = "i686"
	ARCH_noarch = "noarch"
)

// String() for "%s"
func (p Pkg) String() string {
	return p.ToString()
}

func (p *Pkg) ToString() string {
	return p.Name + "-" + p.Version + "." + p.Arch
}

func (p *Pkg) NameVer() string {
	return p.Name + "-" + p.Version
}

func (p *Pkg) Compare(p2 *Pkg) int {
	if rc := strings.Compare(p.Name, p2.Name); rc != 0 {
		return rc
	}
	if rc := strings.Compare(p.Version, p2.Version); rc != 0 {
		return rc
	}
	return strings.Compare(p.Arch, p2.Arch)
}

func PkgSlice2StringSlice(ps []Pkg) []string {

	ss := make([]string, len(ps))

	for i, v := range ps {
		ss[i] = v.ToString()
	}

	return ss
}

// version.archを分解する
func VerArch(verarch string) (ver string, arch string, err error) {
	for _, arch = range []string{ARCH_x86, ARCH_i686, ARCH_noarch} {
		if strings.HasSuffix(verarch, "."+arch) {
			ver = verarch[0 : len(verarch)-len(arch)-1]
			return
		}
	}
	err = fmt.Errorf("Unknown arch. '%s'", verarch)
	return
}

// Pkgスライスのソート(大文字小文字を見る)
func psort(a []Pkg) []Pkg {
	sort.Slice(a, func(i, j int) bool { return a[i].Compare(&a[j]) < 0 })
	return a
}
