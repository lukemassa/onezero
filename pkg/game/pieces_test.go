package game

import (
	"testing"
)

func FuzzCompilePiece(f *testing.F) {
	f.Add("XX\nXX")
	f.Add("XXX\nXXX")
	f.Add("XXX\nXXX\nXXX")
	f.Fuzz(func(t *testing.T, s string) {
		piece, err := compilePiece(s)
		if err != nil {
			return
		}
		if piece.String() != s {
			t.Errorf("%s (len %d) turned into piece %s (len %d)", s, len(s), piece.String(), len(piece.String()))
		}
	})

}

func FuzzCompilePieceOnlyGoodChars(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		s := byteArrayToStringRepr(b)
		piece, err := compilePiece(s)
		if err != nil {
			return
		}
		if piece.String() != s {
			t.Errorf("%s (len %d) turned into piece %s (len %d)", s, len(s), piece.String(), len(piece.String()))
		}
	})

}
