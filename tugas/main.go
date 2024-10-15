package main
import "fmt"

/*
	buatkan func :
	- diketahui produk sepatu adidas, puma, kappa
	-harga adidas 200.000
	- harga puma 150.000
	- kappa 600.000

	-harga diskon
	-jika adidas dan puma potongan 50.000
	- jika puma dan kappa potongan 150.000
	- jika adidas dan kappa potongan 75.000

	tampilkan total harga pembelian
*/

func diskon(adidas int, puma int, kappa int) int {
	diskon1 := 50000
	diskon2 := 150000
	diskon3 := 75000
	
	if adidas > 0 && puma > 0 {
		return diskon1
	} else if puma > 0 && kappa > 0 {
		return diskon2
	} else if adidas > 0 && kappa > 0 {
		return diskon3
	}
	return 0
}

func totalHargaPembelian(beliAdidas int, beliPuma int, beliKappa int) int {
	hargaAdidas := 200000
	hargaPuma := 150000
	hargaKappa := 600000
	
	total := (hargaAdidas * beliAdidas) + (hargaPuma * beliPuma) + (hargaKappa * beliKappa)
	
	diskonTotal := diskon(beliAdidas, beliPuma, beliKappa)
	total -= diskonTotal
	
	return total
}

func main() {
	beliAdidas := 1 
	beliPuma := 0   
	beliKappa := 1  
	
	total := totalHargaPembelian(beliAdidas, beliPuma, beliKappa)
	fmt.Println("Total harga pembelian setelah diskon:", total)
}