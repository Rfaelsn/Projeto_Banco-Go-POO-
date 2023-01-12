package main

import (
	"Banco/clientes"
	"Banco/contas"
	"fmt"
)

func main() {
	clienteBruno := clientes.Titular{"Bruno", "123.543.543-65", "Desenvolvedor"}

	contaDoBruno := contas.ContaCorrente{clienteBruno, 897, 123456, 5000}

	fmt.Println(contaDoBruno)

}
