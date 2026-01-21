package main

import (
	"errors"
	"fmt"
)

func main() {

	bands := []string{"Metallica", "Megadeth", "Slayer", "Anthrax"}

	fmt.Println(len(bands))
}

func nameSurname(firstname string, lastname string) (string, int, error) {

	fullname := firstname + " " + lastname
	var err error

	if firstname == "lol" || lastname == "lol" {
		err = errors.New("invalid name")
	}
	age := len(fullname)
	return fullname, age, err
}
