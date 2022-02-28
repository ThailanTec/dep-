package contas

import "banco/clientes"

type contaPoupanca struct {
	Titular     clientes.Titular
	NumeroAG    int
	NumeroConta int
	Operacao    int
	saldo       float64
}
