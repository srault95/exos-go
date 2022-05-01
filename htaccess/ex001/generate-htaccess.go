package main

import (
	"fmt"
)

var ipAddress = []string{
	"192.168.1.10",
	"172.100.2.11",
}

func main() {
	// Print screen
	fmt.Println("Order Allow,Deny")
	for _, v := range ipAddress {
		fmt.Printf("Allow from %s\n", v)
	}
	fmt.Println("Deny from all")
}
