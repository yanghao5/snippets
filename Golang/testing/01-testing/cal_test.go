package main

import (
	"fmt"
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := AddUpper(10)
	if res != 55 {
		fmt.Printf("error")
		t.Fatalf("error %d", res)
	}
	t.Logf("ok")
}
