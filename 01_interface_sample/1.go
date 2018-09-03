package main

import (
	"fmt"
)

//Sample struct
type Sample struct {
}

//Test is a method
func (s *Sample) name() string {
	return ("cricket")
}

func main() {
	fmt.Println("in main pro")
	findGame(&Sample{})
}
