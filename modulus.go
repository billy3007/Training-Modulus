package main

import "fmt"

func main() {
	var count = 10;
	for i := 1; i <= count; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "genap")	
		} else {
			fmt.Println(i, "ganjil")	
		}
	}
}
