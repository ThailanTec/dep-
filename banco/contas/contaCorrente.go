package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Saque(valorDosaque float64) string {

	podesacar := valorDosaque > 0 && valorDosaque <= c.saldo

	if podesacar {
		c.saldo -= valorDosaque
		return "Saque realizado com sucesso"
	} else {
		return "Valor não é suficiente"
	}

}

func (c *ContaCorrente) Deposito(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		c.saldo += valorDeposito
	}
	return "Deposito realizado com sucesso no valor:", c.saldo

}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) (string, bool) {

	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Deposito(valorTransferencia)

		return "Transferencia realizado com sucesso", true
	} else {
		return "Não foi possivel realizar tranferencia", false
	}
}

func (c *ContaCorrente) Obtersaldo() float64 {
	return c.saldo
}
