package main

import (
	"fmt"
)

type Manusia struct {
	name string
}

func main() {
	var kelasA = []Manusia{
		{name: "Billy"},
		{name: "Hendar"},
		{name: "Zaka"},
		{name: "Achmad"},
		{name: "Fadli"},
		{name: "David"},
		{name: "Ziyadatun"},
		{name: "Zayan"},
		{name: "Fajar"},
		{name: "Rizky"},
	}
	var list []*string
	for i := 0; i < len(kelasA); i++ {
		list = append(list, &kelasA[i].name)
	}
	
	cetak := func(l []*string) {
		for _, v := range list {
			fmt.Println(*v)
		}
	}
	cetak(list)
}
