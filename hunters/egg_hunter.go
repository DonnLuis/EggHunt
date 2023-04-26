package hunters

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"egghunt/asset"
)

func Hunt(target string, startPort, endPort int) {

	//
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
			fmt.Printf("Port %s is closed or filtered\n", asset.Red(strconv.Itoa(port)))
			continue
		} // END IF

		opened = append(opened, port) // append opened ports to slice

		conn.Close() // Closes the connection
		// if a port is opened, this is displayedg
		fmt.Printf("===============>           %s Port %s is open %s         <==============\n", asset.Egg, asset.Green(strconv.Itoa(port)), asset.Egg)
	}
}
