package main

import (
	"fmt"
)

func main() {
	c, err := files.clientes()
	if err != nil {
		fmt.Println("err")
	}
}
