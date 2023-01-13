package main

import (
	"Banco/contas"
	"fmt"
)

func main() {
	contaEmanuel := contas.ContaPoupanca{}
	contaEmanuel.Depositar(500)
	contaEmanuel.Sacar(200)
	fmt.Println(contaEmanuel.ObterSaldo())
}
