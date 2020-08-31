package lib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLsErr1(t *testing.T) {
	rc, err := Ls("../testdata/notexists")
	if err != nil {
		// 存在しないパスでもエラーにはならない
		t.Fatal(err)
	}

	if len(rc) != 0 {
		t.Fatalf("wrong length")
	}
}

func TestLs1(t *testing.T) {
	wants := []string{"c7", "host1", "R8"}
	rc, err := Ls("../testdata/1")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("wants: %+v\n   rc: %+v\n", wants, rc)

	if !reflect.DeepEqual(rc, wants) {
		t.Fatalf("wrong order")
	}
}
