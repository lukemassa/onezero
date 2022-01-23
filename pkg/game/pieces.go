package game

import (
	"errors"
	"fmt"
	"strings"
)

type Piece [][]bool

func (p Piece) Show() {
	fmt.Println("Piece:")
	fmt.Printf("%s\n", p)
}

func (p Piece) Dimensions() (int, int) {
	return len(p[0]), len(p)
}

func (p Piece) String() string {
	var sb strings.Builder
	wroteNewLine := false
	for i := range p {
		for j := range p[i] {
			if p[i][j] {
				sb.WriteString("X")
			} else {
				sb.WriteString(" ")
			}
		}
		wroteNewLine = true
		sb.WriteString("\n")
	}
	ret := sb.String()

	if wroteNewLine {
		ret = ret[:len(ret)-1]
	}
	return ret
}

func byteArrayToStringRepr(b []byte) string {
	var sb strings.Builder
	for i := range b {
		if i%3 == 0 {
			sb.WriteString(" ")
		} else if i%3 == 1 {
			sb.WriteString("X")

		} else if i%3 == 2 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func convertToBoolArray(repr string) ([][]bool, error) {

	x := 0
	y := 0
	for _, line := range strings.Split(repr, "\n") {

		if y != 0 && x != len(line) {
			return nil, errors.New(fmt.Sprintf("First line was %d long, %d-th line is %d long", x, y, len(line)))
		}
		y += 1
		x = len(line)
		for _, char := range line {
			if char != 'X' && char != ' ' && char != '.' {
				return nil, errors.New(fmt.Sprintf("Invalid char in repr: %c", char))
			}
		}
	}
	//fmt.Printf("x=%d, y=%d\n", x, y)
	b := make([][]bool, y)
	for i := range b {
		b[i] = make([]bool, x)
	}
	atLeastOneSquare := false
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			loc := j*(x+1) + i
			//fmt.Printf("%d: %c\n", loc, repr[loc])
			if repr[loc] == 'X' {
				b[j][i] = true
				atLeastOneSquare = true
			}
		}
	}
	if !atLeastOneSquare {
		return nil, errors.New("Must have at least one square in block")
	}
	return b, nil
}

func mustCompilePiece(repr string) Piece {

	piece, err := convertToBoolArray(repr)
	if err != nil {
		panic(err)
	}
	return Piece(piece)
}

func getPiecees() []Piece {
	return []Piece{
		// 1
		mustCompilePiece("X"),
		// 2
		mustCompilePiece("XX"),
		mustCompilePiece("X\nX"),

		// 3
		mustCompilePiece("XX\nX "),
		mustCompilePiece("XX\n X"),
		mustCompilePiece("X \nXX"),
		mustCompilePiece(" X\nXX"),
		mustCompilePiece("XXX"),
		mustCompilePiece("X\nX\nX"),

		// 4
		mustCompilePiece("XXXX"),
		mustCompilePiece("X\nX\nX\nX"),

		// 5
		mustCompilePiece("XXXXX"),
		mustCompilePiece("X\nX\nX\nX\nX"),

		// 9
		mustCompilePiece("XXX\nXXX\nXXX"),
	}
}
