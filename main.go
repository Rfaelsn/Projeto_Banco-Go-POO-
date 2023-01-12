package main

import (
	"Banco/contas"
	"fmt"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{"Silvia", 890, 123456, 400}
	contaDoRafael := contas.ContaCorrente{"Rafael", 895, 123457, 200}

	status := contaDaSilvia.Transferir(-300, &contaDoRafael)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoRafael)

}
