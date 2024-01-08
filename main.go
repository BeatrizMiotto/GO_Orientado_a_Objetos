package main

import (
	"fmt"

	"\banco\clientes\"
)

func main() {

	clientenovo := clientes.Titular{"Mariana", "123.789.456-10", "Desenvolvedora"}
	contaNova := contas.ContaCorrente{clientenovo, 123, 123456, 300}

	//status := contaNova.Transferir(100, &contaNova1)

	fmt.Println(contaNova)

}
