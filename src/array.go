package main

import "fmt"

func main() {
	//deklarasiArray()
	//aksesElemenArray()
	//rubahElementArray()
	// inisialisasiArray()
	cekPanjangArray()
}

func deklarasiArray() {
	var arr1 = [3]int{1,2,3}
	arr2 := [5]int{4,5,6,7,8}
  
	fmt.Println(arr1)
	fmt.Println(arr2)
}

func aksesElemenArray() {
  prices := [3]int{10,20,30}

  fmt.Println(prices[0])
  fmt.Println(prices[2])
}

func rubahElementArray() {
  prices := [3]int{10,20,30}

  prices[2] = 50
  fmt.Println(prices)
}

func inisialisasiArray() {
  arr1 := [5]int{} //not initialized
  arr2 := [5]int{1,2} //partially initialized
  arr3 := [5]int{1,2,3,4,5} //fully initialized

  fmt.Println(arr1)
  fmt.Println(arr2)
  fmt.Println(arr3)
}

func cekPanjangArray() {
  arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  arr2 := [...]int{1,2,3,4,5,6}

  fmt.Println(len(arr1))
  fmt.Println(len(arr2))
}

// ! Running Terminal - go run .\src\array.go