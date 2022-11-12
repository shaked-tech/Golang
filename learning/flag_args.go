package main

import (
	"flag"
	"fmt"
	"os"
)

func print_help() {
    fmt.Println("An example sctipt to manage args using a go command-line.\n" +
                "\n" +
                "Usage:\n" +
                "\n" +
                "   go run flag_arg.go -m {message} [arguments]\n")
    flag.PrintDefaults()
    os.Exit(1)
}

func args_init(num *int, message *string) {
    flag.IntVar(num, "n", 1, "# of iterations")
    flag.StringVar(message, "m", "", "your message")
    flag.Parse()
    if *message == "" {
        fmt.Println("Message is empty")
        print_help()
    }
}

func main() {
    example_1()
    // example_2()
}

func example_1() {
    var num int
    var message string
    args_init(&num, &message)
    
    println("Example 1:")
    println("  Num:", num)
    println("  Message:", message)
}

func example_2() {
    num := flag.Int("n", 1, "# of iterations")
    message := flag.String("m", "", "your message")
    flag.Parse()
    
    n := *num
    m := *message
    
    println("Example 2:")
    println("   Num:", n)
    println("   Message:", m)
}