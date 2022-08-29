package main

import (
	"fmt"
	"training/Training-Modulus/service"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)
	res := userSvc.Register(&service.User{Nama: "Billy"})
	fmt.Println(res)
	res = userSvc.Register(&service.User{Nama: "cahya"})
	fmt.Println(res)
	resGet := userSvc.GetUser()
	fmt.Println("---Hasil get user---")
	for _, v := range resGet {
		fmt.Println(v.Nama)
	}
}