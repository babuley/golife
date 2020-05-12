package main

import "testing"

func TestGrid(t *testing.T) {
	dim := &Config{60, 30, 100, 100}
	grid := makeGrid(dim)
	expectedSize := 60 * 30
	if expectedSize != len(grid) {
		t.Errorf("Expected grid of '%v' elements, but got '%v'", expectedSize, len(grid))
	}
}

func TestNewNeighbourhood(t *testing.T) {
	c := NewCell(10, 10, 1)
	ns := c.DefineNeighbours()
	if len(ns) != 8 {
		t.Errorf("Expected 8, but got '%v'", len(ns))
	}
}

func TestGenerators(t *testing.T) {
	dim := &Config{60, 30, 100, 100}
	gens := GetGenerators(dim)
	expected := dim.Height * dim.Width
	if len(gens) != expected {
		t.Errorf("Expected '%v', but got '%v'", expected, len(gens))
	}
}

func TestIsCellAlive(t *testing.T) {
	c := NewCell(10, 10, 1)
	if !c.IsAlive() {
		t.Errorf("Expected alive, but was dead")
	}
}

func TestNewNeighbour(t *testing.T) {
	n := NewNeighbourCell(11, 11)
	if n.IsAlive() {
		t.Errorf("Expected dead, but was alive")
	}
}
