package main

import "fmt"

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "saque realizado com sucesso"
	} else {
		return " sem saldo"
	}
}

func main() {
	/*//conta := ContaCorrente{titular: "Miotto", agencia: 5023, conta: 123456, saldo: 123.45}
	conta2 := ContaCorrente{"Bia", 5023, 789456, 123.45}
	conta5 := ContaCorrente{"Bia", 5023, 789456, 123.45}
	//conta3 := ContaCorrente{titular: "Miotto", saldo: 123.45}
	fmt.Println(conta2 == conta5)

	var conta4 *ContaCorrente
	conta4 = new(ContaCorrente)
	conta4.titular = "Larissa"

	fmt.Println(*conta4)*/

	contaNova := ContaCorrente{}
	contaNova.titular = "Ana"
	contaNova.saldo = 500

	fmt.Println(contaNova.saldo)

	fmt.Println(contaNova.Sacar(200))
	fmt.Println(contaNova.saldo)

}
