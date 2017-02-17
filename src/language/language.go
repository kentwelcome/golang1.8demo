package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name     string
	AgeYears int
	SSN      int
}

type Aux struct {
	Name     string `json:"full_name"`
	AgeYears int    `json:"age"`
	SSN      int    `json:"social_security"`
}

func readAux() Aux {
	raw, err := ioutil.ReadFile("./data.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var aux Aux
	json.Unmarshal(raw, &aux)
	return aux
}

func main() {
	aux := readAux()
	fmt.Printf("[Aux struct]\n")
	fmt.Printf("  %+v\n", aux)
	person := Person(aux)
	fmt.Printf("[Person struct]\n")
	fmt.Printf("  %+v\n", person)
}
