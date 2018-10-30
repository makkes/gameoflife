package main

import (
	"time"

	gol "github.com/makkes/gameoflife"
)

func main() {
	sizeX := 20
	sizeY := 20
	grid := gol.NewGrid(sizeX, sizeY)

	// this stabilizes after 246 generations

	// Blinker
	grid.SetCell(11, 2, true)
	grid.SetCell(12, 2, true)
	grid.SetCell(13, 2, true)

	// Beacon
	grid.SetCell(3, 1, true)
	grid.SetCell(4, 1, true)
	grid.SetCell(4, 2, true)
	grid.SetCell(1, 3, true)
	grid.SetCell(1, 4, true)
	grid.SetCell(2, 4, true)

	// Glider
	grid.SetCell(2, 5, true)
	grid.SetCell(3, 5, true)
	grid.SetCell(4, 5, true)
	grid.SetCell(4, 6, true)
	grid.SetCell(3, 7, true)

	for {
		grid.Print()
		grid.Gen()
		time.Sleep(30 * time.Millisecond)
	}
}
