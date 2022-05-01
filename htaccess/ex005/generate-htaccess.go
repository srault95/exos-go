package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
)

type Ip struct {
	Address string
	Role    string
}

// Valid ip format
func (ip *Ip) validIp() error {
	p := net.ParseIP(ip.Address)
	if p == nil {
		return errors.New(fmt.Sprintf("invalid IP Address: %s", ip.Address))
	}
	return nil
}

// Load CSV and return ip list + rule
func loadCSV(filepath string) (addr []Ip, err error) {
	csvFile, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("open file %s error : %v", filepath, err)
		return addr, err
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Printf("parse csv file error : %v", err)
		return addr, err
	}

	for i, line := range csvLines {
		// bypass headers line
		if i == 0 {
			continue
		}

		ip := Ip{
			Address: line[0],
			Role:    line[1],
		}

		addr = append(addr, ip)
	}

	return addr, nil
}

// Filter clean IP
func clean(addr []Ip) (cleanAddr []Ip) {
	for _, ip := range addr {
		err := ip.validIp()
		if err != nil {
			log.Println(err)
			continue
		}
		cleanAddr = append(cleanAddr, ip)
	}
	return cleanAddr
}

func htaccess(filepath string, addr []Ip) error {
	f, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("create file %s error : %v", filepath, err)
		return err
	}
	defer f.Close()

	f.WriteString("Order Allow,Deny\n")

	for _, ip := range addr {
		f.WriteString(fmt.Sprintf("%s from %s\n", ip.Role, ip.Address))
	}

	f.WriteString("Deny from all\n")

	return nil
}

func main() {
	csvFilePath := "ip-addresses.csv"
	htaccessFilePath := "htaccess.txt"

	ipList, err := loadCSV(csvFilePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	addr := clean(ipList)

	err = htaccess(htaccessFilePath, addr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
