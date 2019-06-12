package compare


import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type publico struct {
	nome string
	sal  int
}

func main() {

	c, _ := clientes()
	d, _ := data()
	e := compare(c, d)
	fmt.Println(e)
}

func data() ([]publico, error) {
	f, err := os.Open("./arquivos brutos/Remuneracao.txt")

	if err != nil {
		return nil, err
	}

	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ';'

	d, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	var lista []publico

	for i, c := range d {
		v := strings.Split(c[3], ",")
		s, _ := strconv.Atoi(v[0])
		if i != 0 {
			lista = append(lista, publico{c[0], s})

		}
	}

	return lista, nil

}

func clientes() ([]string, error) {
	var clientes []string
	f, err := os.Open("./arquivos brutos/clientes.csv")

	if err != nil {
		return nil, err
	}

	r := csv.NewReader(bufio.NewReader(f))

	d, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	for _, c := range d {
		clientes = append(clientes, c[0])
	}

	return clientes, nil
}

func compare(w []string, p []publico) []publico {
	var bc []publico
	for _, c := range p {
		if c.sal >= 20000 {
			for _, b := range w {
				if c.nome == b {
					bc = append(bc, c)
				}
			}
		}
	}

	return bc
}
