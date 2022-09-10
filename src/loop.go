package main

import "fmt"

func main() {
  //por()
  //kontinue()
  //brik()
  nestedLoov()
}

func por() {
	for i:=0; i < 5; i++ {
		fmt.Println(i)
	  }
}

func kontinue() {
	for i:=0; i < 5; i++ {
		if i == 3 {
		  continue
		}
	   fmt.Println(i)
	  }
}

func brik() {
	for i:=0; i < 5; i++ {
		if i == 3 {
		  break
		}
	   fmt.Println(i)
	  }
}

func nestedLoov() {
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i:=0; i < len(adj); i++ {
	  for j:=0; j < len(fruits); j++ {
		fmt.Println(adj[i],fruits[j])
	  }
	}
}

//go run .\src\loop.go