package gameoflife

import "fmt"

type Grid struct {
	gen    int
	dx, dy int
	grid   [][]bool
}

func NewGrid(dx, dy int) Grid {
	grid := make([][]bool, dx)
	for y := range grid {
		grid[y] = make([]bool, dx)
	}
	return Grid{0, dx, dy, grid}
}

func (g Grid) coords(x, y int) (int, int) {
	x += g.dx
	x %= g.dx
	y += g.dy
	y %= g.dy
	return x, y
}

func (g Grid) GetCell(x, y int) bool {
	idxX, idxY := g.coords(x, y)
	return g.grid[idxY][idxX]
}

func (g Grid) SetCell(x, y int, alive bool) {
	idxX, idxY := g.coords(x, y)
	g.grid[idxY][idxX] = alive
}

func (g Grid) Print() {
	fmt.Println("\033c")
	for y := len(g.grid) - 1; y >= 0; y-- {
		for x := 0; x < len(g.grid[y]); x++ {
			if g.grid[y][x] {
				fmt.Print("██")
			} else {
				fmt.Print("##")
			}
		}
		fmt.Println()
	}
	fmt.Println(g.gen)
}

type CellGenEvent struct {
	x, y  int
	alive bool
}

func (g *Grid) Gen() {
	resCh := make(chan CellGenEvent, g.dx*g.dy)
	for y := 0; y < len(g.grid); y++ {
		for x := 0; x < len(g.grid[y]); x++ {
			g.genCell(x, y, resCh)
		}
	}
	close(resCh)

	for ev := range resCh {
		g.grid[ev.y][ev.x] = ev.alive
	}
	g.gen++

}

func (g Grid) genCell(x, y int, res chan<- CellGenEvent) {
	alive := g.GetCell(x, y)
	cellGenEv := CellGenEvent{x, y, alive}
	aliveNeighbours := 0
	for ny := y - 1; ny <= y+1; ny++ {
		for nx := x - 1; nx <= x+1; nx++ {
			if nx == x && ny == y {
				continue
			}
			if g.GetCell(nx, ny) {
				aliveNeighbours++
			}
		}
	}
	if alive {
		if aliveNeighbours < 2 || aliveNeighbours > 3 {
			cellGenEv.alive = false
		}
	} else {
		if aliveNeighbours == 3 {
			cellGenEv.alive = true
		}
	}

	res <- cellGenEv
}
