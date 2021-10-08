package main

import "testing"

func TestSum(t *testing.T) {
	res := Sum("1", "2")
	if res != "1+2" {
		t.Errorf(`Expected "1+2", received "%s"`, res)
	}
}
