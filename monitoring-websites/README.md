<h1 align="center">Monitorando Sites</h1>

<p align="center">
  <a href="#o-que-é">O que é</a> •
  <a href="#funcionalidades">Funcionalidades</a> •
  <a href="#como-usar">Como usar</a> •
  <a href="#referência">Referência</a>
</p>  

## O que é:   

Monitor de sites em Go. Este projeto é fruto do curso *Go: iniciando com a linguagem do Google*, da Alura. Neste repositório, no entanto, fiz minhas próprias modificações para além do projeto final.   

## Funcionalidades:

- [x] Monitora sites   
    - [x] Identifica se o site está online ou se ocorreu algum erro (exibe o *status code*)
    - [x] Atualmente realiza 2 monitoramentos, com 5 segundos de diferença entre eles   
- [x] Exibe logs de monitoramento com as seguintes informações:
    - Horário do teste  
    - Url do site testado   
    - Status

## Como usar:  

*1* - Você deve ter Go instalado

[Download aqui](https://golang.org/dl/)

*2* - Clone o projeto 

```bash
git clone https://github.com/lealclarissa/go-practice.git
```

*3* - Acesse a pasta */monitoring-websites*  

```bash
cd go-practice/monitoring-websites 
```   

*4* - Rode o projeto 

```bash  
go run main.go    
```   

*5* - Deverá aparecer uma mensagem de saudação inicial perguntando seu nome: responda.  

*6* - Em seguida, será exibido o menu de opções do programa(monitorar, exibir logs, sair do programa). Divirta-se.

---

## Referência:

- [Curso de Go - Alura](https://www.alura.com.br/curso-online-golang)