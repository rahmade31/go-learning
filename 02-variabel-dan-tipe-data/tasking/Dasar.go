package main

import (
	"fmt"
)

func main() {

	// tasking 1
	var nama string = "Rahmad Windi Simanullang"
	var umur int = 25
	var tinggi float64 = 169.5
	var menikah bool = false

	fmt.Printf("Nama: %s, Umur: %d, Tinggi: %.2f, Menikah: %t\n", nama, umur, tinggi, menikah)

	// tasking 2
	var umur int
	var tinggi float32
	var nama string
	var aktif bool

	fmt.Println("umur:", umur)
	fmt.Println("tinggi:", tinggi)
	fmt.Println("nama:", nama)
	fmt.Println("aktif:", aktif)

	//tasking 3

	panjang := 10
	lebar := 5

	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	fmt.Println("Luas:", luas)
	fmt.Println("Keliling:", keliling)

	name := "Rahmad."

	if name == "Rahmad" {
		fmt.Println("Benar")
	} else if {
		fmt.Println("Salah")
	}
}
