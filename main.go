package main

import "fmt"

func main() {
	var nome string = "Pedro"
	var idade int = 12
	if idade >= 12 && idade < 18 {
		fmt.Println(nome, "tem", idade, "anos. E é adolescente")
	} else if idade < 12 {
		fmt.Println(nome, "tem", idade, "anos. E é criança")
	} else {
		fmt.Println(nome, "tem", idade, "anos. E é adulto")
	}
}
