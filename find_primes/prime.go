// enter a number to get all primes between 0 < num.
// ex:
// $ go run prime.go [number]

package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  var num int
  var this_temp int
  var next_temp int

  if len(os.Args) > 2 {
    panic("Cannot except more then one argument")
  } else if len(os.Args) == 2 {
    arg := os.Args[1]
    if arg == "-h" || arg == "--help" {
      fmt.Println("This sctipt prints a comma-seperated string of the numbers from 0 to the input number \n" +
                  "\n" +
                  "-h|--help Print description of the sctipt \n" +
                  " \n" +
                  "Usage: \n" +
                  "       prime [arg]")
      os.Exit(0)
    }
    num, _ = strconv.Atoi(os.Args[1])
  } else {
    fmt.Print("Enter number: ")
    fmt.Scanln(&num)
  }

  // fmt.Println(num)
  for i := 2; i < num-1; i++ {
    this_temp = i*i
    next_temp = (i+1)*(i+1)

    if num > next_temp { 
      fmt.Print(this_temp, ",")
    } else {
      if this_temp != num {
        fmt.Print(this_temp)
      }
      break
    }
  }
  fmt.Println()
}