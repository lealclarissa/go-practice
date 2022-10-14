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
			fmt.Printf("\nExibindo Logs...\n")
			printLogs()
		case 0:
			fmt.Printf("\nSaindo do programa...\n")
			os.Exit(0)
		default:
			fmt.Printf("\nNão conheço este comando\n")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	version := 1.4
	var name string
	fmt.Println("Olá, antes de iniciarmos, qual seu nome?")
	fmt.Scan(&name)

	fmt.Printf("\nOlá, %v! Este programa está na versão %v\nO que deseja fazer? Por favor, digite um dos números abaixo:\n", name, version)
}

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento\n2 - Exibir Logs \n0 - Sair do Programa")
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)

	fmt.Printf("O comando escolhido foi: %v", commandRead)

	return commandRead
}

func startMonitoring() {
	fmt.Printf("\nMonitorando...\n")

	websites := readWebsitesFromFile()

	for i := 0; i < monitoring; i++ {
		for i, website := range websites {
			fmt.Printf("\nTestando site %v: %v\n", i, website)
			testWebsite(website)
		}
		time.Sleep(delay * time.Second)
	}
}

func errorMessage(err error) {
	fmt.Printf("Ocorreu um erro: %v", err)
}

func testWebsite(website string) {
	resp, err := http.Get(website)

	if err != nil {
		errorMessage(err)
	}

	if resp.StatusCode == 200 {
		fmt.Printf("O site %v foi carregado com sucesso!\n", website)
		saveLog(website, true)
	} else {
		fmt.Printf("O site: %v está com problemas! Status Code: %v\n", website, resp.StatusCode)
		saveLog(website, false)
	}
}

func readWebsitesFromFile() []string {
	var websites []string

	file, err := os.Open("websites.txt")
	if err != nil {
		errorMessage(err)
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
		errorMessage(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- está online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		errorMessage(err)
	}

	fmt.Println(string(file))
}
