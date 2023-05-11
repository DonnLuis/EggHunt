package hunters

import (
	"fmt"
	"net"
	"time"

	asset "egghunt/assets"
)

func Hunt(target string, startPort, endPort int, ch chan int) {

	timeout := time.Duration(2 * time.Second) // set the duration
	opened := make([]int, 0)                  // summary of opened ports

	fmt.Print("Eggs found: ")
	for port := startPort; port < endPort; port++ {
		// displays if the checked port is closed or filtered.
		addr := fmt.Sprintf("%s:%d", target, port)
		/*
			   net.DialTimeout(<network type>, <address>, <timeout value>
				 => this functions starts a connection, with a timeout.

		*/

		/*
			addr variable:
					will hold the ip plus the port:
						example: 192.168.1.1:<port 22, 23, 24, and so on>

						the net.DialTimeout function will connect to the ip:port
						through each iteration aswell. it will basically look like this through iteration
						<ip>:22
						<ip>:23
						<ip>:24
						<ip>:25
						and so forth...
		*/
		conn, err := net.DialTimeout("tcp", addr /* <ip>:<port> */, timeout)
		if err != nil { // if err is not null/empty
			continue
		} // END IF

		// append opened ports to slice
		opened = append(opened, port)

		// Closes the connection
		defer conn.Close()

		// if a port is opened, this is displayedg
		fmt.Printf(
			"%s",
			asset.Egg,
		)

	} // END FOR

	// if len is 0, ports are closed or filtered
	if len(opened) == 0 {
		fmt.Println(asset.Red("All ports are Closed or Filtered."))
		// else, they are opened
	} else {
		// list and display the opened ports
		fmt.Printf("%s\n\tWhat a hunt!\nHere are the %s in your %s!\n%s\n",
			"\n============================",
			asset.Egg,
			asset.Basket,
			"============================",
		)
		// send the opened ports to the channel
		for _, port := range opened {
			ch <- port
		}
	}
	// before the function returns, close the channel.
	defer close(ch)
} // END HUNT()
