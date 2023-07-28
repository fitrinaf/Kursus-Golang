package main

import "testing"

func TestFungsiJumlah(t *testing.T) {
	hasil1 := Jumlah(3, 3)
	if hasil1 != 6 {
		t.Errorf("Hasil penjumlahan 3 + 3 = 6; hasil yang ditampil %v", hasil1)
	}
}
