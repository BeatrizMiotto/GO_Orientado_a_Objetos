package main

import "fmt"

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func main() {
	conta := ContaCorrente{titular: "Miotto", agencia: 5023, conta: 123456, saldo: 123.45}
	conta2 := ContaCorrente{"Bia", 5023, 789456, 123.45}
	conta3 := ContaCorrente{titular: "Miotto", saldo: 123.45}
	fmt.Println(conta, conta2, conta3)

}
