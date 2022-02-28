package main

import (
	"banco/contas"
	"fmt"
)
import "banco/clientes"

func main() {

	clienteThailan := clientes.Titular{"Thailan", "123414515", "Developer"}
	contaThailan := contas.ContaCorrente{Titular: clientes.Titular{"Thailan", "123414515", "Developer"}, NumeroConta: 21231, NumeroAgencia: 1223}
	fmt.Println(clienteThailan, contaThailan)
}
