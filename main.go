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

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar((valorTransferencia))
		return true
	} else {
		return false
	}
}

func main() {

	contaNova := ContaCorrente{titular: "Ana", saldo: 1000}
	contaNova1 := ContaCorrente{titular: "Maria", saldo: 300}

	status := contaNova.Transferir(100, &contaNova1)

	fmt.Println(status)
	fmt.Println(contaNova)
	fmt.Println(contaNova1)

}
