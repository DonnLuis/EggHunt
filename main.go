/*
Copyright © 2023 NAME HERE devstuffin@gmail.com
Author: Luis Soto (github.com/DonnLuis)

*/

package main

import (
	"egghunt/asset"
	"egghunt/hunters"
	"egghunt/recon"
	"fmt"
)

func main() {

	// warning
	assets.Warn() // display warning of using this tool

	fmt.Printf("===============[ Ready for the %s hunt? %s ]===================\n", asset.Egg, asset.Evil) // The greeting
	var target string
	fmt.Printf("OS: %s\n", recon.GetOS())
	// target endpoint to scan
	fmt.Print("Pick the hunting ground [target IP]: ")
	fmt.Scanln(&target) // USER INPUT

	// ask user if they want to check a specific subnet
	//ip.IP("192.168.1.0", 24) // testing data, erase when done testing

	// The channel goroutine used to communicate between the go routine
	// and the Hunt()
	ch := make(chan int, 2048)

	startPort, endPort := 22, 443                   // starting/ending ports to scan
	go hunters.Hunt(target, startPort, endPort, ch) // invoke and Hunt for eggs

	// Receive the results from the channel
	for openedPort := range ch {
		fmt.Printf("%s%s\n", asset.Egg, string(asset.Green(openedPort)))
	}
} // END MAIN
