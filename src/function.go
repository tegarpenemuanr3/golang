package main

import "fmt"

//Function Argument
func familyName(fname string) {
  fmt.Println("Hello", fname, "Refsnes")
}

//function return
func myFunction(x int, y int) int {
	return x + y
}

//function name return
func myFunctions(x int, y int) (result int) {
	result = x + y
	return
  }

func main() {
  familyName("Liam")
  familyName("Jenny")
  familyName("Anja")

  fmt.Println(myFunction(1, 2))
  fmt.Println(myFunctions(2, 2))
}

// ! Running Terminal - go run .\src\function.go