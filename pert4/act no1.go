package main

import "fmt"

// Fitri Nur Afifah 10121491

func naikanGaji(gajiAwal *int, gajiAkhir *int) {
	*gajiAwal = *gajiAkhir
}

func main() {
	var expectedSalary, salaryNow int
	fmt.Print("Masukkan gaji anda : ")
	fmt.Scan(&salaryNow)

	fmt.Print("Masukkan gaji yang anda inginkan : ")
	fmt.Scan(&expectedSalary)

	naikanGaji(&salaryNow, &expectedSalary)

	fmt.Printf("\nGaji anda sekarang %d\n", salaryNow)
}
