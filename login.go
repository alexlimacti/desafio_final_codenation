package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	email string
	pass  []byte
}

func u() user {
	e := "alf@g.com"
	p := "test"
	//he, _ := bcrypt.GenerateFromPassword([]byte(e), 14)
	hp, _ := bcrypt.GenerateFromPassword([]byte(p), 14)
	m := user{e, hp}
	
	fmt.Println(hp)
	return m
}
func login(w, s string) bool {
	m := u()
	hp := bcrypt.CompareHashAndPassword([]byte(m.pass), []byte(s))
	//he := bcrypt.CompareHashAndPassword([]byte(m.email), []byte(w))
	return hp == nil && m.email == w
}

func main() {

	fmt.Println(login("alf@g.com", "test"))
}
