package main

import rl "github.com/gen2brain/raylib-go/raylib"
type vec2int struct {
	x,y int32
}
type frame struct {
	x,y,w,h int32
	
	parent *frame
}


func BeginWindow(width int32, height int32, title string) frame{
	windowFrame := frame{x: 0, y: 0, w: width, h: height, parent: nil}

	rl.InitWindow(width, height, title)
	
	return windowFrame
}

func EndWindow() {
	rl.CloseWindow()
}

func BeginFrame(x int32, y int32, width int32, height int32, parent frame) frame{
	newFrame := frame{
		x: x, y: y, w: width, h: height,
		parent: &parent,
	}
	real := getRelativePos(vec2int{x,y}, parent)
	rl.DrawRectangleLines(real.x, real.y, width, height, rl.Black)
	
	return newFrame
}

func EndFrame() {
	currentFrame = *currentFrame.parent
}

func DrawRect(x int32, y int32, width int32, height int32) {
	real := getRelativePos(vec2int{x,y}, currentFrame)
	rl.DrawRectangle(real.x, real.y, width, height, rl.Blue)
}

func getRelativePos(v vec2int, f frame) vec2int{
	realX := v.x
	realY := v.y
	// if f.parent == nil {
	// 	return vec2int{x:realX,y:realY}
	// }
	cf := f
	for cf.parent != nil {
		realX += cf.x
		realY += cf.y
		cf = *cf.parent
	}
	return vec2int{x:realX,y:realY}
}