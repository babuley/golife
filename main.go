package main

import (
	"fmt"
	"time"

	"github.com/danicat/simpleansi"
)

func main() {
	Initialise()
	defer Cleanup()

	//Dimension of the grid
	var dim = &Config{60, 30, 100}
	grid := makeGrid(dim)
	for {
		// repeat
		dumpGrid(grid)
		runTick(grid)
		time.Sleep(dim.Sleep * time.Millisecond)
	}
}

func makeGrid(config *Config) []*Cell {
	var grid []*Cell
	for y := 0; y < config.Height; y++ {
		for x := 0; x < config.Width; x++ {
			alive := 0
			if x > 0 && x < config.Width-1 && y%3 == 0 && y > 3 {
				alive = 1
			}
			grid = append(grid, NewCell(y, x, alive))
		}
	}
	return grid
}

func getLiveNeighbours(cands []*Cell) []*Cell {
	var alive []*Cell
	for _, c := range cands {
		if c.IsAlive() {
			alive = append(alive, c)
		}
	}
	return alive
}

//-Any live cell with two or three neighbors survives.
//-Any dead cell with three live neighbors becomes a live cell.
//-All other live cells die in the next generation. Similarly, all other dead cells stay dead.
func runTick(grid []*Cell) {
	var toRevive []*Cell
	var toKill []*Cell

	for _, cand := range grid {
		ns := getNeighbours(cand, grid)
		aliveNeighbours := getLiveNeighbours(ns)

		if cand.IsAlive() {
			var twoOrThreeLiveNeighbours = len(aliveNeighbours) == 2 || len(aliveNeighbours) == 3
			if !twoOrThreeLiveNeighbours {
				toKill = append(toKill, cand)
			}
		} else {
			if len(aliveNeighbours) == 3 {
				toRevive = append(toRevive, cand)
			}
		}
	}
	setAlive(1, toRevive)
	setAlive(0, toKill)
}

func setAlive(alive int, cells []*Cell) {
	for _, c := range cells {
		c.Value = alive
	}
}

func dumpGrid(grid []*Cell) {
	simpleansi.ClearScreen()
	for _, c := range grid {
		simpleansi.MoveCursor(c.X, c.Y)
		if c.IsAlive() {
			fmt.Print(simpleansi.WithBackground("@", simpleansi.BLUE))
		} else {
			fmt.Print(simpleansi.WithBackground(" ", simpleansi.GREY))
		}
	}
}

func defineNeighbours(cell *Cell) []Cell {
	var neighbours []Cell
	for j := -1; j < 2; j++ {
		for i := -1; i < 2; i++ {
			neighbours = append(neighbours, NewNeighbourCell(i, j))
		}
	}
	return neighbours
}

func transformNeighbours(cell *Cell, neighbors []Cell) []Cell {
	var calculated []Cell
	for _, n := range neighbors {
		calculated = append(calculated, NewNeighbourCell(cell.X+n.X, cell.Y+n.Y))
	}
	return calculated
}

func getNeighbours(cell *Cell, grid []*Cell) []*Cell {
	var neighbours []*Cell
	definedNeighbours := defineNeighbours(cell)
	calculated := transformNeighbours(cell, definedNeighbours)

	for _, cand := range grid {
		if cell.ID == cand.ID {
			//Skip myself
			continue
		}
		for _, nd := range calculated {
			if nd.X == cand.X && nd.Y == cand.Y {
				neighbours = append(neighbours, cand)
			}
		}
	}
	return neighbours
}
