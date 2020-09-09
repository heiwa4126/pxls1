package lib

import (
	"testing"
)

func TestCommon1(t *testing.T) {
	p1 := Pkg{"python-rhsm", "1.18.7-1.el6_9", "x86_64"}
	s1 := p1.ToString()
	if s1 != "python-rhsm-1.18.7-1.el6_9.x86_64" {
		t.Fatalf("type Pkg BUG")
	}

	p2 := Pkg{"redhat-access-insights", "1.0.13-2.el6", "noarch"}
	s2 := p2.ToString()
	if s2 != "redhat-access-insights-1.0.13-2.el6.noarch" {
		t.Fatalf("type Pkg BUG")
	}

	if TailIs(s1, ".x86_64") != true {
		t.Fatalf("TailIs() BUG")
	}
	if TailIs(s1, ".noarch") == true {
		t.Fatalf("TailIs() BUG")
	}
	if TailIs(s2, ".x86_64") == true {
		t.Fatalf("TailIs() BUG")
	}
	if TailIs(s2, ".noarch") != true {
		t.Fatalf("TailIs() BUG")
	}

	var err error
	var ver, arch string

	ver, arch, err = VerArch(s1)
	if err != nil {
		t.Fatal(err)
	}
	if !(ver == p1.NameVer() && arch == p1.Arch) {
		t.Fatalf("VerArch() BUG. version='%s'.arch='%s'", ver, arch)
	}

	ver, arch, err = VerArch(s2)
	if err != nil {
		t.Fatal(err)
	}
	if !(ver == p2.NameVer() && arch == p2.Arch) {
		t.Fatalf("VerArch() BUG. version='%s'.arch='%s'", ver, arch)
	}
}
