package main

import "testing"

func TestTrue(t *testing.T) {
	if false {
		t.Error("Wrong result")
	}
}
