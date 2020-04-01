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
	Width, Height int
	Sleep         time.Duration
}

//IsAlive returns true for active cell
func (c *Cell) IsAlive() bool {
	return c.Value != 0
}

//NewNeighbourCell - returns new dummy cell to the given coordinates
func NewNeighbourCell(x int, y int) Cell {
	return Cell{x, y, uuid.Nil, 0}
}

//NewCell - returns  set of neighbouring cells to the given cell
func NewCell(x int, y int, active int) *Cell {
	return &Cell{x, y, uuid.New(), active}
}
