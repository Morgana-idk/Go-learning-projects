package main

import "fmt"

func main() {
	nomes := []string{
		"Pedro",
		"Davi",
		"Alysson",
		"Nicole",
	}
	nomes = append(nomes, "Não sei")
	for _, nome := range nomes {
		fmt.Println((nome))
	}
	fmt.Println("brutal")
}
d