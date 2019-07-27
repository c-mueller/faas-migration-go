package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Params struct {
	Endpoint string `json:"cloudant_host"`
	Username string `json:"cloudant_username"`
	Password string `json:"cloudant_password"`
	DBName   string `json:"cloudant_db_name"`
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Reading Failed")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Read Document\n")

	var p Params
	err = json.Unmarshal(data, &p)
	if err != nil {
		fmt.Println("Reading Failed")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	c
}
