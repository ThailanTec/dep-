package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitor = 3
const delay = 5

func main() {
	execIntro()

	for {
		showmenu()

		comand := readcomand()

		switch comand {
		case 1:
			startmonitor()
		case 2:
			fmt.Println("Exibindo logs... ")
			printlog()
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

	sites := readsitesArqui()
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
		registerlog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", response.StatusCode)
		registerlog(site, false)
	}

}

func readsitesArqui() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)

	}

	reader := bufio.NewReader(arquivo)

	for {
		linha, err := reader.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break

		}

	}
	arquivo.Close()

	return sites
}

func registerlog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(arquivo)

	var convertido = strconv.FormatBool(status)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site + " - online: " + convertido + "\n")

	arquivo.Close()

}

func printlog() {

	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {

		fmt.Println(err)

	}
	fmt.Println(string(arquivo))
}
