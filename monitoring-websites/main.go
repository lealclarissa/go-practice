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

const monitoring = 2
const delay = 5

func main() {

	showIntroduction()
	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
			fmt.Println("")
			printLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	version := 1.3
	var name string
	fmt.Println("Olá, antes de iniciarmos, qual seu nome?")
	fmt.Scan(&name)
	fmt.Println("")
	fmt.Println("Olá, ", name, "!", "Este programa está na versão", version)
	fmt.Println("O que deseja fazer? Por favor, digite um dos números abaixo:")
	fmt.Println("")
}

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)
	fmt.Println("O comando escolhido foi:", commandRead)
	fmt.Println("")

	return commandRead
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	fmt.Println("")

	websites := readWebsitesFromFile()

	for i := 0; i < monitoring; i++ {
		for i, website := range websites {
			fmt.Println("Testando site", i, ":", website)
			testWebsite(website)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testWebsite(website string) {
	resp, err := http.Get(website)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site", website, "foi carregado com sucesso!")
		saveLog(website, true)
	} else {
		fmt.Println("O site", website, "está com problemas! Status Code:", resp.StatusCode)
		saveLog(website, false)
	}
}

func readWebsitesFromFile() []string {
	var websites []string

	file, err := os.Open("websites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		websites = append(websites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return websites
}

func saveLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- está online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(string(file))
}
