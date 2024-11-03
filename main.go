package main

import (
	"fmt"

	"github.com/DC-Ops/puppy"
)

func main() {
	fmt.Println("Hello Go!")
	s1 := puppy.BasicBark()
	s2 := puppy.GshepBark()
	fmt.Println("Usually, a puppy says... " + s1)
	fmt.Println("While, a German Shepherd puppy says! " + s2)
}
