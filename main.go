package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque Realizado com Sucesso"
	} else {
		return "Saldo insuficiênte"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Depósito Realizado com sucesso", c.saldo
	} else {
		return "Falha no depósito", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if c.saldo > valorTransferencia && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	contaDaSilvia := ContaCorrente{"Silvia", 890, 123456, 400}
	contaDoRafael := ContaCorrente{"Rafael", 895, 123457, 200}

	status := contaDaSilvia.Transferir(-300, &contaDoRafael)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoRafael)

}
