package ip

import (
	"fmt"
	"strconv"
)

type Subnet struct {
	ip     string
	subnet int
}

func IP(ip string, subnet int) {
	sub := strconv.Itoa(subnet)
	cidr := ip + "/" + sub
	fmt.Println(cidr)
}

func SubnetScan(cidr string) {

}

/*




	I'm still working on creating the appropriate functions to scan a subnet entered by the user.

	if a user enters a cidr of 192.168.0.0/24, I need to parse and scan the ip based on the subnet, where that is a 24,32, etc..

	should I use BufferReaders/Writers for this process? using the standard library: bufio.












*/
