package main

//Create a new type of 'deck'
//slice of strings
import (
	"fmt"
)

type deck []string

func (x deck) print() {
	for i, card := range x {
		fmt.Println(i, card)
	}

}
