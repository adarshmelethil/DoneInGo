
.PHONY: sudokueasy
sudokueasy:
	go run main.go sudoku brute --debug -p data/sudoku/easy/puzzle.sudoku

.PHONY: alt
alt:
	go run main.go sudoku alternate solve --debug -p data/sudoku/easy/puzzle.sudoku
