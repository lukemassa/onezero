package game

import (
	"fmt"
)

const DEFAULT_BOARD_SIZE = 12

type Board [][]bool

type Location struct {
	X int
	Y int
}

func NewBlankBoard() Board {
	size := DEFAULT_BOARD_SIZE
	b := make([][]bool, size)
	for i := range b {
		b[i] = make([]bool, size)
	}
	return b
}

func NewBoardFromString(s string) Board {
	boolArray, err := convertToBoolArray(s)
	if err != nil {
		panic(err)
	}
	return Board(boolArray)
}

func (b Board) Show() {
	fmt.Println("Board:")
	for j := range b {
		for i := range b[j] {
			if b[j][i] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (b Board) ShowWithPotentialPiece(p Piece, l Location) {
	fmt.Println("Board:")
	size := len(b)
	p_x, p_y := p.Dimensions()
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			potentialPiece := i >= l.X && j >= l.Y && i < l.X+p_x && j < l.Y+p_y && p[j-l.Y][i-l.X]
			alreadyFilled := b[j][i]
			if potentialPiece {
				if alreadyFilled {
					fmt.Print("?")
				} else {
					fmt.Print("o")
				}
			} else if alreadyFilled {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (b Board) Overlaps(p Piece, l Location) bool {
	size := len(b)
	p_x, p_y := p.Dimensions()
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			potentialPiece := i >= l.X && j >= l.Y && i < l.X+p_x && j < l.Y+p_y && p[j-l.Y][i-l.X]
			alreadyFilled := b[j][i]
			if potentialPiece && alreadyFilled {
				return true
			}
		}
	}
	return false
}

func (b Board) GetPossibleLocations(piece Piece) []Location {
	ret := make([]Location, 0)
	piece_x, piece_y := piece.Dimensions()
	size := len(b)
	for j := range b {
		for i := range b[j] {
			legalMove := true
			for y := 0; y < piece_y; y++ {
				for x := 0; x < piece_x; x++ {
					test_x := i + x
					test_y := j + y
					// If this is out of range, not valid
					if test_x >= size || test_y >= size {
						legalMove = false
						break
					}
					// If overlap with piece, not valid
					if b[test_y][test_x] && piece[y][x] {
						legalMove = false
						break
					}
				}
				if !legalMove {
					break
				}
			}
			if !legalMove {
				continue
			}
			// This is a legal move, add it
			ret = append(ret, Location{i, j})
		}
	}
	return ret
}

func (b *Board) Place(piece Piece, l Location) bool {
	cleared := false
	piece_x, piece_y := piece.Dimensions()
	board := *b
	size := len(board)
	for j := 0; j < piece_y; j++ {
		for i := 0; i < piece_x; i++ {
			// If piece has a block at j, i
			// fill in the board at j, i, offset by location
			if piece[j][i] {
				//fmt.Printf("ADDING at (%d %d)\n", j+l.y, i+l.x)
				board[j+l.Y][i+l.X] = true
			}
		}
	}
	// Clean rows and columns

	fullRows := make([]int, 0)
	fullColumns := make([]int, 0)

	// Check for full rows
	for j := 0; j < size; j++ {
		isRowFull := true
		for i := 0; i < size; i++ {
			if !board[j][i] {
				isRowFull = false
				break
			}
		}
		if isRowFull {
			fullRows = append(fullRows, j)
		}
	}

	// Check for full columns
	for i := 0; i < size; i++ {
		isColumnFull := true
		for j := 0; j < size; j++ {
			if !board[j][i] {
				isColumnFull = false
				break
			}
		}
		if isColumnFull {
			fullColumns = append(fullColumns, i)
		}
	}

	// We can't clear as we we find them snice you might
	// clear a column and row simultaneously
	for _, row := range fullRows {
		cleared = true
		for i := 0; i < size; i++ {
			board[row][i] = false
		}
	}
	for _, column := range fullColumns {
		cleared = true
		for j := 0; j < size; j++ {
			board[j][column] = false
		}
	}
	return cleared
}
