//  buatkan func untuk menghitung luas dan keliling bangunan datar, 
package main
import "fmt"

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