package main

import (
"fmt"
)

func publicHello() {
    fmt.Println("This is a public function, Hello!")
    // return true
}

func sayGreeting(greeting, name string) {
	name = "Ted"
	fmt.Println(greeting, name)
}

func sayGreetingPointer(greeting, name *string) {
	*name = "France"
	fmt.Println(*greeting, *name) // de-referance the poniter
}

func sum(values ...int) *int { // returning the pointer address  
	println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}
													// named return
func sumReturnPointer(values ...int) (result int) { // resault is declared as the return value, 
	println(values)                                 // this is less recomended, because it seems that the vale is returned not the pointer add.
	result = 0
	for _, v := range values {
		result += v
	}
	return result
}

func devide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot devide by Zero")
	} 
	return a / b, nil
}

func main() {
	// greet := "hello"
	// name := "shaked"
	
	// sayGreeting(greet, name)
	// println("Before poninter func", name)
	// sayGreetingPointer(&greet, &name) // Sending the adress of the var
	// println("After pointer func", name)
	
	// println("Reg:", greet, name)
	// println("Address:", &greet, &name)
	// // println("Point:", *greet, *name) // error
	
	// s := sum(1, 2, 3, 4, 5)
	// println("The sum is:", *s)

	// d, err := devide(5.0, 0.0) // error
	// d, err := devide(5.0, 3.0) // res
	// if err != nil {
	// 	// println("Error:", err) // Does not print the Error message...?
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// println("devided result is:", d)

	// // anonomus function
	// for i := 0; i <= 5; i++ {
	// 	f := func(i int) {
	// 		msg := "anonomus func. called:"
	// 		fmt.Println(msg, i)
	// 	}
	// 	f(i)
	// }

	// var devide func(float64, float64) (float64, error) // declare var of type function
	// devide = func (a,b float64) (float64, error) { //
	// 	if b == 0.0 {
	// 		return 0.0, fmt.Errorf("Cannot devide by Zero")
	// 	} 
	// 	return a / b, nil
	// }
	// d, err := devide(5.0, 3.0)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// println(d)

}