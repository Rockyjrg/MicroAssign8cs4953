package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 400, "raylib MicroAssign8 - window")

	defer rl.CloseWindow()

	weirdMouse := rl.LoadTexture("C:/_dev/Go/RaylibAssets/Mouse.png")
	newColor1 := rl.NewColor(150, 150, 150, 150)
	newColor2 := rl.NewColor(150, 245, 255, 255)
	newColor3 := rl.NewColor(255, 100, 0, 200)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawTextureEx(weirdMouse, rl.NewVector2(100, 0), 60, 4, newColor1)

		rl.DrawTextureEx(weirdMouse, rl.NewVector2(400, 120), 90, 10, newColor2)

		rl.DrawTextureEx(weirdMouse, rl.NewVector2(600, 150), 180, 2, newColor3)

		rl.EndDrawing()
	}
}
