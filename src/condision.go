package main

import "fmt"

func main() {
	nilai := 200

	if nilai >= 90 && nilai <=100 {
		fmt.Println("Sangat Memuaskan")
	} else if nilai >= 80 && nilai < 90 {
		fmt.Println("Memuaskan")
	} else if nilai >= 70 && nilai < 80 {
		fmt.Println("Baik")
	} else if nilai >= 60 && nilai < 70 {
		fmt.Println("Cukup")
	} else if nilai < 60 {
		fmt.Println("Tidak Lulus")
	} else {
		fmt.Println("Hello")
	}
}

//go run .\src\condision.go