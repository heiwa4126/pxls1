package lib

import (
	"fmt"
	"log"
	"testing"
)

func TestReadJSON7_1(t *testing.T) {
	rc, err := ReadJSON7("../test/7/web02.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("rc: %+v\n", rc)

	// t.Fatal("test")
}
