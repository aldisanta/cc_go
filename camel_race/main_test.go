package camel

import "testing"

func TestTew(t *testing.T) {
	expect := 0
	onta := 3
	roy := []int{
		3, 2, 1,
	}
	ananto := []int{
		1, 2, 3,
	}
	abi := []int{
		1, 2, 3,
	}

	if sum, _ := Tew(onta, roy, ananto, abi); expect != sum {
		t.Errorf("Expected %d, got %d", expect, sum)
	}
}

func TestTetew(t *testing.T) {
	expect := 3
	onta := 4
	roy := []int{
		2, 3, 1, 4,
	}
	ananto := []int{
		2, 1, 4, 3,
	}
	abi := []int{
		2, 4, 3, 1,
	}

	if sum, _ := Tew(onta, roy, ananto, abi); expect != sum {
		t.Errorf("Expected %d, got %d", expect, sum)
	}
}

func TestTetewtew(t *testing.T) {
	onta := 4
	// error here
	roy := []int{
		2, 3, 1,
	}
	ananto := []int{
		2, 1, 4, 3,
	}
	abi := []int{
		2, 4, 3, 1,
	}

	if _, err := Tew(onta, roy, ananto, abi); err == nil {
		t.Error("Expecting error")
	}
}

func TestTetewtewtew(t *testing.T) {
	// error here
	onta := 1
	roy := []int{
		2, 3, 1,
	}
	ananto := []int{
		2, 1, 4, 3,
	}
	abi := []int{
		2, 4, 3, 1,
	}

	if _, err := Tew(onta, roy, ananto, abi); err == nil {
		t.Error("Expecting error")
	}
}
