/*
* @Author: Ivy
* @Date: 2022/4/26 8:34 PM
**/
package main

import "fmt"

func main() {
	grid := [][]int{{1, 2}, {3, 4}}
	projection := projectionArea(grid)
	fmt.Println(projection)

	grid2 := [][]int{{2}}
	projection2 := projectionArea(grid2)
	fmt.Println(projection2)
}

func projectionArea(grid [][]int) int {
	xyProjectionArea := 0 //xOy投影面积
	xzProjectionArea := 0 //xOz投影面积
	yzProjectionArea := 0 //yOz投影面积

	// xOy投影面积
	for _, i := range grid {
		for _, j := range i {
			if j != 0 {
				xyProjectionArea = xyProjectionArea + 1

			}
		}
	}
	//fmt.Println(xyProjectionArea)

	// xOz投影面积
	for _, i := range grid {
		highest := 0
		for _, j := range i {
			if j > highest {
				highest = j
			}
		}
		xzProjectionArea = xzProjectionArea + highest

	}
	//fmt.Println(xzProjectionArea)

	// yOz投影面积
	for j := 0; j < len(grid); j++ {
		highest := 0
		for i := 0; i < len(grid[j]); i++ {
			if grid[i][j] > highest {
				highest = grid[i][j]
			}
		}
		yzProjectionArea = yzProjectionArea + highest
	}
	//fmt.Println(yzProjectionArea)

	return xyProjectionArea + xzProjectionArea + yzProjectionArea
}
