package players

import (
	"fmt"

	"github.com/lukemassa/onezero/pkg/game"
	"github.com/pkg/term"
)

type TUIPlayer struct{}

func (t TUIPlayer) Decide(b game.Board, p game.Piece, locations []game.Location) int {
	panic("Can't decide amongst choices")
}

// Returns either an ascii code, or (if input is an arrow) a Javascript key code.
func getChar() (ascii int, keyCode int, err error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)

	var numRead int
	numRead, err = t.Read(bytes)
	if err != nil {
		return
	}
	if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
		// Three-character control sequence, beginning with "ESC-[".

		// Since there are no ASCII codes for arrow keys, we use
		// Javascript key codes.
		if bytes[2] == 65 {
			// Up
			keyCode = 38
		} else if bytes[2] == 66 {
			// Down
			keyCode = 40
		} else if bytes[2] == 67 {
			// Right
			keyCode = 39
		} else if bytes[2] == 68 {
			// Left
			keyCode = 37
		}
	} else if numRead == 1 {
		ascii = int(bytes[0])
	} else {
		// Two characters read??
	}
	t.Restore()
	t.Close()
	return
}

func (i TUIPlayer) Move(b game.Board, p game.Piece) *game.Location {
	x_dim, y_dim := p.Dimensions()
	size := len(b)
	possibleLocations := b.GetPossibleLocations(p)
	if len(possibleLocations) == 0 {
		return nil
	}
	location := possibleLocations[0]
	for {
		fmt.Print("\033[H\033[2J")

		b.ShowWithPotentialPiece(p, location)
		p.Show()
		ascii, keyCode, err := getChar()
		if err != nil {
			panic(err)
		}
		if ascii == 3 {
			panic("")
		}
		if ascii == 110 {
			found := false
			for i := 0; i < len(possibleLocations)-1; i++ {
				if possibleLocations[i].X == location.X && possibleLocations[i].Y == location.Y {
					location = possibleLocations[i+1]
					found = true
					break
				}
			}
			if !found {
				location = possibleLocations[0]
			}
			continue
		}
		if ascii == 98 {
			found := false
			for i := 1; i < len(possibleLocations); i++ {
				if possibleLocations[i].X == location.X && possibleLocations[i].Y == location.Y {
					location = possibleLocations[i-1]
					found = true
					break
				}
			}
			if !found {
				location = possibleLocations[len(possibleLocations)-1]
			}
			continue
		}
		if keyCode == 0 {
			if ascii == 13 {
				// If it's valid, take it
				for i := 0; i < len(possibleLocations); i++ {
					if location.X == possibleLocations[i].X && location.Y == possibleLocations[i].Y {
						return &location
					}
				}
			}
			continue
		}
		fmt.Println(keyCode)
		x_inc, y_inc := 0, 0
		if keyCode == 37 {
			x_inc, y_inc = -1, 0
		} else if keyCode == 38 {
			x_inc, y_inc = 0, -1
		} else if keyCode == 39 {
			x_inc, y_inc = 1, 0
		} else if keyCode == 40 {
			x_inc, y_inc = 0, 1
		}
		newLocation := game.Location{X: location.X, Y: location.Y}
		for {
			if newLocation.X+x_inc < 0 || newLocation.Y+y_inc < 0 || newLocation.X+x_dim+x_inc > size || newLocation.Y+y_dim+y_inc > size {
				newLocation = location
				break
			}
			newLocation.X = newLocation.X + x_inc
			newLocation.Y = newLocation.Y + y_inc
			if !b.Overlaps(p, newLocation) {
				fmt.Println("no overlap")
				break
			}
			//fmt.Printf("%d x %d\n (incremented %d x %d)", newLocation.X, newLocation.Y, x_inc, y_inc)

		}

		location = newLocation
	}
}
