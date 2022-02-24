package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitor = 3
const delay = 5

func main() {
	readsitesArqui()
	execIntro()

	for {
		showmenu()

		comand := readcomand()

		switch comand {
		case 1:
			startmonitor()
		case 2:
			fmt.Println("Exibindo logs... ")
		case 0:
			fmt.Println("Thau! Até mais! ")
			os.Exit(0)

		default:
			fmt.Println("Desconheço esse comando.")
			os.Exit(-1)

		}
	}

}

func execIntro() {
	name := "Thailan"
	version := 0.1

	fmt.Println("Olá Sr.", name)
	fmt.Println("Este programa está na versão: ", version)

}

func readcomand() int {

	var comand int

	fmt.Scanf("%d", &comand)

	return comand
}

func showmenu() {

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func startmonitor() {
	fmt.Println("Monitorando... ")

	sites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br/", "https://www.signativa.com.br"}

	for i := 0; i < monitor; i++ {

		for i, site := range sites {
			fmt.Println("Testando site", i, ": ", site)
			testingsite(site)
			fmt.Println("")
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func testingsite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "Foi carregado com sucesso.")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", response.StatusCode)
	}

}

func readsitesArqui() []string {

	arquivo, err := os.Open("sites.txt")
	fmt.Println(arquivo)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)

	}

	var sites []string

	return sites
}
