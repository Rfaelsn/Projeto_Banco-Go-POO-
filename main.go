package main

import (
	"Banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaEmanuel := contas.ContaPoupanca{}
	contaEmanuel.Depositar(500)
	PagarBoleto(&contaEmanuel, 100)

	fmt.Println(contaEmanuel.ObterSaldo())

	contaRafael := contas.ContaCorrente{}
	contaRafael.Depositar(5000)
	PagarBoleto(&contaRafael, 1000)

	fmt.Println(contaRafael)

}
