package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
)

type IPAddressEntry struct {
	IPAddress string
	Role      string
}

func (ip *IPAddressEntry) validateIpAddress() error {
	p := net.ParseIP(ip.IPAddress)
	if p == nil {
		return errors.New(fmt.Sprintf("invalid IP Address: %s", ip.IPAddress))
	}
	return nil
}

func main() {

	csvFile, err := os.Open("ip-addresses.csv")
	if err != nil {
		fmt.Printf("open file ip-addresses.csv error : %v", err)
		os.Exit(1)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Printf("read csv file error : %v", err)
		os.Exit(1)
	}

	f, err := os.Create("htaccess.txt")
	if err != nil {
		fmt.Printf("create file htaccess.txt error : %v", err)
		os.Exit(1)
	}
	defer f.Close()

	f.WriteString("Order Allow,Deny\n")

	for i, line := range csvLines {
		// bypass headers line
		if i == 0 {
			continue
		}

		// Optional Struct creation
		ip := IPAddressEntry{
			IPAddress: line[0],
			Role:      line[1],
		}
		err := ip.validateIpAddress()
		if err != nil {
			log.Println(err)
			continue
		}
		f.WriteString(fmt.Sprintf("%s from %s\n", ip.Role, ip.IPAddress))
	}
	f.WriteString("Deny from all\n")
}
