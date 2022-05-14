package main

import (
	"fmt"
	"strconv"
)

func Public() {
    fmt.Print("This is a public function")
    return true
}

func main() {
    // Variables:
    var name string = "noName"    
    name = "shaked"
    fmt.Println("Hello, World! My name is", name)
    fmt.Printf("Hello, World! My name is %s (type '%T')\n", name, name)
    
    // a := 22
    // b := 22.2
    // fmt.Printf("a: %d (type %T) \nb: %f (type %T)\n", a, a, b, b)
    // var c int = int(b)
    // fmt.Printf("c: %d (type %T), data loss!\n", c ,c)

    i := 100
    f := float32(i)
    s := strconv.Itoa(i)
    fmt.Printf("%T: %d, %T: %f, %T: %s\n", i, i, f, f, s, s)

    // Primitives
    var t bool = true
    fmt.Printf("%v, %T\n", t, t)

    m := 1 == 2
    fmt.Printf("%v, %T\n", m, m)

    var e bool
    fmt.Printf("%v, %T\n", e, e)

    var n uint16 = 42 
    fmt.Printf("%v, %T\n", n, n)

    a := 10
    b := 3
    fmt.Println(a + b)
    fmt.Println(a - b)
    fmt.Println(a * b)
    fmt.Println(a / b)
    fmt.Println(a % b)

    var Int int = 8 // 
    var Int8 int8 = 8 // 
    var Int16 int16 = 16 // 
    var Int32 int32 = 32 // 
    var Int64 int64 = 64 // 

    fmt.Printf("%d, %T\n", Int8, Int8)
    fmt.Printf("%d, %T\n", Int16, Int16)
    fmt.Printf("%d, %T\n", Int32, Int32)
    fmt.Printf("%d, %T\n", Int64, Int64)

    fmt.Println(Int + int(Int16)) // cant apply actions between different types, must be converted

}