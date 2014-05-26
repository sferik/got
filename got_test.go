package main

import "testing"

func TestRuler(t *testing.T) {
	for indent, expected := range map[int]string{
		0:  "----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|",
		5:  "     ----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|",
		-1: "----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|",
	} {
		got := ruler(indent)
		if got != expected {
			t.Errorf("with indent %d\n%8s: %s\n%8s: %s", indent, "expected", expected, "got", got)
		}
	}
}
