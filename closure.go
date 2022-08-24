package main

import (
	"fmt"
)

func main() {
	var kelasA = []string{"Billy", "Hendar", "Zaka", "Achmad", "Fadli", "David", "Ziyadatun"}
	var list []*string
	for i := 0; i < len(kelasA); i++ {
		list = append(list, &kelasA[i])
	}
	
	cetak := func(l []*string) {
		for _, v := range list {
			fmt.Println(*v)
		}
	}
	cetak(list)
}
