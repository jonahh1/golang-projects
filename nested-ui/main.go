package main

// https://www.freecodecamp.org/news/go-beginners-handbook/#the-go-workspace

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rootFrame := BeginWindow(800, 450, "raylib [core] example - basic")
	font := rl.LoadFont("../assets/regular.ttf")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		if rl.IsWindowResized() {
			rootFrame.rect.w = int32(rl.GetScreenWidth())
			rootFrame.rect.h = int32(rl.GetScreenHeight())
		}
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		frameA := BeginFrame(rectint{10, 10, 50, 50}, rootFrame)
		frameB := BeginFrame(rectint{5, 5, 10, 10}, frameA)
		DrawRectAnchored(rectint{0, 0, 20, 20}, middleCenter, rl.Color{255, 0, 0, 128}, frameB)
		DrawRectAnchored(rectint{0, 0, 20, 20}, middleCenter, rl.Color{255, 0, 0, 128}, frameA)
		DrawRectAnchored(rectint{0, 0, 20, 20}, middleCenter, rl.Color{255, 0, 0, 128}, rootFrame)
		//EndFrame()
		//EndFrame()

		rl.DrawTextEx(font, "sjdviov", rl.Vector2{X: 10, Y: 10}, 32, 0, rl.Black)
		rl.EndDrawing()
	}

	EndWindow()
}
