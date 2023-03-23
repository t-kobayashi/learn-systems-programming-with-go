package main

import (
	"fmt"
)

type Talker interface {
	Talk()
}

type Greeeter struct {
	name string
}

func (g *Greeeter) Talk() {
	fmt.Printf("Hello, my name is %s", g.name)
}

func main() {
	var talker Talker
	talker = &Greeeter{"Tetsuya"}
	talker.Talk()
}
