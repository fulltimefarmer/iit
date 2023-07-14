// Basic socket knowlege
// IP Types in go

package main

import (
	"fmt"
	"net"
)

func main() {
	// Converting an IPv4 or IPv6 string into an IP type
	// enforces the standards of what and IP should be
	// returns nil if formatting is incorrect

	//ipv4 := "127.0.0.1"
	ipString := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"

	ip := net.ParseIP(ipString) // type ip is an array of bytes
	if ip == nil {
		fmt.Println("The supplied address is not valid")
	}
	fmt.Println(ip)

	fmt.Printf("Loopback:    %t\n", ip.IsLoopback())
	fmt.Println("Unicast:")
	fmt.Printf("Global:    %t\n", ip.IsGlobalUnicast())
	fmt.Printf("Link:    %t\n", ip.IsLinkLocalUnicast())
	fmt.Println("Multicast:")
	fmt.Printf("Global:    %t\n", ip.IsMulticast())
	fmt.Printf("Interface    %t\n", ip.IsInterfaceLocalMulticast())
	fmt.Printf("Link    %t\n", ip.IsLinkLocalMulticast())

}
