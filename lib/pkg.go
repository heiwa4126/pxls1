package lib

import "strings"

type Pkg struct {
	Name    string
	Version string
	Arch    string
}

func (p Pkg) String() string {
	return p.Name + "-" + p.Version + "." + p.Arch
}

func (p Pkg) NameVer() string {
	return p.Name + "-" + p.Version
}

func (p Pkg) Compare(p2 *Pkg) int {
	rc := strings.Compare(p.Name, p2.Name)
	if rc != 0 {
		return rc
	}
	rc = strings.Compare(p.Version, p2.Version)
	if rc != 0 {
		return rc
	}
	return strings.Compare(p.Arch, p2.Arch)
}

func PkgSlice2StringSlice(ps []Pkg) []string {

	ss := make([]string, len(ps))

	for i, v := range ps {
		ss[i] = v.String()
	}

	return ss
}
