package sdk

import "testing"

func TestAdd(t *testing.T) {
	var r int
	r = Add(1, 2)
	if r != 3 {
		t.Error("Add(1,2) failed . Go %d, expected 3.", r)
	} else {
		t.Log("test Add sucess")
	}
}

func TestAdd2(t *testing.T) {
	r := Add(3, 2)
	if r != 5 {
		t.Error("Add(1,2) failed . Go %d, expected 5.", r)
	} else {
		t.Log("test Add sucess")
	}
}
