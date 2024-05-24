// Game tree generator and analyzer
package gametree

import (
	"fmt"

	"github.com/UrsusArctos/wtttsolver/util/auxmath"
	"github.com/UrsusArctos/wtttsolver/util/output"
	"github.com/UrsusArctos/wtttsolver/wttt/playfield"
)

type (
	TPossibleMove struct {
		node *TTreeNode
		rank int32
	}

	TTreeNode struct {
		// Defines the situation that may have further outcomes
		SitCode playfield.TSitCode
		// Defines all possible outcomes
		nextMove [3][auxmath.FlatLen]TPossibleMove // MarkNone index shoud not be used
	}
)

// Build complete subtree under the current situation node
func (Node *TTreeNode) BuildFrom() {
	// Unravel game situation
	gamesit := playfield.NewSituationFromCode(Node.SitCode)
	// 1. check if the situation is already winning
	if !gamesit.IsWinning() {
		// 2. Enumerate all cells for all marks
		for cMark := playfield.TMark(playfield.MarkCross); cMark <= playfield.MarkNought; cMark++ {
			for ci := range gamesit.Cell {
				// 3. Check if the cell is empty
				if gamesit.Cell[ci] == playfield.MarkNone {
					// 4.1 Make a move into that cell
					newMove := playfield.NewSituation(&gamesit, uint8(ci), cMark)
					// 4.2 Construct new tree node
					newNode := TTreeNode{SitCode: newMove.GetSitCode()}
					// 4.3 Link this new tree node in current tree node
					Node.nextMove[cMark][ci].node = &newNode
					Node.nextMove[cMark][ci].rank = 0
					// 4.4. Explore subtree
					newNode.BuildFrom()
				}
			}
		}
	}
}

// [WIP] Dump the subtree
func (Node TTreeNode) Dump() {
	for cMark := playfield.TMark(playfield.MarkCross); cMark <= playfield.MarkNought; cMark++ {
		for pm := range Node.nextMove[cMark] {
			if Node.nextMove[cMark][pm].node != nil {
				// possible move found, get it
				PossibleMove := *Node.nextMove[cMark][pm].node
				fmt.Printf("Possible move %d\n", PossibleMove.SitCode)
				output.SitPrint(playfield.NewSituationFromCode(PossibleMove.SitCode))
				if playfield.NewSituationFromCode(PossibleMove.SitCode).IsWinning() {
					fmt.Println("GAME ENDS: WIN!")
				}
				PossibleMove.Dump()
			}
		}
	}
	///////
	fmt.Println("=== End of branch")
}
