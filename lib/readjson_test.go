package lib

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestReadJSON1(t *testing.T) {
	wants := []string{
		"package01",
		"Package02",
		"package03",
		"Package04",
		"package05",
		"PacKage06",
		"package07",
	}

	rc, err := ReadJSON("../testdata/1/host1.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("wants: %+v\n   rc: %+v\n", wants, rc)

	if !reflect.DeepEqual(rc, wants) {
		t.Fatalf("wrong order")
	}
}

func TestReadJSON2(t *testing.T) {
	rc, err := ReadJSON("../testdata/1/c7.json")
	if err != nil {
		log.Fatal(err)
	}

	if len(rc) != 0 {
		t.Fatalf("filtering failed")
	}
}
