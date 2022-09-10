package main

import "fmt"

func main() {
	angka := 1

	switch {
	case angka == 1:
		{
			fmt.Println("Angka = 1")
			break
		}
	case angka == 2:
		{
			fmt.Println("Angka = 2")
			break
		}
	default:
		{
			fmt.Println("Angka tidak diketahui")
		}
	}
}

// go run .\src\switchcase.go