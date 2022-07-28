package main

import (
	"fmt"
)

// Finds a blank box and returns it's coordinates 
// If the grid doesn't have a blank box returns false

func blankBoxFinder(grid [][]int) (check bool, num_row, num_col int){
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if grid[row][col] == 0 {
				check = true
				num_row = row
				num_col = col
				return
			}
		}
	}
	check = false
	num_row = 9
	num_col = 9
	return
}

//  Check if the given value exists in the row.

func rowCheck(grid [][]int, row int, num int) bool{
	for col:=0; col <9; col++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

// Check if the given value exists in the column.

func colCheck(grid [][]int, col int, num int) bool{
	for row:=0; row <9; row++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

// Checks if the given value exists in the box.

func boxCheck(grid [][]int, boxRow int, boxCol int, num int) bool{
	for row:=0; row<3; row++ {
		for col:=0; col<3; col++ {
			if grid[row+boxRow][col+boxCol] == num {
				return true
			}
		}
	}
	return false
}

// Checks if the given value is valid in that position or not.

func validityCheck(grid [][]int, row int, col int, num int) bool{
	if !rowCheck(grid, row, num) && !colCheck(grid, col, num) && !boxCheck(grid, 3*(row/3), 3*(col/3), num) && grid[row][col] == 0 {
		return true
	}
	return false
}

// Recursive function which gives us the solution of the puzzle.

func sudokuSolver(grid [][]int) bool {
	check, row, col := blankBoxFinder(grid)

	if !check {
		return true
	}
	
	for i:=1; i<=9; i++ {
		if validityCheck(grid, row, col, i) {
			grid[row][col] = i					// Putting the value in it and checking if it gives us the required solution
			if sudokuSolver(grid) {
				return true
			}
			grid[row][col] = 0					// If it doesn't fit, then we change it back to the original value
		}
	}
	return false
}


func main() {
	// Initializing the input grid using slices.
	grid := make([][]int, 9)
	for i:=0; i<9; i++ {
		grid[i] = make([]int, 9)
	}
	count := 0
	fmt.Println("Please give the input grid below:")
	for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			fmt.Scanf("%d", &grid[i][j])
			if grid[i][j] != 0 {
				count++
			}
		}
	}

	// Minimum no of clues required to solve the puzzle is 17.
	if count < 17 {
		fmt.Println("\nSorry, the input that you have given doesn't have a unique solution.")
		fmt.Println("Please give atleast 17 non-zero valid clues to solve the puzzle.\n")
	} else {
		if sudokuSolver(grid) {
			fmt.Println("This will be the solution of the given grid.")
			for i:=0; i<9; i++ {
				for j:=0; j<9; j++ {
					fmt.Printf("%d ", grid[i][j])
				}
				fmt.Println("")
			}
		} else {
			fmt.Println("The given grid is not valid.")
		}
	}
}