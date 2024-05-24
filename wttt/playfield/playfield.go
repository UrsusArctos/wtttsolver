// Defines WTTT playfield
package playfield

import "github.com/UrsusArctos/wtttsolver/util/auxmath"

const (
	// Define possible marks on a playfield: none, cross, nought
	MarkNone = iota
	MarkCross
	MarkNought
)

type (
	// Define elementary types
	TMark    = uint8
	TSitCode = uint16

	// Define flattened playfield
	TPlayField = [auxmath.FlatLen]TMark

	// Define game situation class
	TGameSit struct {
		Cell TPlayField
	}
)

// Constuct new TGameSit instance from another TGameSit template
func NewSituation(origin *TGameSit, cellIndex uint8, newMark TMark) (newsit TGameSit) {
	if origin != nil {
		// copy existing situation
		copy(newsit.Cell[:], origin.Cell[:])
		// place new mark
		newsit.Cell[cellIndex] = newMark
	}
	return newsit
}

// Constuct new TGameSit instance from situation code
func NewSituationFromCode(sitCode TSitCode) (newsit TGameSit) {
	newsit.SetSituation(sitCode)
	return newsit
}

// Derive situation code from current position on the playfield
func (GS TGameSit) GetSitCode() TSitCode {
	var sitcode TSitCode = 0
	for pos, trit := range GS.Cell {
		sitcode += auxmath.Pow3[pos] * uint16(trit)
	}
	return sitcode
}

// Set current position on the playfield based on the situation code
func (GS *TGameSit) SetSituation(sitcode TSitCode) {
	for pos := range GS.Cell {
		GS.Cell[pos] = TMark((sitcode / auxmath.Pow3[pos]) % 3)
	}
}

// Check sameness or marks on three cell indices
func (GS TGameSit) areSame(ci0 uint8, ci1 uint8, ci2 uint8) bool {
	return (GS.Cell[ci0] != MarkNone) && (GS.Cell[ci0] == GS.Cell[ci1]) && (GS.Cell[ci0] == GS.Cell[ci2])
}

// Check whether the current situation is winning
func (GS TGameSit) IsWinning() bool {
	return (GS.areSame(0, 1, 2) || GS.areSame(3, 4, 5) || GS.areSame(6, 7, 8)) || // any horizontal row
		(GS.areSame(0, 3, 6) || GS.areSame(1, 4, 7) || GS.areSame(2, 5, 8)) || // any vertical row
		(GS.areSame(0, 4, 8) || GS.areSame(2, 4, 6)) // any diagonal
}
