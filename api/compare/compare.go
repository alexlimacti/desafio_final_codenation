package compare

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

// Publico contém o nome do funcionário e o salário do mês atual
type Publico struct {
	nome string
	sal  int
}

// Run faz todo o trabalho de comparar os dois arquivos de entrada
func Run() ([]Publico, error) {
	d, err := data()
	if err != nil {
		return nil, err
	}
	c, err := clientes()
	if err != nil {
		return nil, err
	}
	R := comparar(c, d)

	return R, nil
}

func data() ([]Publico, error) {
	f, err := os.Open("./arquivos/Remuneracao.txt")

	if err != nil {
		return nil, err
	}

	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ';'

	d, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	var lista []Publico

	for i, c := range d {
		v := strings.Split(c[3], ",")
		s, _ := strconv.Atoi(v[0])
		if i != 0 {
			lista = append(lista, Publico{c[0], s})

		}
	}

	return lista, nil

}

func clientes() ([]string, error) {
	var clientes []string
	f, err := os.Open("./arquivos/clientes.csv")

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

func comparar(w []string, p []Publico) []Publico {
	var bc []Publico
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
