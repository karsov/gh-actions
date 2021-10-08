package main

import "testing"

func TestSum(t *testing.T) {
	res := Sum("1", "2")
	if res != "12" {
		t.Errorf(`Expected "12", received "%s"`, res)
	}
}
