/*
Author: Luis Soto (github.com/DonnLuis)

*/

package main

import (
	"fmt"

	"egghunt/asset"
	"egghunt/hunters"
)

func main() {

	// tool title
	title :=
		`
	▓█████   ▄████   ▄████  ██░ ██  █    ██  ███▄    █ ▄▄▄█████▓
	▓█   ▀  ██▒ ▀█▒ ██▒ ▀█▒▓██░ ██▒ ██  ▓██▒ ██ ▀█   █ ▓  ██▒ ▓▒
	▒███   ▒██░▄▄▄░▒██░▄▄▄░▒██▀▀██░▓██  ▒██░▓██  ▀█ ██▒▒ ▓██░ ▒░
	▒▓█  ▄ ░▓█  ██▓░▓█  ██▓░▓█ ░██ ▓▓█  ░██░▓██▒  ▐▌██▒░ ▓██▓ ░ 
	░▒████▒░▒▓███▀▒░▒▓███▀▒░▓█▒░██▓▒▒█████▓ ▒██░   ▓██░  ▒██▒ ░ 
	░░ ▒░ ░ ░▒   ▒  ░▒   ▒  ▒ ░░▒░▒░▒▓▒ ▒ ▒ ░ ▒░   ▒ ▒   ▒ ░░   
	░ ░  ░  ░   ░   ░   ░  ▒ ░▒░ ░░░▒░ ░ ░ ░ ░░   ░ ▒░    ░    
	░   ░ ░   ░ ░ ░   ░  ░  ░░ ░ ░░░ ░ ░    ░   ░ ░   ░      
	░  ░      ░       ░  ░  ░  ░   ░              ░          
    `
	// User's Warning
	// ASCII WARNING taken from
	// https://patorjk.com/software/taag/#p=display&f=Big%20Money-sw&t=WARNING
	warning1 := `
 __       __   ______   _______   __    __  ______  __    __   ______  
/  |  _  /  | /      \ /       \ /  \  /  |/      |/  \  /  | /      \ 
$$ | / \ $$ |/$$$$$$  |$$$$$$$  |$$  \ $$ |$$$$$$/ $$  \ $$ |/$$$$$$  |
$$ |/$  \$$ |$$ |__$$ |$$ |__$$ |$$$  \$$ |  $$ |  $$$  \$$ |$$ | _$$/ 
$$ /$$$  $$ |$$    $$ |$$    $$ |$$$$  $$ |  $$ |  $$$$  $$ |$$ |/   |
$$ $$/$$ $$ |$$$$$$$$ |$$$$$$$  |$$ $$ $$ |  $$ |  $$ $$ $$ |$$ |$$$$ |
$$$$/  $$$$ |$$ |  $$ |$$ |  $$ |$$ |$$$$ | _$$ |_ $$ |$$$$ |$$ \__$$ |
$$$/    $$$ |$$ |  $$ |$$ |  $$ |$$ | $$$ |/ $$   |$$ | $$$ |$$    $$/ 
$$/      $$/ $$/   $$/ $$/   $$/ $$/   $$/ $$$$$$/ $$/   $$/  $$$$$$/  

`
	warning2 :=
		`
		(Unless you have permission)
		Use this tool only on your own systems, and no one else's.
		I created this tool as a fun way to harden my own network.
		This tool is not meant to be used to break laws.
		You have been warned! 
		`

	// tool starts
	fmt.Println(asset.Red(warning1))
	fmt.Println(asset.Yellow(warning2))
	fmt.Println(title)                                               // tool title
	fmt.Printf("Ready for the %s hunt? %s\n", asset.Egg, asset.Evil) // The greeting
	var target string                                                // target endpoint to scan
	fmt.Print("Pick your hunting ground [target IP]: ")
	fmt.Scanln(&target) // USER INPUT

	// The channel goroutine used to communicate between the go routine
	// and the Hunt()
	ch := make(chan int, 2048)

	startPort, endPort := 22, 443               // starting/ending ports to scan
	go hunters.Hunt(target, startPort, endPort, ch) // invoke and Hunt for eggs

	// Receive the results from the channel
	for openedPort := range ch {
		fmt.Printf("%s%s\n", asset.Egg, string(asset.Green(openedPort)))
	}
} // END MAIN
