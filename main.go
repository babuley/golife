package main

import (
	"fmt"
	"math/rand"
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

type ActiveCondition func(x int, y int, config *Config) bool

var generators = map[int]ActiveCondition{
	0: activeCondition1,
	1: activeCondition2,
	2: activeCondition3,
}

func activeCondition1(x int, y int, config *Config) bool {
	return x > 0 && x < config.Width-1 && y%3 == 0 && y > 3
}
func activeCondition2(x int, y int, config *Config) bool {
	return x > 3 && x < config.Width-3 && y%5 == 0 && y > 3
}
func activeCondition3(x int, y int, config *Config) bool {
	return x < 5 && x < config.Width-2 && y%2 != 0 && y > 3
}

func makeGrid(config *Config) []*Cell {
	var grid []*Cell
	rand.Seed(time.Now().UnixNano())
	gen := generators[rand.Intn(len(generators)-1)]
	for y := 0; y < config.Height; y++ {
		for x := 0; x < config.Width; x++ {
			alive := 0
			if gen(x, y, config) {
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
	//winner := make(chan *Cell)
	neighboursDef := NewNeighbourCell(0, 0).DefineNeighbours()
	for _, cand := range grid {
		ns := getNeighbours(cand, grid, neighboursDef)
		aliveNeighbours := getLiveNeighbours(ns)

		if cand.IsAlive() {
			var twoOrThreeLiveNeighbours = len(aliveNeighbours) == 2 || len(aliveNeighbours) == 3
			if !twoOrThreeLiveNeighbours {
				toKill = append(toKill, cand)
			}
		} else {
			if len(aliveNeighbours) == 3 {
				//winner <- cand
				toRevive = append(toRevive, cand)
			}
		}
	}
	setAlive(1, toRevive)
	//resurrect(winner)
	setAlive(0, toKill)
}

func resurrect(winners chan *Cell) {
	go func(winners chan *Cell) {
		for cell := range winners {
			cell.Value = 1
		}
	}(winners)
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

func transformNeighbours(cell *Cell, neighbors []*Cell) []*Cell {
	var calculated []*Cell
	for _, n := range neighbors {
		calculated = append(calculated, NewNeighbourCell(cell.X+n.X, cell.Y+n.Y))
	}
	return calculated
}

func getNeighbours(cell *Cell, grid []*Cell, definedNeighboursDef []*Cell) []*Cell {
	var neighbours []*Cell
	calculated := transformNeighbours(cell, definedNeighboursDef)

	for _, cand := range grid {
		for _, nd := range calculated {
			if nd.X == cand.X && nd.Y == cand.Y {
				neighbours = append(neighbours, cand)
			}
		}
	}
	return neighbours
}
