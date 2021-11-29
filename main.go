package main

import (
	"fmt"
	"oop-go/contas"
)

func main() {
	contaDoDenis := contas.ContaCorrente{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	contaDaLuisa := contas.ContaPoupanca{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 400)

	fmt.Println(contaDoDenis)
	fmt.Println(contaDaLuisa)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}
