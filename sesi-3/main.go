package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	nama string
	alamat string
	pekerjaan string
	alasan string
}


func main() {
	args := os.Args[1:]

        if len(args) == 0 {
		fmt.Println("Tolong masukan nama atau nomor absen")
	        os.Exit(0)
	}
	
	var temanKelas = []Teman{
		{nama:"reni", alamat:"bekasi", pekerjaan:"product analyst", alasan:"keren"},
		{nama:"manda", alamat:"cinere", pekerjaan:"head of data", alasan:"mantap"},
		{nama:"nadya", alamat:"deplu", pekerjaan:"ux researcher", alasan:"cepat"},
		{nama:"jilly", alamat:"bintaro", pekerjaan:"software engineer", alasan:"simpel"},
		{nama:"dizty", alamat:"nusaloka", pekerjaan:"lead qa", alasan:"ikut jilly"},
		{nama:"triisya", alamat:"karang tengah", pekerjaan:"qa", alasan:"ikut dizty"},
	}

	arg := args[0]
	
	// Cek jika input merupakan int
	cariNama := false
	x , err:= strconv.Atoi(arg)
	if err != nil {
        	cariNama = true
        }

	found := false
	for index, value := range temanKelas {
		if cariNama {
			found = value.nama == arg
		} else {
			found = index == x
		}

		if found {
			fmt.Println("id:", index)
			fmt.Println("nama:", value.nama)
			fmt.Println("alamat:", value.alamat)
			fmt.Println("pekerjaan:", value.pekerjaan)
			fmt.Println("alasan:", value.alasan)
			break
		}
	}

	if !found {
		fmt.Println("Tidak ada teman dengan nama atau index:", arg)
	}
	
}
