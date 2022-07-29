package main

import (
	"fmt"
	"sudoku/validity"
	"sudoku/gridops"
)

// sudokuSolver: Recursive function which gives us the solution of the puzzle.

func sudokuSolver(grid [][]int) bool {
	check, row, col := gridops.BlankBoxFinder(grid)

	if !check {
		return true
	}
	
	for i:=1; i<=9; i++ {
		if gridops.ValidityCheck(grid, row, col, i) {
			grid[row][col] = i					// Putting the value in it and checking if it gives us the required solution
			if sudokuSolver(grid) {
				return true
			}
			grid[row][col] = 0					// If it doesn't fit, then we change it back to the original value
		}
	}
	return false
}

// errorHandler: Handles the errors and if there is no error it runs the solver and gives the solution.

func errorHandler(grid [][]int, count int) {
	if count < 17 {
		fmt.Println("\nSorry, the input that you have given doesn't have a unique solution,")
		fmt.Println("Please give atleast 17 non-zero valid clues to solve the puzzle.\n")
	} else {
		if !validity.GridValidity(grid) {
			fmt.Println("\nThe given grid is not valid. Please give a valid grid.\n")
		} else {
			if sudokuSolver(grid) {
				fmt.Println("\nThis is the solution of the given grid:\n")
				gridops.PrintGrid(grid)
			}
		}
	}
}

func main() {
	// Initializing the input grid using slices.
	grid, count := gridops.InputGrid()

	// Minimum no of clues required to solve the puzzle is 17.
	errorHandler(grid, count)
}