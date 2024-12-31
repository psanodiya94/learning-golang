package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	jsonString := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]`

	var persons []Person

	err := json.Unmarshal([]byte(jsonString), &persons)
	if err != nil {
		log.Println("Error Unmarshalling JSON: ", err)
	}

	log.Printf("Persons: %+v", persons)

	var persons1 []Person
	persons1 = append(persons1, Person{
		FirstName: "Peter",
		LastName:  "Parker",
		HairColor: "black",
		HasDog:    false,
	})

	persons1 = append(persons1, Person{
		FirstName: "Tony",
		LastName:  "Stark",
		HairColor: "black",
		HasDog:    false,
	})

	jsonString1, err := json.MarshalIndent(persons1, "", "    ")
	if err != nil {
		log.Println("Error Marshalling JSON: ", err)
	}

	fmt.Printf(string(jsonString1))
}
