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

	// Emoji Unicodes
	egg := "\U0001F95A" // egg emoji unicode
	evil := "\U0001F608"
	opened := make([]int, 0) // summary of opened ports

	// tool starts
	fmt.Println(asset.Red(warning1))
	fmt.Println(asset.Yellow(warning2))
	fmt.Println(title)                                   // tool title
	fmt.Printf("Ready for the %s hunt? %s\n", egg, evil) // The greeting
	var target string                                    // target endpoint to scan
	fmt.Print("Pick your hunting ground [target IP]: ")
	fmt.Scanln(&target) // USER INPUT

	startPort, endPort := 22, 443            // starting/ending ports to scan
	hunters.Hunt(target, startPort, endPort) // invoke and Hunt for eggs

	if len(opened) == 0 { // if len is 0, the port is closed/filtered
		fmt.Println(asset.Red("All ports are Closed or Filtered."))
	} else { // if len is not 0, port n is open
		fmt.Printf("\nEggs found!\n")
	}

	// list and display the opened ports
	for i := 0; i < len(opened); i++ {
		fmt.Printf("%s%s\n", egg, string(asset.Green(opened[i])))
	}
} // END MAIN
