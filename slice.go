package main

import "fmt"

func main() {
	var kelasA = []string{"Billy", "Hendar", "Zaka", "Achmad", "Fadli", "David", "Ziyadatun"}
	kelasA = append(kelasA, "Zayan", "Fajar", "Rizky")

	for _, value := range kelasA {
		fmt.Println(value)
	}
}
