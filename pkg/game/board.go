package game

import (
	"fmt"
)

const DEFAULT_BOARD_SIZE = 4

type Board [][]bool

type Location struct {
	x int
	y int
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

func (b Board) getPossibleLocations(piece Piece) []Location {
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
					// If this is out of range, or the spot is covered
					if test_x >= size || test_y >= size || b[test_y][test_x] {
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

func (b *Board) place(piece Piece, l Location) {
	piece_x, piece_y := piece.Dimensions()
	for j := 0; j < piece_y; j++ {
		for i := 0; i < piece_x; i++ {
			// If piece has a block at j, i
			// fill in the board at j, i, offset by location
			if piece[j][i] {
				//fmt.Printf("ADDING at (%d %d)\n", j+l.y, i+l.x)
				(*b)[j+l.y][i+l.x] = true
			}
		}
	}
}
