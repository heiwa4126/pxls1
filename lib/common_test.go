package lib

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommon1(t *testing.T) {
	p1 := Pkg{"python-rhsm", "1.18.7-1.el6_9", "x86_64"}
	s1 := p1.String()
	assert.Equal(t, s1, "python-rhsm-1.18.7-1.el6_9.x86_64", "type Pkg BUG")

	p2 := Pkg{"redhat-access-insights", "1.0.13-2.el6", "noarch"}
	s2 := p2.String()
	assert.Equal(t, s2, "redhat-access-insights-1.0.13-2.el6.noarch", "type Pkg BUG")

	fmt.Printf("'%s','%s'\n", p1, p2)

	assert.True(t, strings.HasSuffix(s1, ".x86_64"), "strings.HasSuffix() BUG")
	assert.False(t, strings.HasSuffix(s1, ".noarch"), "strings.HasSuffix() BUG")

	assert.False(t, strings.HasSuffix(s2, ".x86_64"), "strings.HasSuffix() BUG")
	assert.True(t, strings.HasSuffix(s2, ".noarch"), "strings.HasSuffix() BUG")

	for _, p := range []Pkg{p1, p2} {
		namever, arch, err := VerArch(p.String())
		if err != nil {
			t.Fatal(err)
		}
		assert.True(t,
			namever == p.NameVer() && arch == p.Arch,
			"VerArch() BUG. version='%s', arch='%s'", namever, arch)
	}

}
