package main

import (
	"fmt"
  
)
type greeter struct {
	name string
	greeting string
}

func (gr greeter) greet() {
	fmt.Println(gr.greeting, gr.name)
}

func (g *greeter) greetPointer() {
	fmt.Println(g.greeting, g.name)
}

func main() {
  g := greeter{
    greeting: "Hello",
		name: "Jerry",
	}
	g.greetPointer()
	g.greet()
  g.name = "Tom"
	g.greetPointer()
}
