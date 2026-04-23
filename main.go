package main

import "fmt"

type pessoas []struct {
	nome  string
	forca int
}

func main() {
	pessoasMain := pessoas{
		{nome: "Pedro", forca: 90},
		{nome: "Davizin", forca: 90},
		{nome: "Alysson", forca: 85},
		{nome: "Nicole", forca: 100},
		{nome: "Modulescript", forca: 1000000000000000000},
	}

	for _, pessoa := range pessoasMain {
		if pessoa.forca >= 90 {
			if pessoa.nome == "Pedro" {
				fmt.Println("[Master] --", pessoa.nome, "[RIVAL] --", pessoa.forca)
			} else {
				fmt.Println("[Master] --", pessoa.nome, "--", pessoa.forca)
			}
		} else {
			if pessoa.nome == "Pedro" {
				fmt.Println("[Master] --", pessoa.nome, "[RIVAL] --", pessoa.forca)
			} else {
				fmt.Println("[Grinding] --", pessoa.nome, "--", pessoa.forca)
			}
		}
	}
}
