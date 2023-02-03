
// screen[i * width + j] = pixel // filling full screen

package main

import (
	"fmt"
)


const width int = 120
const height int = 30
const screenLenght int = width * height
var screen [screenLenght]string

var pixel string = " "
var borderRelation float64 = float64(width) / float64(height) // relation width to height ( 4 : 1 )
var symbolFrameRelation float64 = 11.0 / 24.0
// var symbolFrameRelation float64 = 8.0 / 16.0


func main() {
	
	for k := 0; k < 10000; k++ {

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				pixel = " "
				
				var x float64 = 2 * ( ( float64(j) / float64(width) ) - 0.5 ) // normalizing width (from -1 to 1)
				var y float64 = 2 * ( ( float64(i) / float64(height) ) - 0.5 ) // normalizing height (from -1 to 1)
				x *= borderRelation * symbolFrameRelation // x and y coordinates in consider with screen and symbol border relations


				if ( (x * x + y * y) < 0.5 ) {
					pixel = "@"
				}
				screen[i * width + j] = pixel
				//fmt.Println(x, y)
			}
		}

		PrintScreen(screen)

	}

}

func PrintScreen(scr [screenLenght]string) {
	for i := 0; i < screenLenght; i++ {
		fmt.Print(scr[i])
	}
}