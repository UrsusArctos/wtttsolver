// Contains functions for printing of in-game entities in human-readable format
package output

import (
	"fmt"

	"github.com/UrsusArctos/wtttsolver/wttt/playfield"
)

var (
	symbol = [3]string{" ", "x", "o"}
)

// Prints deflattened playfield
func SitPrint(sit playfield.TGameSit) {
	fmt.Printf(" %s \u2502 %s \u2502 %s\n", symbol[sit.Cell[0]], symbol[sit.Cell[1]], symbol[sit.Cell[2]])
	fmt.Println("\u2500\u2500\u2500\u253c\u2500\u2500\u2500\u253c\u2500\u2500\u2500")
	fmt.Printf(" %s \u2502 %s \u2502 %s\n", symbol[sit.Cell[3]], symbol[sit.Cell[4]], symbol[sit.Cell[5]])
	fmt.Println("\u2500\u2500\u2500\u253c\u2500\u2500\u2500\u253c\u2500\u2500\u2500")
	fmt.Printf(" %s \u2502 %s \u2502 %s\n", symbol[sit.Cell[6]], symbol[sit.Cell[7]], symbol[sit.Cell[8]])
}
