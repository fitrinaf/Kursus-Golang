package main

// Fitri Nur Afifah 10121491
import "fmt"

var nilai1, nilai2 float64

func main() {
	defer fmt.Println("---- SELESAI ----")
	fmt.Print("Masukkan Bilangan 1: ")
	fmt.Scan(&nilai1)
	fmt.Print("Masukkan Bilangan 2: ")
	fmt.Scan(&nilai2)
	hasil := nilai1 / nilai2
	fmt.Printf("Hasil dari Nilai1 / Nilai2 = %.3f\n", hasil)
}
