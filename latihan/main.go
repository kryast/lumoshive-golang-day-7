package main

import "fmt"

//  buatkan func untuk menghitung luas dan keliling bangunan datar, 
	func contoh(panjang int, lebar int) (int, int) {
	luasPersegi := panjang * lebar
	kelilingPersegi := (2 * panjang) + (2 * lebar)
	return luasPersegi, kelilingPersegi
}

func main () {
	
	Lp, Kp := contoh(5,10)

	fmt.Println(" Luas Persegi :", Lp)
	fmt.Println(" Keliling Persegi :", Kp)
}