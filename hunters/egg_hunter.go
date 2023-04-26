package hunters

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/fatih/color"

)

func Hunt(target string, startPort, endPort int) {
	// Emoji Unicodes
	opened := make([]int, 0)                         // summary of opened ports
	green := color.New(color.FgGreen).SprintFunc()   // creates a green instance
	red := color.New(color.FgRed).SprintFunc()       // creates a red instance

	timeout := time.Duration(2 * time.Second)
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
			fmt.Printf("===============>           %s Port %s is open %s         <==============\n", Emote.Egg, green(strconv.Itoa(port)), Emote.Egg)
		}
}