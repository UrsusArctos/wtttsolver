// Contains auxiliary math routines such as base3 conversion
package auxmath

const (
	// Define length of the playfield flattened into an array
	FlatLen = 9
)

var (
	// Precomputed powers of three (0..8)
	Pow3 = [FlatLen]uint16{1, 3, 9, 27, 81, 243, 729, 2187, 6561}
)
