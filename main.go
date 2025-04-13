package main

import "fmt"

func main() {
	var lista []string

	lista = append(lista, "Teste")
	lista = append(lista, "Teste 2")

	itemLista := lista[1]

	fmt.Println(itemLista)
}
