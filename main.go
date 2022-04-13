package main

import (
	"encoding/json"
	"log"
)

type Dog struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	HasTail   bool   `json:"has_tail"`
}

func main() {
	myJson := `[
{
		"first_name": "Millie",
		"last_name": "Collins", 
		"age": 3, 
		"has_tail": false
},
{
		"first_name": "Blueberry", 
		"last_name": "Muffinz", 
		"age": 5, 
		"has_tail": false
}
	]`
	// take pseudo received json and resolve into a struct

	var unmarshalled []Dog
	// note that json.Unmarshal second argument takes a reference

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("error in unmarshalling", err)
	}

	log.Println(unmarshalled)

	// now lets have a slice of structs that we need to marshal into json

	type Person struct {
		FirstName string
		LastName  string
		Email     string
	}

	var demoSlice []Person

	person1 := Person{
		FirstName: "Jack",
		LastName:  "Collins",
		Email:     "jc@test.com",
	}

	person2 := Person{
		FirstName: "LouLou",
		LastName:  "Collins",
		Email:     "pretty@test.com",
	}

	demoSlice = append(demoSlice, person1)
	demoSlice = append(demoSlice, person2)

	newJson, err := json.MarshalIndent(demoSlice, "", "   ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	log.Println(string(newJson))
}
