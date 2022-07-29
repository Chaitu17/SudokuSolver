# SudokuSolver

This is a sudoku solver made in Go language.

As part of my learning, I decided to build a simple application. So, I chose to solve a sudoku solver.

This solver uses backtracking approach to solve the puzzle.

Validity module has the required validity checker functions which check if the inputted grid is valid or not.

Gridops module has many grid operation functions like InputGrid, PrintGrid which as the name suggest take input of a grind and print a grid respectively.

Gridops module also has many functions like rowCheck, colCheck, boxCheck, BlankBoxFinder etc. which help the recursive function in solving the grid.

Any optimisations and improvisations are welcomed.
