package main

import (
	"fmt"

	"github.com/UrsusArctos/wtttsolver/util/output"
	"github.com/UrsusArctos/wtttsolver/wttt/gametree"
	"github.com/UrsusArctos/wtttsolver/wttt/playfield"
)

const (
	sitCode = 11535 // arbitrary
)

func main() {
	fmt.Printf("Given situation: %d\n", sitCode)
	output.SitPrint(playfield.NewSituationFromCode(sitCode))
	if playfield.NewSituationFromCode(sitCode).IsWinning() {
		fmt.Println("Game over: WIN")
	} else {
		fmt.Println("Game in progress")
		Root := gametree.TTreeNode{SitCode: sitCode}
		Root.BuildFrom()
		fmt.Println("GameTree built")
		Root.Dump()
	}
}
