package main

import (
	"fmt"
	"os"
)

var ipAddress = []string{
	"192.168.1.10",
	"172.100.2.11",
}

func main() {

	f, err := os.Create("htaccess.txt")
	if err != nil {
		fmt.Printf("create file htaccess.txt error : %v", err)
		os.Exit(1)
	}
	defer f.Close()

	// Write to file
	f.WriteString("Order Allow,Deny\n")
	for _, v := range ipAddress {
		f.WriteString(fmt.Sprintf("Allow from %s\n", v))
	}
	f.WriteString("Deny from all\n")
}
