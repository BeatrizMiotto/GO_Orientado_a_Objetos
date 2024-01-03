package main

import (
	"github.com//BeatrizMiotto/GO_Orientado_a_Objetos/banco/contas"

	"fmt"
)

func main() {

	contaNova := contas.ContaCorrente{Titular: "Ana", Saldo: 1000}
	contaNova1 := contas.ContaCorrente{Titular: "Maria", Saldo: 300}

	status := contaNova.Transferir(100, &contaNova1)

	fmt.Println(status)
	fmt.Println(contaNova)
	fmt.Println(contaNova1)

}
