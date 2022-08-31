package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Nama string `json:"nama"`
}

type UserService struct {
	db []*User
}

type UserIface interface {
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	GetUserHandler(w http.ResponseWriter, r *http.Request)
	Register(u *User) []string
	GetUser() []*User
}

func NewUserService(db []*User) *UserService {
	return &UserService{
		db: db,
	}
}

func (u *UserService) Register(user *User) string {
	u.db = append(u.db, user)
	return user.Nama + " berhasil didaftarkan"
}

func (u *UserService) GetUser() []*User {
	return u.db
}

func (u *UserService) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var user User
		err := decoder.Decode(&user)
		if err != nil {
			fmt.Println("Error data user")
			return
		}
		u.db = append(u.db, &user)
		w.Write([]byte(user.Nama + " berhasil ditambahkan"))
	} else {
		w.Write([]byte("invalid http method"))
	}
}

func (u *UserService) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(u.db)
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
