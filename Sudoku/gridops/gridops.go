package gridops

import "fmt"

// inputGrid: Asking and taking the grid input from the user

func InputGrid() ([][]int, int){
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
	return grid, count
}

// printGrid: This function will print the given sudoku grid.

func PrintGrid(grid [][]int) {
	for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println("")
	}
}

// blankBoxFinder: Finds a blank box and returns it's coordinates 
// If the grid doesn't have a blank box returns false

func BlankBoxFinder(grid [][]int) (check bool, numRow, numCol int){
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if grid[row][col] == 0 {
				check = true
				numRow = row
				numCol = col
				return
			}
		}
	}
	check = false
	numRow = 9
	numCol = 9
	return
}

// rowCheck: Checks if the given value exists in that row.

func rowCheck(grid [][]int, row int, num int) bool{
	for col:=0; col <9; col++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

// colCheck: Check if the given value exists in the column.

func colCheck(grid [][]int, col int, num int) bool{
	for row:=0; row <9; row++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

// boxCheck: Checks if the given value exists in the box.

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

// validityCheck: Checks if the given value is valid in that position or not.

func ValidityCheck(grid [][]int, row int, col int, num int) bool{
	if !rowCheck(grid, row, num) && !colCheck(grid, col, num) && !boxCheck(grid, 3*(row/3), 3*(col/3), num) && grid[row][col] == 0 {
		return true
	}
	return false
}