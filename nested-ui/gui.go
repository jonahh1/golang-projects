package main

import rl "github.com/gen2brain/raylib-go/raylib"

type rectint struct {
	x, y, w, h int32
}
type vec2int struct {
	x, y int32
}

const (
	topLeft   = 0
	topCenter = 1
	topRight  = 2

	middleLeft   = 3
	middleCenter = 4
	middleRight  = 5

	bottomLeft   = 6
	bottomCenter = 7
	bottomRight  = 8
)

func BeginWindow(width int32, height int32, title string) rectint {
	windowFrame := rectint{0, 0, width, height}

	rl.InitWindow(width, height, title)

	return windowFrame
}

func EndWindow() {
	rl.CloseWindow()
}

func BeginFrame(rect rectint, anchor int, parent rectint) rectint {
	/*newFrame := frame{
		rect:   rect,
		parent: &parent,
	}*/
	newRect := rectint{w: rect.w, h: rect.h}

	//DrawRectLines(rect, rl.Black, parent)

	return newRect
}

/*
func EndFrame(frame) {
	currentFrame = *currentFrame.parent
}*/

func DrawRect(rect rectint, col rl.Color, parent rectint) {
	DrawRectAnchored(rect, topLeft, col, parent)
}
func DrawRectLines(rect rectint, col rl.Color, parent rectint) {
	real := getRelativePos(rect, topLeft, parent)
	rl.DrawRectangleLines(real.x, real.y, rect.w, rect.h, col)
}
func DrawRectAnchored(rect rectint, anchor int, col rl.Color, parent rectint) {
	real := getRelativePos(rect, anchor, parent)
	rl.DrawRectangle(real.x, real.y, rect.w, rect.h, col)
} /*
func getRelativePos(v vec2int, f frame) vec2int {
	realX := v.x
	realY := v.y
	// if f.parent == nil {
	// 	return vec2int{x:realX,y:realY}
	// }
	currentFrame := f
	for currentFrame.parent != nil {
		realX += currentFrame.rect.x
		realY += currentFrame.rect.y
		currentFrame = *currentFrame.parent
	}
	return vec2int{x: realX, y: realY}
}*/
func getRelativePos(rect rectint, anchor int, parent rectint) rectint {
	newRect := rectint{w: rect.w, h: rect.h}

	// find the position of the rectangle as if it were at 0,0
	// modify the rect position based on the anchor

	switch a := anchor; a {
	case topLeft:
		newRect.x = rect.x
		newRect.y = rect.y
	case topCenter:
		newRect.x = (parent.w-rect.w)/2 + rect.x
		newRect.y = rect.y
	case topRight:
		newRect.x = (parent.w - rect.w) - rect.x
		newRect.y = rect.y

	case middleLeft:
		newRect.x = rect.x
		newRect.y = (parent.h-rect.h)/2 + rect.y
	case middleCenter:
		newRect.x = (parent.w-rect.w)/2 + rect.x
		newRect.y = (parent.h-rect.h)/2 + rect.y
	case middleRight:
		newRect.x = (parent.w - rect.w) - rect.x
		newRect.y = (parent.h-rect.h)/2 + rect.y

	case bottomLeft:
		newRect.x = rect.x
		newRect.y = (parent.h - rect.h) - rect.y
	case bottomCenter:
		newRect.x = (parent.w-rect.w)/2 + rect.x
		newRect.y = (parent.h - rect.h) - rect.y
	case bottomRight:
		newRect.x = (parent.w - rect.w) - rect.x
		newRect.y = (parent.h - rect.h) - rect.y
	}
	newRect.x -= parent.x
	newRect.y -= parent.y
	return newRect
}
