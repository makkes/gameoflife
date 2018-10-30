package main

import (
	"fmt"
	"time"
)

type Cell struct {
	x     int
	y     int
	alive bool
	grid  *Grid
}

type Grid struct {
	gen    int
	dx, dy int
	grid   [][]Cell
}

func NewGrid(dx, dy int) Grid {
	res := Grid{}
	res.dx = dx
	res.dy = dy
	grid := make([][]Cell, dx)
	for y := 0; y < dy; y++ {
		grid[y] = make([]Cell, dx)
		for x := 0; x < dx; x++ {
			grid[y][x] = Cell{x, y, false, &res}
		}
	}
	res.grid = grid
	return res
}

func (g Grid) GetCell(x, y int) *Cell {
	var idxX, idxY int
	if y < 0 {
		idxY = len(g.grid) + y
	} else {
		idxY = y % len(g.grid)
	}
	if x < 0 {
		idxX = len(g.grid[idxY]) + x
	} else {
		idxX = x % len(g.grid[idxY])
	}
	return &g.grid[idxY][idxX]
}

func (g Grid) Print() {
	fmt.Println("\033c")
	for y := len(g.grid) - 1; y >= 0; y-- {
		for x := 0; x < len(g.grid[y]); x++ {
			if g.grid[y][x].alive {
				fmt.Print("██")
			} else {
				fmt.Print("##")
			}
		}
		fmt.Println()
	}
	fmt.Println(g.gen)
}

func (g *Grid) Gen() {
	resCh := make(chan Cell, g.dx*g.dy)
	for y := 0; y < len(g.grid); y++ {
		for x := 0; x < len(g.grid[y]); x++ {
			g.genCell(x, y, resCh)
		}
	}
	close(resCh)

	for newCell := range resCh {
		g.grid[newCell.y][newCell.x].alive = newCell.alive
	}
	g.gen++

}

func (g Grid) genCell(x, y int, res chan<- Cell) {
	cell := *g.GetCell(x, y)
	aliveNeighbours := 0
	for ny := y - 1; ny <= y+1; ny++ {
		for nx := x - 1; nx <= x+1; nx++ {
			if nx == x && ny == y {
				continue
			}
			n := g.GetCell(nx, ny)
			if n.alive {
				aliveNeighbours++
			}
		}
	}
	if cell.alive {
		if aliveNeighbours < 2 || aliveNeighbours > 3 {
			cell.alive = false
		}
	} else {
		if aliveNeighbours == 3 {
			cell.alive = true
		}
	}

	res <- cell
}

func main() {
	sizeX := 20
	sizeY := 20
	grid := NewGrid(sizeX, sizeY)

	// this stabilizes after 246 generations

	// Blinker
	grid.GetCell(11, 2).alive = true
	grid.GetCell(12, 2).alive = true
	grid.GetCell(13, 2).alive = true

	// Beacon
	grid.GetCell(3, 1).alive = true
	grid.GetCell(4, 1).alive = true
	grid.GetCell(4, 2).alive = true
	grid.GetCell(1, 3).alive = true
	grid.GetCell(1, 4).alive = true
	grid.GetCell(2, 4).alive = true

	// Glider
	grid.GetCell(2, 5).alive = true
	grid.GetCell(3, 5).alive = true
	grid.GetCell(4, 5).alive = true
	grid.GetCell(4, 6).alive = true
	grid.GetCell(3, 7).alive = true

	for {
		grid.Print()
		grid.Gen()
		time.Sleep(250 * time.Millisecond)
	}
}
