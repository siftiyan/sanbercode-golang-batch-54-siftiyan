package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	// Soal 1
	var kataBootcamp = "Bootcamp"
	var kataDigital = "Digital"
	var kataSkill = "Skill"
	var kataSanbercode = "Sanbercode"
	var kataGolang = "Golang"

	kalimat := kataBootcamp + " " + kataDigital + " " + kataSkill + " " + kataSanbercode + " " + kataGolang
	fmt.Println("Soal 1:", kalimat)

	// Soal 2
	halo := "Halo Dunia"
	halo = strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println("Soal 2:", halo)

	// Soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	output := kataPertama + " " + strings.Title(kataKedua) + " " + strings.ToLower(kataKetiga) + " " + strings.ToUpper(kataKeempat)
	fmt.Println("Soal 3:", output)

	// Soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	angkaSatu, _ := strconv.Atoi(angkaPertama)
	angkaDua, _ := strconv.Atoi(angkaKedua)
	angkaTiga, _ := strconv.Atoi(angkaKetiga)
	angkaEmpat, _ := strconv.Atoi(angkaKeempat)

	total := angkaSatu + angkaDua + angkaTiga + angkaEmpat
	fmt.Println("Soal 4:", total)

	// Soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	hasil := strings.Replace(kalimat, "halo", "Hi", 2) + " - " + strconv.Itoa(angka)
	fmt.Println("Soal 5:", hasil)
}
