package lib

type Pkg struct {
	Name    string
	Version string
	Arch    string
}

func (p *Pkg) ToString() string {
	return p.Name + "-" + p.Version + "." + p.Arch
}
func (p *Pkg) NameVer() string {
	return p.Name + "-" + p.Version
}
