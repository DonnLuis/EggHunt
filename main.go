package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func main() {

	// User's Warning
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
	evil := "\U0001F47F"
	opened := make([]int, 0)                         // summary of opened ports
	green := color.New(color.FgGreen).SprintFunc()   // creates a green instance
	red := color.New(color.FgRed).SprintFunc()       // creates a red instance
	yellow := color.New(color.FgYellow).SprintFunc() // creates a yellow instance
	// tool starts
	fmt.Println(red(warning1))
	fmt.Println(yellow(warning2))

	fmt.Printf("Ready for the %s hunt? %s\n", egg, evil) // The greeting
	var target string                                    // target endpoint to scan
	fmt.Print("Pick your hunting ground [target IP]: ")
	fmt.Scanln(&target) // USER INPUT

	startPort, endPort := 22, 443             // starting/ending ports to scan
	timeout := time.Duration(2 * time.Second) // variable holding a 2 second duration

	// while (port) 80 < 443, port++
	for port := startPort; port < endPort; port++ {
		// displays if the checked port is closed or filtered.
		addr := fmt.Sprintf("%s:%d", target, port)

		/*
		   net.DialTimeout(<network type>, <address>, <timeout value>
		     => this function is used to attempt to connect to the port
		        at the specified address
		*/
		conn, err := net.DialTimeout("tcp", addr, timeout)
		if err != nil { // if err is not null/empty
			fmt.Printf("Port %s is closed or filtered\n", red(strconv.Itoa(port)))
			continue
		} // END IF

		opened = append(opened, port) // append opened ports to slice

		conn.Close() // Closes the connection
		// if a port is opened, this is displayed
		fmt.Printf("===============>           %s Port %s is open %s         <==============\n", egg, green(strconv.Itoa(port)), egg)
	} // END FOR

	fmt.Printf("\nEggs found!\n")
	for i := 0; i < len(opened); i++ {
		fmt.Printf("%s%s\n", egg, string(green(opened[i])))
	}
} // END MAIN
