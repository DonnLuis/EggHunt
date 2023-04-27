package ip

import (
	"fmt"
	"strconv"
)

type Subnet struct {
	sub []string
}

func IP(ip string, subnet int) {
	sub := strconv.Itoa(subnet)
	cidr := ip + "/" + sub
	fmt.Println(cidr)
}
