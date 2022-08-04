package main

import (
	"fmt"
	"sudoku/validity"
	"sudoku/gridops"
	"io/ioutil"
	"strings"
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
	// Taking the input from a text file.
	data,_ := ioutil.ReadFile("test.txt")

	// Taking the test cases into an array
	// Splitting the data when we encounter a new line.
	tests := strings.Split(string(data), "\n")
	for i:=0; i<len(tests); i++ {
		// Initializing a 9 X 9 grid using slices 
		grid := make([][]int, 9)
		for i:=0; i<9; i++ {
			grid[i] = make([]int, 9)
		}

		// Parsing the string input into a 9 X 9 integer grid.
		count := 0
		nums := string(tests[i])
		for i:=0; i<len(nums); i++ {
			num := int(nums[i]) - 48
			grid[i/9][i - 9*(i/9)] = num
			if num != 0 {
			// Storing the count of non-zero values
				count++
			}
		}

		// Prints the inputted grid from the test.txt file.
		fmt.Println("This is the input grid:\n")
		gridops.PrintGrid(grid)

		// // Initializing the input grid using slices using the InputGrid function.
		// grid, count := gridops.InputGrid()

		// Minimum no of clues required to solve the puzzle is 17.
		errorHandler(grid, count)
		fmt.Println("")
	}	
}
