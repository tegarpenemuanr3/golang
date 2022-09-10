package main

import (
	"fmt"
	"strconv"
)

func main() {
	var stringToAngka string = "10"
	angkaInt, _ := strconv.ParseInt(stringToAngka,10,64)
	fmt.Printf("Data tipe: %T value: %v \n",angkaInt, angkaInt)

	//ParseFloat
	//ParseBool -> 0 dan 1
}

// ! Running Terminal - go run .\src\konversi_tipe_data.go