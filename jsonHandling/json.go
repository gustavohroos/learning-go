package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Name        string
	Surname     string
	Age         int
	Profession  string
	BankAccount float64
}

func main() {
	jamesbond := person{
		"James", "Bond", 42, "Secret Agent", 253329.23,
	}
	darthvader := person{
		"Anakin", "Skywalker", 50, "Intergalatic Vilain", 10000000000.54,
	}

	jb, err := json.Marshal(jamesbond)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(string(jb))
	}

	dv, err := json.Marshal(darthvader)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(string(dv))
	}

	var persons []*person

	dvPerson := new(person)
	err = json.Unmarshal(dv, dvPerson)
	if err != nil {
		fmt.Println("Error unmarshaling Darth Vader:", err)
	} else {
		persons = append(persons, dvPerson)
	}

	jbPerson := new(person)
	err = json.Unmarshal(jb, jbPerson)
	if err != nil {
		fmt.Println("Error unmarshaling James Bond:", err)
	} else {
		persons = append(persons, jbPerson)
	}

	file := "persons.json"
	f, err := os.Create(file)
	defer f.Close()
	if err != nil {
		fmt.Println("Error creating file:", err)
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "    ")

	err = enc.Encode(persons)
	if err != nil {
		fmt.Println("Error encoding person:", err)
	}

	persons = []*person{}
	file = "personsToRead.json"
	f, err = os.Open(file)
	defer f.Close()
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(&persons)
	if err != nil {
		fmt.Println("Error decoding person:", err)
	}

	for _, p := range persons {
		fmt.Printf("Name: %v\nSurname: %v\nAge: %v\nProfession: %v\nBank Account: %v\n\n", p.Name, p.Surname, p.Age, p.Profession, p.BankAccount)
	}

}
