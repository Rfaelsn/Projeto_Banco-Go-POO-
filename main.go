package main

import (
	"Banco/contas"
	"fmt"
)

func main() {
	contaExemplo := contas.ContaCorrente{}

	contaExemplo.Depositar(-500)

	fmt.Println(contaExemplo.ObterSaldo())
}
