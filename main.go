package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const port = 8080

func main() {

	// fetching responses
	res, err := http.Get("https://random-data-api.com/api/v2/banks")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// convert from byte to json
	JsonRes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(JsonRes))

	requiredDetailsMap := map[string]string{}

	json.Unmarshal(JsonRes, &requiredDetailsMap)

	fmt.Println("iban: " + requiredDetailsMap["iban"])
	fmt.Println("bank_name: " + requiredDetailsMap["bank_name"])
	fmt.Println("routing_number: " + requiredDetailsMap["routing_number"])
}
