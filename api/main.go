package main

import (
	"fmt"
	"github.com/marcospedro97/desafio_final_codenation/tree/creating-view-and-api-dir/api/files/compare"
)

func main() {
	c, err := files.clientes()
	if err != nil {
		fmt.Println("err")
	}
}
