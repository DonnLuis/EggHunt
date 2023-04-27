package hunters

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"egghunt/asset"
)

func Hunt(target string, startPort, endPort int, ch chan int) {
	// help format how outputs look
	arrows_left := "===============>"
	arrows_right := "<=============="
	// set the duration
	timeout := time.Duration(2 * time.Second)
	// Emoji Unicodes
	opened := make([]int, 0) // summary of opened ports
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
			fmt.Printf(
				"Port %s is closed or filtered\n",
				asset.Red(strconv.Itoa(port)),
			)
			continue
		} // END IF

		opened = append(opened, port) // append opened ports to slice

		defer conn.Close() // Closes the connection

		// if a port is opened, this is displayedg
		fmt.Printf(
			"%s           %s Port %s is open %s         %s\n",
			asset.Green(arrows_left),
			asset.Egg,
			asset.Green(strconv.Itoa(port)),
			asset.Egg,
			asset.Green(arrows_right),
		)

	} // END FOR

	if len(opened) == 0 { // if len is 0, all ports are closed/filtered
		fmt.Println(asset.Red("All ports are Closed or Filtered."))
	} else {
		// list and display the opened ports
		fmt.Printf("%s\n\tWhat a hunt!\nHere are the %s in your %s!\n%s\n",
			"============================",
			asset.Egg,
			asset.Basket,
			"============================",
		)
		// send the opened ports to the channel
		for _, port := range opened {
			ch <- port
		}
	}

	defer close(ch)
} // END HUNT()
