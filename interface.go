package main

import (
	"sync"
	"fmt"
	"training/Training-Modulus/service"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)
	res := userSvc.Register(&service.User{Nama: "Billy"})
	fmt.Println(res)
	res = userSvc.Register(&service.User{Nama: "Suhendar"})
	fmt.Println(res)
	resGet := userSvc.GetUser()
	fmt.Println("---Hasil get user---")
	
	var wg sync.WaitGroup
	wg.Add(len(resGet))
	for _, v := range resGet {
		go cetakNama(&wg, v.Nama)
	}
	wg.Wait()
}

func cetakNama (wg *sync.WaitGroup, name string) {
	fmt.Println(name)
	wg.Done()
}