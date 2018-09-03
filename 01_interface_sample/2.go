package main

import "fmt"

//Game is
type Game interface {
	name() string
}

//Hello s
func findGame(i Game) {
	fmt.Println(i.name())
}
