package main

import "fmt"

type ipAddr [4]byte

func (ipAddr ipAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}

func main() {
	hosts := map[string]ipAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
