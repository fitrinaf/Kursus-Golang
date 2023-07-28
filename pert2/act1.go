package main

import (
	"fmt"
	"strconv"
)

// Membuat kalkulator perhitungan sederhana
// - Input: angka pertama, angka kedua, operator
// - Output: hasil perhitungan

// Contoh
// - Input: Masukkan angka pertama : 10
// - Input: Masukkan angka kedua : 5
// - Input: Masukkan operator : +
// pertambahan, pengurangan, perkalian, pembagian, pangkat, modulus
// - Output: Hasilnya adalah : 15

func main() {
	var (
		a, b, c  int
		operator string
	)
	fmt.Print("Masukkan angka pertama : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan angka kedua : ")
	fmt.Scan(&b)
	fmt.Println("Masukkan operator")
	var result string = "0"
	fmt.Println("+: Addition")
	fmt.Println("-: substraction")
	fmt.Println("*: Multiplication")
	fmt.Println("/: Division")
	fmt.Println("^: rank")
	fmt.Println("%: Mod")
	fmt.Print("Masukkan operator: ")
	fmt.Scan(&operator)

	if operator == "+" {
		c = a + b
		result = strconv.Itoa(c)
		fmt.Println("Hasilnya adalah :", result)
	} else if operator == "-" {
		c = a - b
		result = strconv.Itoa(c)
		fmt.Println("Hasilnya adalah :", result)
	} else if operator == "*" {
		c = a * b
		result = strconv.Itoa(c)
		fmt.Println("Hasilnya adalah :", result)
	} else if operator == "/" {
		c = a / b
		result = strconv.Itoa(c)
		fmt.Println("Hasilnya adalah :", result)
	} else if operator == "^" {
		c = a ^ b
		result = strconv.Itoa(c)
		fmt.Println("Hasilnya adalah :", result)
	} else if operator == "%" {
		c = a % b
		result = strconv.Itoa(c)
		fmt.Println("Hasilnya adalah :", result)
	} else {
		fmt.Println("Invalid operator")
	}

}

// FITRI NUR AFIFAH 10121491
