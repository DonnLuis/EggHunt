package main

type Color struct {

	// Emoji Unicodes
	egg := "\U0001F95A" // egg emoji unicode
	evil := "\U0001F608"
	opened := make([]int, 0)                         // summary of opened ports
	green := color.New(color.FgGreen).SprintFunc()   // creates a green instance
	red := color.New(color.FgRed).SprintFunc()       // creates a red instance
	yellow := color.New(color.FgYellow).SprintFunc() // creates a yellow instance
}