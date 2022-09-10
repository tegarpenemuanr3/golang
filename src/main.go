package main

import "fmt"

// * Konstanta
const Nama string = "Tegar Penemuan"

// * Entry Point main()
func main() {
    fmt.Println("Hello World")

    pariabel()
    sptf()
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

func sptf() {
    var number float32 = 10.0
    fmt.Printf("\nTipe data: %T dan Value: %v \n",number,number)

    //tampung dalam variabel
    var myString string = fmt.Sprintf("Tipe data: %T dan Value: %v",number,number)
    fmt.Println(myString)

    //menampilkan true dan false
    fmt.Printf("TRUE atau FALSE: %t \n", 1 < 2)

    //menampilkan binary (base 2) dari sebuah angka
    fmt.Printf("Binary: %b \n",10)

    //menampilkan oktal
    fmt.Printf("Oktal: %o \n", 10)

    //menampilkan desimal
    fmt.Printf("Desimal: %d \n", 10)

    //menampilkan hexadecimal
    fmt.Printf("Hexa: %x \n", 10)

    //menampilkan angka decimal dan dibulatkan
    fmt.Printf("Hasil: %f \n",2.1029322222222222)

    //menampilkan angka decimal secara menyeluruh
    fmt.Printf("Hasil: %g \n",2.1029322222222222)

    //menampilkan string default
    fmt.Printf("Bahasa: %s \n","Golang")

    //menampilkan string dalam tanda kutip
    fmt.Printf("Bahasa: %q \n", "Golang")

    //Panjang ke kiri
    fmt.Printf("Hello: %12s Tegar \n", "Penemuan")

    //Panjang ke kanan
    fmt.Printf("Hello: %-12s Tegar \n", "Penemuan")

    //panjang default, presisi 2
    fmt.Printf("Angka %.2f \n", 10.11231)
}



// ! Running Terminal - go run .\src\main.go