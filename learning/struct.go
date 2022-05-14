package main

import (
	"fmt"
)

func main() {
	g := greeter{ 
		greeting: "Hello",
		name: "Jerry",
	}
	g.greetPointer()
	fmt.Println(g.name)
}

type greeter struct {
	greeting string
	name string
}

func (gr greeter) greet() {
	fmt.Println(gr.greeting, gr.name)
}

func (g *greeter) greetPointer() {
	fmt.Println(g.greeting, g.name)
 	g.name = "Tom"
}
