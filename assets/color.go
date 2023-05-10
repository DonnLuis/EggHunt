package asset

import (
	"github.com/fatih/color"
)

/*
type Color struct {

	// Emoji Unicodes
	opened := make([]int, 0)                         // summary of opened ports
	green := color.New(color.FgGreen).SprintFunc()   // creates a green instance
	red := color.New(color.FgRed).SprintFunc()       // creates a red instance
	yellow := color.New(color.FgYellow).SprintFunc() // creates a yellow instance
}
*/

// colors available to use
var (
	Red    = color.New(color.FgRed).SprintFunc()    // creates a red instance
	Green  = color.New(color.FgGreen).SprintFunc()  // creates a green instance
	Yellow = color.New(color.FgYellow).SprintFunc() // creates a yellow instance
)
