package main

import "fmt"

// * Konstanta
const Nama string = "Tegar Penemuan"

// * Entry Point main()
func main() {
    fmt.Println("Hello World")

    pariabel()
}

func pariabel() {
    // * explisit tipe
    var hello string = "Hello "+Nama
    fmt.Println(hello)
    fmt.Printf("Cek type data: %T \n",hello)

    var on bool = true
    fmt.Println(on)

    var angka uint = 255
    fmt.Println(angka)

    var angkaFloat float64 = 255.898
    fmt.Println(angkaFloat)

    //default int = 0, string = ""

    // * implisit type (membuat variabel tanpa memberi type datanya/type data dihandle dari Go nya)
    var hai = "Hai Tegar Penemuan"
    fmt.Println(hai)

    // * shorcut Variabel
    hola := "Hai dalam bahasa Spanyol"
    fmt.Println(hola)
}

// ! Running Terminal - go run .\src\main.go