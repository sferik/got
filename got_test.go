package main

import "testing"

func TestRuler(t *testing.T) {
	expected := "----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|"
	got := ruler(0)
	if got != expected {
		t.Errorf("\n%8s: %s\n%8s: %s", "expected", expected, "got", got)
	}
}
