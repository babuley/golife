package main

import (
	"time"

	"github.com/google/uuid"
)

type Cell struct {
	X, Y  int
	ID    uuid.UUID
	Value int
}

type Config struct {
	Width, Height       int
	Sleep               time.Duration
	NumberOfGenerations int
}

//IsAlive returns true for active cell
func (c *Cell) IsAlive() bool {
	return c.Value != 0
}

//NewNeighbourCell - returns new dummy cell to the given coordinates
func NewNeighbourCell(x int, y int) *Cell {
	return &Cell{x, y, uuid.Nil, 0}
}

//NewCell - returns  set of neighbouring cells to the given cell
func NewCell(x int, y int, active int) *Cell {
	return &Cell{x, y, uuid.New(), active}
}

//DefineNeighbours -- returns []Cell definitions of locations of neighbours for the given cell
func (c *Cell) DefineNeighbours() []*Cell {
	var neighbours []*Cell
	for j := -1; j < 2; j++ {
		for i := -1; i < 2; i++ {
			if j == 0 && i == 0 {
				//Exclude myself
				continue
			}
			neighbours = append(neighbours, NewNeighbourCell(i, j))
		}
	}
	return neighbours
}
