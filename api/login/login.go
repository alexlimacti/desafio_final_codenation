package login

import (
	"golang.org/x/crypto/bcrypt"
)

// User é a struct recebida no formulário de login
type User struct {
	email string
	pass  []byte
}

// Login é a função principal do package
func Login(w, s string) bool {
	hp := bcrypt.CompareHashAndPassword([]byte(m.pass), []byte(s))
	return hp == nil && m.email == w
}
