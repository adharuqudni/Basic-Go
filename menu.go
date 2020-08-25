package main

import (
	"fmt"
	"os"
)

func main() {
	var pilihan int
	var kembali string
	var jumlahData int
	var nama [10]string
	var nim [10]string
	var ipk [10]float32
	k := 0

	iya := true
	for iya {
		iya = true
		pilihan = 0
		fmt.Println("::::::::::::: MENU :::::::::::::")
		fmt.Println("1. INPUT DATA")
		fmt.Println("2. TAMPILKAN DATA")
		fmt.Println("3. KELUAR")
		fmt.Print("PILIH : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Println("INPUT DATA")
			fmt.Print("banyak data : ")
			fmt.Scan(&jumlahData)
			for i := 0; i < jumlahData; i++ {
				fmt.Print("nama : ")
				fmt.Scan(&nama[i])
				fmt.Print("nim : ")
				fmt.Scan(&nim[i])
				fmt.Print("ipk : ")
				fmt.Scan(&ipk[i])
				k++
			}
			fmt.Print("kembali? (y/n) : ")
			fmt.Scan(&kembali)
			if kembali == "y" {
				iya = true
			} else {
				iya = false
			}
		}
		if pilihan == 2 {
			for i := 0; i < k; i++ {
				fmt.Println("nama : ", nama[i])
				fmt.Println("nim : ", nim[i])
				fmt.Println("ipk : ", ipk[i])
			}
			fmt.Print("kembali? (y/n) : ")
			fmt.Scan(&kembali)
			if kembali == "y" {
				iya = true
			} else {
				iya = false
			}
		}
		if pilihan == 3 {
			os.Exit(0)
		}
	}
}
