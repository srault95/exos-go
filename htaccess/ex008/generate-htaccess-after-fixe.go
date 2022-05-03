package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"text/template"
)

type Ip struct {
	Address string
	Role    string
}

type Options struct {
	CsvPath      string
	HtaccessPath string
	TmplPath     string
	Debug        bool
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

func htaccess(tmplPath string, filepath string, addr []Ip) error {

	f, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("create file %s error : %v", filepath, err)
		return err
	}
	defer f.Close()

	template, err := template.ParseFiles(tmplPath)
	if err != nil {
		fmt.Printf("parse template error : %v", err)
		return err
	}
	err = template.Execute(f, addr)
	if err != nil {
		fmt.Printf("parse template error : %v", err)
		return err
	}

	return nil
}

func parseOptions() (options Options) {

	csvPath := "ip-addresses.csv"
	htaccessPath := "htaccess.txt"
	tmplPath := "htaccess.tmpl"

	val, ok := os.LookupEnv("CSV_PATH")
	if ok {
		csvPath = val
	}

	val, ok = os.LookupEnv("HTACCESS_PATH")
	if ok {
		htaccessPath = val
	}

	val, ok = os.LookupEnv("TEMPLATE_PATH")
	if ok {
		tmplPath = val
	}

	options = Options{
		CsvPath:      *flag.String("c", csvPath, "CSV filepath"),
		HtaccessPath: *flag.String("h", htaccessPath, "Htaccess filepath"),
		TmplPath:     *flag.String("t", tmplPath, "Template filepath"),
		//Debug:        *flag.Bool("D", false, "Debug mode"),
	}
	flag.Parse()

	return options
}

func main() {

	options := parseOptions()

	ipList, err := loadCSV(options.CsvPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	addr := clean(ipList)

	err = htaccess(options.TmplPath, options.HtaccessPath, addr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
