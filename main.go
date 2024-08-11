package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 40
	height = 20
)

type World [height][width]bool

func main() {
	world := randomWorld()
	for {
		printWorld(world)
		world = nextGeneration(world)
		time.Sleep(time.Second / 2)
	}
}

func randomWorld() World {
	var world World
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			world[i][j] = rand.Intn(2) == 1
		}
	}
	return world
}

func printWorld(w World) {
	fmt.Print("\033[H\033[2J") // Limpia la pantalla
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if w[i][j] {
				fmt.Print("■")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func nextGeneration(w World) World {
	var next World
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			count := countNeighbors(w, i, j)
			if w[i][j] {
				next[i][j] = count == 2 || count == 3
			} else {
				next[i][j] = count == 3
			}
		}
	}
	return next
}

func countNeighbors(w World, x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			nx, ny := x+i, y+j
			if nx >= 0 && nx < height && ny >= 0 && ny < width && w[nx][ny] {
				count++
			}
		}
	}
	return count
}
