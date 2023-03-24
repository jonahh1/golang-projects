package main

// Conway's game of life

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const width = 100
const height = 60
const size = 10

type cell struct {
	state         int
	canManipulate bool
}

func updateState(state [width][height]cell, scroll float32) [width][height]cell {
	newState := state

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			xLess1 := x - 1
			if xLess1 < 0 {
				xLess1 = width - 1
			}
			xPlus1 := x + 1
			if xPlus1 > width-1 {
				xPlus1 = 0
			}

			yLess1 := y - 1
			if yLess1 < 0 {
				yLess1 = height - 1
			}
			yPlus1 := y + 1
			if yPlus1 > height-1 {
				yPlus1 = 0
			}

			ntl := state[xLess1][yLess1].state // neighbor top left
			ntc := state[x][yLess1].state      // neighbor top centre
			ntr := state[xPlus1][yLess1].state // neighbor top right

			nml := state[xLess1][y].state // neighbor middle left
			nmr := state[xPlus1][y].state // neighbor middle left

			nbl := state[xLess1][yPlus1].state // neighbor bottom left
			nbc := state[x][yPlus1].state      // neighbor bottom centre
			nbr := state[xPlus1][yPlus1].state // neighbor bottom right

			totalNeighbors := ntl + ntc + ntr + nml + nmr + nbl + nbc + nbr
			if scroll > 0 {
				if totalNeighbors < 2 || totalNeighbors > 3 {
					newState[x][y].state = 0
				}
				if totalNeighbors == 3 {
					newState[x][y].state = 1
				}
			}
		}
	}
	return newState
}

func main() {
	var cells [width][height]cell
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			cells[x][y] = cell{state: 0, canManipulate: true}
		}
	}
	rl.InitWindow(width*size-1, height*size-1, "Conway's Game of Life")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		cells = updateState(cells, rl.GetMouseWheelMove())
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				col := rl.Black
				if cells[x][y].state == 1 {
					col = rl.White
				}
				rl.DrawRectangle((int32)(x*size), (int32)(y*size), size-1, size-1, col)
				mousePos := rl.GetMousePosition()
				if int(mousePos.X) > x*size && int(mousePos.Y) > y*size && int(mousePos.X) < x*size+size && int(mousePos.Y) < y*size+size {
					if rl.IsMouseButtonDown(0) && cells[x][y].canManipulate {
						cells[x][y].state = 1
						cells[x][y].canManipulate = false
					}
					if rl.IsMouseButtonDown(1) && cells[x][y].canManipulate {
						cells[x][y].state = 0
						cells[x][y].canManipulate = false
					}
				}
			}
		}
		if rl.IsMouseButtonReleased(0) {
			for x := 0; x < width; x++ {
				for y := 0; y < height; y++ {
					cells[x][y].canManipulate = true
				}
			}
		}
		if rl.IsKeyPressed(rl.KeyC) {
			for x := 0; x < width; x++ {
				for y := 0; y < height; y++ {
					cells[x][y].state = 0
				}
			}
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

// mouse down
// canManipulate = false
//
// mouse up
// canManipulate = true
