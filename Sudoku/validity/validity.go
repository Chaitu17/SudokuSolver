package validity

// rowcolValidity: Checks the validity of sudoku grid row-wise and column-wise.

func rowcolValidity(grid [][]int, setting bool) bool {
	if setting {
		for row:=0; row<9; row++ {
			hash := make(map[int]int)
			for col:=0; col<9; col++ {
				if grid[row][col] != 0 {
					hash[grid[row][col]]++
					if hash[grid[row][col]] > 1 {
						return false
					}
				}
			}
		}
	} else {
		for col:=0; col<9; col++ {
			hash := make(map[int]int)
			for row:=0; row<9; row++ {
				if grid[row][col] != 0 {
					hash[grid[row][col]]++
					if hash[grid[row][col]] > 1 {
						return false
					}
				}
			}
		}
	}
	return true
}

// boxValidity: Checks the validity of the sudoku grid box-wise.

func boxValidity(grid [][]int) bool {
	for boxRow:=0; boxRow<7; boxRow += 3 {
		for boxCol:=0; boxCol<7; boxCol += 3 {
			hash := make(map[int]int)
			for i:=0; i<3; i++ {
				for j:=0; j<3; j++ {
					if grid[i+boxRow][j+boxCol] != 0 {
						hash[grid[i+boxRow][j+boxCol]]++
						if hash[grid[i+boxRow][j+boxCol]] > 1 {
							return false
						}
					}
				}
			}
		}
	}
	return true
}

// gridValidity: Checks the validity of the sudoku grid overall using rowcolValidity and boxValidity functions.

func GridValidity(grid [][]int) bool{
	return rowcolValidity(grid,true) && rowcolValidity(grid,false) && boxValidity(grid)
}