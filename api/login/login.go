package login

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// User é a struct recebida no formulário de login
type User struct {
	email string
	pass  []byte
}

// Login é a função principal do package
func Login(w, s string) bool {
	m := data()
	hp := bcrypt.CompareHashAndPassword([]byte(m.pass), []byte(s))
	return hp == nil && m.email == w
}

func data() User {
	bytes, err := bcrypt.GenerateFromPassword([]byte("12345"), 14)
	if err != nil {
		fmt.Println(err)
	}
	m := User{"alph@gmail.com", bytes}
	return m
}
