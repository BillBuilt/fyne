//go:build ((!gles && !arm && !arm64) || darwin) && !js && !wasm && !web
// +build !gles,!arm,!arm64 darwin
// +build !js
// +build !wasm
// +build !web

package glfw

import "github.com/go-gl/glfw/v3.3/glfw"

func initWindowHints() {
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 0)

	glfw.WindowHint(glfw.CocoaGraphicsSwitching, glfw.True)
}
