package main

// https://www.freecodecamp.org/news/go-beginners-handbook/#the-go-workspace

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	BeginWindow(800, 450, "raylib [core] example - basic")
	font := rl.LoadFont("assets/regular.ttf")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		BeginFrame(10, 10, 50, 50)
		BeginFrame(5, 5, 10, 10)
		//DrawRect(0,0,20,20)
		EndFrame()
		EndFrame()
		
		rl.DrawTextEx(font,"sjdviov", rl.Vector2{X: 10, Y: 10}, 32, 0, rl.Black);
		rl.EndDrawing()
	}

	EndWindow()
}