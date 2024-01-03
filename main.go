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

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso", c.saldo //mutiplos retornos
	} else {
		return "Valor invalido ", c.saldo
	}
}

func main() {

	contaNova := ContaCorrente{}
	contaNova.titular = "Ana"
	contaNova.saldo = 500

	fmt.Println(contaNova.saldo)

	fmt.Println(contaNova.Depositar(500))

}
