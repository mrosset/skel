package skel

import (
	"testing"
)

func TestPrint(t *testing.T) {
	if Print() != "HELLO" {
		t.Fail()
	}
}
