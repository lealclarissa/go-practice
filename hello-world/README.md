<h1 align="center">Hello, World! 👋🏽</h1>  


<p align="center">
  <a href="#o-que-é">O que é</a> •
  <a href="#como-iniciar">Como iniciar</a> •
  <a href="#como-usar">Como usar</a>
</p>  

---
## O que é:

O bom e velho primeiro passo ao aprender qualquer linguagem de programação: imprimir uma frase no terminal. Por conta da [maldição](https://digitalinnovation.one/artigos/cuidado-para-os-iniciantes-a-maldicao-do-ola-mundo), recomenda-se que seja o bom e velho "Hello, World!". A escolha da frase fica por sua conta e risco. 😂

## Como iniciar:  
> Opção para imprimir uma frase do zero, a partir de um repositório seu. E partindo do princípio de que já tem [Go](https://golang.org/dl/) instalado na máquina

* Crie uma pasta para seu projeto

```bash
mkdir nome_do_projeto
````

* Acesse a pasta do projeto

```bash
cd nome_do_projeto
```

* Inicialize o módulo:  

```bash
go mod init <url_do_provedor/usuario_ou_organizacao/nome_do_projeto>
```   

> Para este projeto, usei o seguinte: `go mod init github.com/lealclarissa/go-practice/tree/main/hello-world`  

* Crie um arquivo `main.go`, onde rodará o projeto

```bash
touch main.go
```

* Abra o arquivo `main.go` na IDE de sua preferência

```bash
code main.go .
```  

* Faça a declaração do pacote no início do código. Ex: `package main`  

* Importe o que for necessário para rodar o projeto. Neste caso: `import "fmt"`    

* Crie a função principal - *main* - a partir da qual a mágica acontecerá.    
    ```go
    func main() {  
        fmt.Println("Olá, mundo!")
    }
    ```  
  
* Compile e rode o projeto com o seguinte comando:  

```bash  
go run main.go    
```  

> Pronto, em seu terminal deve ter aparecido: `Hello, World!`

---

## Como usar:  
> Opção para apenas rodar este projeto, visualizando o "Hello, World" no terminal.  

*1* - Você deve ter Golang instalado

[Download aqui](https://golang.org/dl/)

*2* - Clone o projeto 
> se já tiver o projeto clonado, pule esta etapa.

```bash
git clone https://github.com/lealclarissa/go-practice.git
```

*3* - Acesse a pasta *hello-world/*  

```bash
cd go-practice/hello-world/
```   

*4* - Compile e rode o projeto com o seguinte comando:  

```bash  
go run main.go    
```  

> Pronto, em seu terminal deve ter aparecido: `Hello, World!`