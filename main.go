package main

import (
	"bufio"
	"fmt"
	"os"
)



func main() {
	var address, adresstemen string
	var umur int // deklarasi with var

	var umus int = 2 //deklararasi with var and value

	alamat := "Jawa timur " //deklarasi with  inference

	no := 90 // deklarasi wit inference

	var (
		nama string
		nim  int
		ipk  float32
	) //deklarasi with var majemuk

	_ = "jkshdfkshdskdhfksdhsdfsdf" //deklarasi underscore
	
	fmt.Println(umus)
	fmt.Println(alamat)
	fmt.Println(nama)
	fmt.Println(no)
	fmt.Println(ipk)
	fmt.Println(nim)
	fmt.Println(umur)
	fmt.Println("")

	fmt.Println()
	fmt.Print("Masukkan nama Anda :")
	fmt.Scan(nama)
	fmt.Println(nama)



	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	address = scanner.Text()
	fmt.Println("ini alamat pertama : ", address)


	scanner.Scan()
	adresstemen = scanner.Text()
	fmt.Println("ini adres temen",adresstemen)
	ukuran_baju := 23
	ukuran_celana := *&ukuran_baju

	fmt.Println(ukuran_celana)



}

//
