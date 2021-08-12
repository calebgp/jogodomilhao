package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

type Pergunta struct {
	Questao  string
	Op1      string
	Op2      string
	Op3      string
	Op4      string
	Gabarito string
}

var fileScanner *bufio.Scanner
var perguntas []Pergunta

func proximaLinha() string {
	fileScanner.Scan()
	return fileScanner.Text()
}

func main() {
	rand.Seed(1000)
	stdin := bufio.NewReader(os.Stdin)
	file, err := os.Open("perguntas.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner = bufio.NewScanner(file)

	for fileScanner.Scan() {
		pergunta := Pergunta{
			Questao:  proximaLinha(),
			Op1:      proximaLinha(),
			Op2:      proximaLinha(),
			Op3:      proximaLinha(),
			Op4:      proximaLinha(),
			Gabarito: proximaLinha(),
		}
		perguntas = append(perguntas, pergunta)
	}

	premiacoes := [11]int{500, 1000, 2500, 5000, 10000, 25000, 50000, 100000, 200000, 500000, 1000000}
	fmt.Printf("Seja bem vindo ao Jogo do Calebão!\n")
	fmt.Printf("Deseja começar?\n 1- Sim\n 2- Não\n ")
	var comecar string
	var comecou bool
	fmt.Scanf("%s", &comecar)
	if comecar == "1" {
		comecou = true
	} else {
		fmt.Printf(" Volte Sempre!\n")
		comecou = false
	}
	if comecou == true {
		premio := 0
		for {
			var comando string
			fmt.Scanf("%s", &comando)
			for comando != "Q" {
				fmt.Printf("Valendo R$ %d!\n", premiacoes[premio])

				pergunta := perguntas[rand.Intn(len(perguntas)-1)]

				fmt.Printf("%s\n\n", pergunta.Questao)
				fmt.Printf("1 - %s\n", pergunta.Op1)
				fmt.Printf("2 - %s\n", pergunta.Op2)
				fmt.Printf("3 - %s\n", pergunta.Op3)
				fmt.Printf("4 - %s\n", pergunta.Op4)
				for {
					_, err := fmt.Scanf("%s\n", &comando)
					if err == nil {
						break
					}

					stdin.ReadString('\n')
					fmt.Println("Sorry, invalid input. Please enter an integer: ")
				}
				if comando == "C" {
					fmt.Printf("Escolha uma carta (número de 1 a 4)\n")
					fmt.Scanf("%s", &comando)
					rand.Intn(4)
				}
				if comando == "U" {

				}
				if comando > "4" || comando < "0" {
					fmt.Printf("Resposta Inválida, Digite outro número por favor\n")
					fmt.Scanf("%s", &comando)
				}
				if comando != pergunta.Gabarito {
					comando = "Q"
					fmt.Printf("Game Over!\n")
					if premio >= 2 {
						fmt.Printf("Você ganhou %d reais\n", premiacoes[premio-2])
					} else {
						fmt.Printf("Você não ganhou nada!\n")
					}
					premio = 0
					break
				} else {
					fmt.Printf("Correto!\n")
					premio++
				}
				if premio == 11 {
					fmt.Printf("Parabéns você é o grande vencedor!\n")
					break
				}

			}
		}
	}
}
