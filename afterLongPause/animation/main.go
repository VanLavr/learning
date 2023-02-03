package main

import (
	"fmt"
)


const width int = 120
const height int = 30
const screenLenght int = width * height
var screen [screenLenght]string

var pixel string = " "


func main() {
	
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			
			var x float64 = 1 * ( float64(i) / float64(width) )
			var y float64 = 1 * ( float64(j) / float64(height) )

			if ( ((x * x) + (y * y)) < 0.5 ) {
				pixel = "@"
			}
			screen[i + j * width] = pixel
		}
	}
	PrintScreen(screen)

}

func PrintScreen(scr [screenLenght]string) {
	for i := 0; i < screenLenght; i++ {
		fmt.Print(scr[i])
	}
}