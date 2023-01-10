package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contaDoRafael := ContaCorrente{titular: "Rafael", numeroAgencia: 833,
		numeroConta: 123456, saldo: 5000}

	contaDaKerssia := ContaCorrente{"Kérssia", 834, 123457, 10000}
	fmt.Println(contaDoRafael)
	fmt.Println(contaDaKerssia)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	fmt.Println(*contaDaCris)

}
