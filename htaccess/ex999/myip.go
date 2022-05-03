package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// {"ip":"66.249.75.9","country":"United States","cc":"US"}
type MyIp struct {
	Ip      string `json:'ip'`
	Country string `json:'country'`
	CC      string `json:'cc'`
}

func main() {
	res, err := http.Get("https://api.myip.com")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// res.StatusCode

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	data_obj := MyIp{}
	jsonErr := json.Unmarshal(data, &data_obj)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(data_obj)

}
