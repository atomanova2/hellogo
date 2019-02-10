package game

import "fmt"

func Game() {

	var grid [3][3]int

	grid[0][0] = 1
	grid[0][1] = 1
	grid[1][0] = 1

	fmt.Printf("grid( %v, %v) = %v)", 0, 0, grid[0][0])

}
