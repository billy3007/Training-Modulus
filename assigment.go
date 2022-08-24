package main

import (
	"fmt"
	"os"
	"strconv"
)

type Manusia struct {
	name string
	alamat string
	pekerjaan string
}

func main() {

	// cetak := func(l []string) {
	// 	for _, v := range kelasA {
	// 		fmt.Println(v.name)
	// 	}
	// }
	// cetak
	key := os.Args[1]
	i, _ := strconv.Atoi(key)
	var res = listPeople(i)
	fmt.Println(res)
	
}

func listPeople(i int) Manusia{
	var kelasA = []Manusia{
		{name: "Billy", alamat: "Jakarta", pekerjaan: "Programmer"},
		{name: "Hendar", alamat: "Jakarta", pekerjaan: "Programmer"},
		{name: "Zaka", alamat: "Jakarta", pekerjaan: "Programmer"},
		{name: "Achmad", alamat: "Jakarta", pekerjaan: "Programmer"},
		{name: "Fadli", alamat: "Jakarta", pekerjaan: "Programmer"},
		{name: "David", alamat: "Jakarta", pekerjaan: "Programmer"},
		{name: "Ziyadatun", alamat: "Jakarta", pekerjaan: "Programmer"},
		{name: "Zayan", alamat: "Jakarta", pekerjaan: "Programmer"},
		{name: "Fajar", alamat: "Jakarta", pekerjaan: "Programmer"},
		{name: "Rizky", alamat: "Jakarta", pekerjaan: "Programmer"},
	}

	return kelasA[i]
}